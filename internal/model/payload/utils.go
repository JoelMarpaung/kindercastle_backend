package payload

import "mime/multipart"

type UploadImageResponse struct {
	ImageURL string `json:"image_url"`
}

type UploadImage struct {
	Image multipart.FileHeader `form:"file"`
}
