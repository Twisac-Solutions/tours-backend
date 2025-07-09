package utils

import (
	"context"
	"mime/multipart"
	"time"

	"github.com/Twisac-Solutions/tours-backend/config"
	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

var cld *cloudinary.Cloudinary

// InitCloudinary initializes the Cloudinary client
func InitCloudinary() error {
	var err error
	cld, err = cloudinary.NewFromURL(config.CloudinaryURL)
	if err != nil {
		return err
	}
	return nil
}

// UploadImageToCloudinary uploads an image to Cloudinary and returns the URL
func UploadImageToCloudinary(file *multipart.FileHeader, folder string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Open the file
	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	// Upload to Cloudinary
	uploadResult, err := cld.Upload.Upload(ctx, src, uploader.UploadParams{
		Folder:       folder,
		ResourceType: "auto",
	})
	if err != nil {
		return "", err
	}

	return uploadResult.SecureURL, nil
}
