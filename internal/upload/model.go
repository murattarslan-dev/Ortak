package upload

import "time"

type File struct {
	ID           string    `json:"id"`
	OriginalName string    `json:"original_name"`
	FileName     string    `json:"file_name"`
	Size         int64     `json:"size"`
	ContentType  string    `json:"content_type"`
	Path         string    `json:"path"`
	UploadedBy   string    `json:"uploaded_by"`
	UploadedAt   time.Time `json:"uploaded_at"`
	MimeType     string    `json:"mime_type"`
	Width        int       `json:"width"`
	Height       int       `json:"height"`
	ThumbURL     string    `json:"thumb_url"`
}

type UploadResponse struct {
	FileID    string `json:"file_id"`
	FileName  string `json:"file_name"`
	Size      int64  `json:"size"`
	URL       string `json:"url"`
	ThumbURL  string `json:"thumb_url,omitempty"`
	MimeType  string `json:"mime_type"`
	Width     int    `json:"width,omitempty"`
	Height    int    `json:"height,omitempty"`
}