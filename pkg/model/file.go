package model

import (
	"encoding/base64"
	"log"
	"mime"
	"os"
)

// File is struct
type File struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	ContentType string `json:"contentType"`
	Base64      string `json:"base64"`
}

// Create create file
func (f File) Create() error {
	data, _ := base64.StdEncoding.DecodeString(f.Base64)

	fileType, err := mime.ExtensionsByType(f.ContentType)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	fileName := f.Name + fileType[0]
	file, err := os.Create("./uploads/" + fileName)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	defer file.Close()

	file.Write(data)

	return err
}
