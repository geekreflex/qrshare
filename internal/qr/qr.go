package qr

import (
	qrcodeTerminal "github.com/Baozisoftware/qrcode-terminal-go"
)

func PrintQRCode(url string) {
	qr := qrcodeTerminal.New()
	qr.Get(url).Print()
}
