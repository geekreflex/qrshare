package files

import (
	"encoding/json"
	"net/http"
	"os"
	"path/filepath"
)

type FileItem struct {
	Name  string `json:"name"`
	IsDir bool   `json:"isDir"`
	Path  string `json:"path"`
}

func ListFilesHandler(basePath string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		queryPath := r.URL.Query().Get("path")
		fsPath := filepath.Join(basePath, queryPath)

		entries, err := os.ReadDir(fsPath)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		files := []FileItem{}
		for _, entry := range entries {
			files = append(files, FileItem{
				Name:  entry.Name(),
				IsDir: entry.IsDir(),
				Path:  filepath.Join(queryPath, entry.Name()),
			})
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(files)
	}
}
