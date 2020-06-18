package model

// File is struct
type File struct {
	FileID      string `json:"fileID"`
	FileName    string `json:"fileName"`
	ContentType string `json:"contentType"`
	FileData    string `json:"fileData"`
}
