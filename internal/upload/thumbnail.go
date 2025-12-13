package upload

import (
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"github.com/nfnt/resize"
)

func (h *Handler) createThumbnail(originalPath, fileID, ext string) string {
	file, err := os.Open(originalPath)
	if err != nil {
		return ""
	}
	defer file.Close()

	img, _, err := image.Decode(file)
	if err != nil {
		return ""
	}

	// 200x200 thumbnail
	thumbnail := resize.Thumbnail(200, 200, img, resize.Lanczos3)

	thumbName := fileID + "_thumb" + ext
	thumbPath := filepath.Join(h.uploadDir, "thumbs", thumbName)

	thumbFile, err := os.Create(thumbPath)
	if err != nil {
		return ""
	}
	defer thumbFile.Close()

	// Format'a göre encode et
	switch strings.ToLower(ext) {
	case ".jpg", ".jpeg":
		err = jpeg.Encode(thumbFile, thumbnail, &jpeg.Options{Quality: 80})
	case ".png":
		err = png.Encode(thumbFile, thumbnail)
	default:
		return ""
	}

	if err != nil {
		os.Remove(thumbPath)
		return ""
	}

	return thumbPath
}