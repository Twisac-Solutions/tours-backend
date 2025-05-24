package utils

import (
	"fmt"
	"mime/multipart"
	"os"
	"path/filepath"
	"time"
)

func SaveFile(files []*multipart.FileHeader) (string, error) {
	if len(files) == 0 {
		return "", nil
	}
	file := files[0]
	uploadPath := "uploads/" + fmt.Sprintf("%d_%s", time.Now().UnixNano(), filepath.Base(file.Filename))
	if err := os.MkdirAll("uploads", os.ModePerm); err != nil {
		return "", err
	}
	if err := saveUploadedFile(file, uploadPath); err != nil {
		return "", err
	}
	return "/" + uploadPath, nil
}

func saveUploadedFile(file *multipart.FileHeader, dst string) error {
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = out.ReadFrom(src)
	return err
}
