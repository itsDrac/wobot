package types

type CreateUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type LoginUserPayload struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

type UploadFilePayload struct {
	FileName string `json:"file_name" validate:"required"`
}

type FileInfo struct {
	FileName string `json:"file_name" validate:"required"`
	FileSize string `json:"file_size" validate:"required"`
}

type Files struct {
	FilesInfo []FileInfo `json:"files_info" validate:"required"`
}

type ApiResponse struct {
	Message string `json:"message"`
}

type LoginUserResponsePayload struct {
	Token string `json:"token"`
}

type UserContextKey string
