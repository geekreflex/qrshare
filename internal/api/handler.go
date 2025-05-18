package api

import (
	"encoding/json"
	"mime"
	"net/http"
	"os"
	"path/filepath"
)

type FileItem struct {
	Name string `json:"name"`
	Path string `json:"path"`
	Type string `json:"type"` // file or directory
	Mime string `json:"mime,omitempty"`
	Size int64  `json:"size,omitempty"`
}

// GET /api/files/list?path=/some/path
func HandleListFiles(w http.ResponseWriter, r *http.Request) {
	queryPath := r.URL.Query().Get("path")
	if queryPath == "" {
		queryPath = "."
	}

	files, err := os.ReadDir(queryPath)
	if err != nil {
		http.Error(w, "Failed to read directory", http.StatusInternalServerError)
		return
	}

	var result []FileItem
	for _, file := range files {
		fullPath := filepath.Join(queryPath, file.Name())
		item := FileItem{
			Name: file.Name(),
			Path: fullPath,
		}

		if file.IsDir() {
			item.Type = "directory"
		} else {
			item.Type = "file"
			item.Mime = mime.TypeByExtension(filepath.Ext(file.Name()))
			if info, err := file.Info(); err == nil {
				item.Size = info.Size()
			}
		}

		result = append(result, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}

// GET /api/files/raw?path=/full/path/to/file
func HandleRawFile(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Query().Get("path")
	if path == "" {
		http.Error(w, "Path is required", http.StatusBadRequest)
		return
	}

	f, err := os.Open(path)
	if err != nil {
		http.Error(w, "Failed to open file", http.StatusNotFound)
		return
	}
	defer f.Close()

	contentType := mime.TypeByExtension(filepath.Ext(path))
	if contentType == "" {
		contentType = "application/octet-stream"
	}

	w.Header().Set("Content-Type", contentType)
	http.ServeFile(w, r, path)
}
