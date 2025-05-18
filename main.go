package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"path/filepath"

	// "github.com/geekreflex/qrshare/internal/api"
	"github.com/geekreflex/qrshare/internal/qr"
	"github.com/geekreflex/qrshare/internal/server"
	"github.com/geekreflex/qrshare/internal/utils"
)

func main() {
	port := 3000
	basePath, _ := os.UserHomeDir()
	distPath := "./web/dist"

	mux := http.NewServeMux()

	// API routes
	// mux.HandleFunc("/api/files/list", api.HandleListFiles)
	// mux.HandleFunc("/api/files/raw", api.HandleRawFile)

	server.RegisterRoutes(mux, basePath)

	// Serve React UI if built
	if _, err := os.Stat(filepath.Join(distPath, "index.html")); err == nil {
		fs := http.FileServer(http.Dir(distPath))
		mux.Handle("/", fs)
	}

	// Start the server
	ip := utils.GetLocalIP() // your own util func that gets 192.168.x.x
	url := fmt.Sprintf("http://%s:%d", ip, port)
	fmt.Printf("📡 Serving at: %s\n", url)

	// Print QR code
	qr.PrintQRCode(url)
	fmt.Println("Scan the QR to access files:", url)

	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), mux))
}
