package server

import (
	"net/http"

	"github.com/geekreflex/qrshare/internal/files"
)

func RegisterRoutes(mux *http.ServeMux, basePath string) {
	// Register API routes under /api
	mux.Handle("/api/files/list", http.HandlerFunc(files.ListFilesHandler(basePath)))

	// Serve files from local filesystem (for downloads, previews, etc.)
	mux.Handle("/api/files/raw/", http.StripPrefix("/api/files/raw/", http.FileServer(http.Dir(basePath))))
}
