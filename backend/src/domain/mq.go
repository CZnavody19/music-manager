package domain

type CompleteMessage struct {
	Track    Track   `json:"track"`
	FilePath *string `json:"file_path"`
	Error    *string `json:"error"`
}
