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
