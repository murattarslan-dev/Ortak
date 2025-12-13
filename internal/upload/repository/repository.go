package repository

import (
	"database/sql"
	"time"
)

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

type Repository interface {
	Create(file File) error
	GetByID(id string) (*File, error)
	GetByUserID(userID string) ([]File, error)
	GetAll() ([]File, error)
}

type RepositoryImpl struct {
	db *sql.DB
}

func NewRepositoryImpl(db *sql.DB) Repository {
	return &RepositoryImpl{db: db}
}

func (r *RepositoryImpl) Create(file File) error {
	if r.db == nil {
		return nil
	}

	query := `
		INSERT INTO files (id, original_name, file_name, size, content_type, path, uploaded_by, uploaded_at, mime_type, width, height, thumb_url)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)
	`
	
	_, err := r.db.Exec(query, 
		file.ID, file.OriginalName, file.FileName, file.Size, 
		file.ContentType, file.Path, file.UploadedBy, file.UploadedAt,
		file.MimeType, file.Width, file.Height, file.ThumbURL,
	)
	
	return err
}

func (r *RepositoryImpl) GetByID(id string) (*File, error) {
	if r.db == nil {
		return nil, nil
	}

	query := `
		SELECT id, original_name, file_name, size, content_type, path, uploaded_by, uploaded_at, mime_type, width, height, thumb_url
		FROM files WHERE id = $1
	`
	
	var file File
	err := r.db.QueryRow(query, id).Scan(
		&file.ID, &file.OriginalName, &file.FileName, &file.Size,
		&file.ContentType, &file.Path, &file.UploadedBy, &file.UploadedAt,
		&file.MimeType, &file.Width, &file.Height, &file.ThumbURL,
	)
	
	if err != nil {
		return nil, err
	}
	
	return &file, nil
}

func (r *RepositoryImpl) GetByUserID(userID string) ([]File, error) {
	if r.db == nil {
		return []File{}, nil
	}

	query := `
		SELECT id, original_name, file_name, size, content_type, path, uploaded_by, uploaded_at, mime_type, width, height, thumb_url
		FROM files WHERE uploaded_by = $1 ORDER BY uploaded_at DESC
	`
	
	rows, err := r.db.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(
			&file.ID, &file.OriginalName, &file.FileName, &file.Size,
			&file.ContentType, &file.Path, &file.UploadedBy, &file.UploadedAt,
			&file.MimeType, &file.Width, &file.Height, &file.ThumbURL,
		)
		if err != nil {
			continue
		}
		files = append(files, file)
	}
	
	return files, nil
}

func (r *RepositoryImpl) GetAll() ([]File, error) {
	if r.db == nil {
		return []File{}, nil
	}

	query := `
		SELECT id, original_name, file_name, size, content_type, path, uploaded_by, uploaded_at, mime_type, width, height, thumb_url
		FROM files ORDER BY uploaded_at DESC
	`
	
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var files []File
	for rows.Next() {
		var file File
		err := rows.Scan(
			&file.ID, &file.OriginalName, &file.FileName, &file.Size,
			&file.ContentType, &file.Path, &file.UploadedBy, &file.UploadedAt,
			&file.MimeType, &file.Width, &file.Height, &file.ThumbURL,
		)
		if err != nil {
			continue
		}
		files = append(files, file)
	}
	
	return files, nil
}