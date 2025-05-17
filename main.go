package main

import (
	"fmt"
	"os"

	"github.com/geekreflex/qrshare/internal/qr"
	"github.com/geekreflex/qrshare/internal/server"
	"github.com/geekreflex/qrshare/internal/utils"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: qrshare <folder-to-share>")
		return
	}
	folder := os.Args[1]
	ip := utils.GetLocalIP()
	url := fmt.Sprintf("http://%s:8080", ip)

	go server.ServeFiles(folder)
	qr.PrintQRCode(url)

	fmt.Println("Scan the QR to access files:", url)

	select {} // Keep running
}
