package utils

import (
	"context"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"kindercastle_backend/internal/model/payload"
	"kindercastle_backend/internal/pkg/constant"
	"kindercastle_backend/internal/pkg/custom"
	"kindercastle_backend/internal/pkg/logging"

	"github.com/google/uuid"
)

func (s Service) UploadImage(ctx context.Context, fileHeader *multipart.FileHeader) (payload.UploadImageResponse, error) {
	if fileHeader == nil {
		logging.Logger().Error().Msg("uploaded file is empty")
		return payload.UploadImageResponse{}, custom.ErrInternalServer
	}

	ext := filepath.Ext(fileHeader.Filename)
	if _, ok := constant.WhitelistPictureExtension[ext]; !ok {
		return payload.UploadImageResponse{}, custom.ErrInvalidImageExtension
	}

	newFileName := uuid.New().String() + ext
	imageURL := fmt.Sprintf("/assets/%s", newFileName)

	openedImg, err := fileHeader.Open()
	if err != nil {
		logging.Logger().Error().Err(err).Msg("failed to open image")
		return payload.UploadImageResponse{}, custom.ErrInternalServer
	}
	defer openedImg.Close()

	filePath := fmt.Sprintf("assets/%s", newFileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		logging.Logger().Error().Err(err).Msg("failed to create file in assets folder")
		return payload.UploadImageResponse{}, custom.ErrInternalServer
	}
	defer outFile.Close()

	if _, err = io.Copy(outFile, openedImg); err != nil {
		logging.Logger().Error().Err(err).Msg("failed to copy image to assets folder")
		return payload.UploadImageResponse{}, custom.ErrInternalServer
	}

	return payload.UploadImageResponse{
		ImageURL: imageURL,
	}, nil
}
