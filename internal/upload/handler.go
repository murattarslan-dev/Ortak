package upload

import (
	"fmt"
	"image"
	_ "image/gif"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"ortak/internal/upload/repository"
	"ortak/internal/upload/service"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type Handler struct {
	uploadDir string
	maxSize   int64
	service   *service.Service
}

func NewHandler(uploadDir string, maxSize int64, service *service.Service) *Handler {
	os.MkdirAll(uploadDir, 0755)
	os.MkdirAll(filepath.Join(uploadDir, "thumbs"), 0755)
	return &Handler{
		uploadDir: uploadDir,
		maxSize:   maxSize,
		service:   service,
	}
}

func (h *Handler) UploadFile(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Dosya yüklenemedi",
		})
		return
	}
	defer file.Close()

	if header.Size > h.maxSize {
		c.JSON(http.StatusBadRequest, gin.H{
			"success": false,
			"message": "Dosya boyutu çok büyük",
		})
		return
	}

	fileID := uuid.New().String()
	ext := filepath.Ext(header.Filename)
	fileName := fmt.Sprintf("%s%s", fileID, ext)
	filePath := filepath.Join(h.uploadDir, fileName)

	if err := c.SaveUploadedFile(header, filePath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"success": false,
			"message": "Dosya kaydedilemedi",
		})
		return
	}

	userID, _ := c.Get("user_id")
	uploadedBy := ""
	if userID != nil {
		uploadedBy = userID.(string)
	}

	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		mimeType = header.Header.Get("Content-Type")
	}

	fileInfo := File{
		ID:           fileID,
		OriginalName: header.Filename,
		FileName:     fileName,
		Size:         header.Size,
		ContentType:  header.Header.Get("Content-Type"),
		Path:         filePath,
		UploadedBy:   uploadedBy,
		UploadedAt:   time.Now(),
		MimeType:     mimeType,
	}

	response := UploadResponse{
		FileID:   fileInfo.ID,
		FileName: fileInfo.OriginalName,
		Size:     fileInfo.Size,
		URL:      fmt.Sprintf("/uploads/%s", fileName),
		MimeType: mimeType,
	}

	if strings.HasPrefix(mimeType, "image/") {
		if file, err := os.Open(filePath); err == nil {
			if config, _, err := image.DecodeConfig(file); err == nil {
				response.Width = config.Width
				response.Height = config.Height
				fileInfo.Width = config.Width
				fileInfo.Height = config.Height
			}
			file.Close()
		}

		// Thumbnail oluştur
		if thumbPath := h.createThumbnail(filePath, fileID, ext); thumbPath != "" {
			thumbURL := fmt.Sprintf("/uploads/thumbs/%s_thumb%s", fileID, ext)
			response.ThumbURL = thumbURL
			fileInfo.ThumbURL = thumbURL
		}
	}

	// DB'ye kaydet
	dbFile := repository.File{
		ID:           fileInfo.ID,
		OriginalName: fileInfo.OriginalName,
		FileName:     fileInfo.FileName,
		Size:         fileInfo.Size,
		ContentType:  fileInfo.ContentType,
		Path:         fileInfo.Path,
		UploadedBy:   fileInfo.UploadedBy,
		UploadedAt:   fileInfo.UploadedAt,
		MimeType:     fileInfo.MimeType,
		Width:        fileInfo.Width,
		Height:       fileInfo.Height,
		ThumbURL:     fileInfo.ThumbURL,
	}

	if err := h.service.SaveFile(dbFile); err != nil {
		fmt.Printf("Failed to save file info to DB: %v\n", err)
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"message": "Dosya başarıyla yüklendi",
		"data":    response,
	})
}

func (h *Handler) GetUploadInfo(c *gin.Context) {
	id := c.Query("id")

	if id != "" {
		// Tek dosya bilgisi
		file, err := h.service.GetFile(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{
				"success": false,
				"message": "Dosya bulunamadı",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    file,
		})
	} else {
		// Tüm dosyalar
		userID, exists := c.Get("user_id")
		if !exists {
			c.JSON(http.StatusUnauthorized, gin.H{
				"success": false,
				"message": "Yetkisiz erişim",
			})
			return
		}

		// Admin kontrolü
		userRole, _ := c.Get("user_role")
		isAdmin := userRole == "admin"

		var files []repository.File
		var err error

		if isAdmin {
			// Admin tüm dosyaları görebilir
			files, err = h.service.GetAllFiles()
		} else {
			// Normal kullanıcı sadece kendi dosyalarını görebilir
			files, err = h.service.GetUserFiles(userID.(string))
		}

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"success": false,
				"message": "Dosyalar getirilemedi",
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"success": true,
			"data":    files,
		})
	}
}
