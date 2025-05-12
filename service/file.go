package service

import (
	"context"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/itsDrac/wobot/store"
	"github.com/itsDrac/wobot/utils"
)

type FileService struct {
	UploadFolder string
	store        store.Store
}

func NewFileService(s store.Store) FileService {
	uploadFolder := utils.GetStringEnv("UPLOAD_FOLDER", "uploads")
	if err := os.MkdirAll(uploadFolder, os.ModePerm); err != nil {
		log.Fatalf("Error in creating upload folder %s", err.Error())
	}
	log.Printf("Upload folder %s created", uploadFolder)
	return FileService{
		UploadFolder: uploadFolder,
		store:        s,
	}
}

func (f *FileService) UploadFile(ctx context.Context, file multipart.File, fileInfo *multipart.FileHeader) error {
	// Get user from context.
	user, ok := ctx.Value("user").(*store.User)
	if !ok {
		return fmt.Errorf("user not found in context")
	}
	// Check if user have storage to add file,
	remainingStorage := user.TotalStorage - user.CurrentStorage
	// If not return error saying storage full
	if fileInfo.Size > remainingStorage {
		log.Printf("Not enough storage space: requested %d, available %d", fileInfo.Size, remainingStorage)
		return fmt.Errorf("not enough storage space")
	}
	fileName := filepath.Clean(fileInfo.Filename)
	filePath := filepath.Join(f.UploadFolder, user.Username, fileName)
	if err := os.MkdirAll(filepath.Dir(filePath), os.ModePerm); err != nil {
		return fmt.Errorf("failed to create directory: %s", err.Error())
	}
	// Upload the file init.
	finalFile, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("failed to create file: %s", err.Error())
	}
	defer finalFile.Close()
	if _, err := io.Copy(finalFile, file); err != nil {
		return fmt.Errorf("failed to copy file: %s", err.Error())
	}
	// Add the file_size to current_storage in database.
	if err := f.store.UpdateCurrentStorage(ctx, user, fileInfo.Size); err != nil {
		return fmt.Errorf("failed to update current storage: %s", err.Error())
	}
	return nil
}

func (f *FileService) GetRemainingStorage(ctx context.Context) (string, error) {
	// Get user from context.
	user, ok := ctx.Value("user").(*store.User)
	if !ok {
		return "", fmt.Errorf("user not found in context")
	}
	// Get the remaining storage.
	remainingStorage := user.TotalStorage - user.CurrentStorage
	return utils.FormatBytes(remainingStorage), nil
}
