package server

import (
	"net/http"

	"github.com/geekreflex/qrshare/internal/files"
)

func RegisterRoutes(mux *http.ServeMux, basePath string, distPath string) {
	// Serve React UI
	mux.Handle("/", http.FileServer(http.Dir(distPath)))

	// Serve file metadata (API)
	mux.Handle("/api/files", files.ListFilesHandler(basePath))

	// Serve actual files for download
	mux.Handle("/files/", http.StripPrefix("/files/", http.FileServer(http.Dir(basePath))))
}
