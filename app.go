package main

import (
	"context"
	"dev-tools/utils/base64"
	"dev-tools/utils/qrcode"
	"dev-tools/utils/url_code"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// GenerateQRCode .
func (a *App) GenerateQRCode(txt string) string {
	return qrcode.Generate(txt)
}

// Base64Encode .
func (a *App) Base64Encode(txt string) string {
	return base64.Encode(txt)
}

// Base64Decode .
func (a *App) Base64Decode(txt string) string {
	return base64.Decode(txt)
}

// UrlEncode .
func (a *App) UrlEncode(txt string) string {
	return url_code.Encode(txt)
}

// UrlDecode .
func (a *App) UrlDecode(txt string) string {
	return url_code.Decode(txt)
}
