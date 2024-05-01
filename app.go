package main

import (
	"context"
	"dev-tools/utils/base64"
	"dev-tools/utils/md5"
	"dev-tools/utils/qrcode"
	"dev-tools/utils/str_length"
	"dev-tools/utils/timex"
	"dev-tools/utils/url_code"
	"dev-tools/utils/uuid"
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
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	timex.DeleteTmpFile()
	return false
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

// Md5Encode .
func (a *App) Md5Encode(txt string) string {
	return md5.Encode(txt)
}

// StrLength .
func (a *App) StrLength(txt string) int {
	return str_length.Count(txt)
}

// TimeInit .
func (a *App) TimeInit() timex.TimeOption {
	return timex.TimeInit()
}

// GetTimeOption .
func (a *App) GetTimeOption() timex.TimeOption {
	return timex.GetTimeOption()
}

// SetUnix .
func (a *App) SetUnix(timestamp int64) {
	timex.SetUnix(timestamp)
}

// SetLocation .
func (a *App) SetLocation(timezone string) {
	timex.SetLocation(timezone)
}

// SetFormat .
func (a *App) SetFormat(layout string) {
	timex.SetFormat(layout)
}

// SetDatetime .
func (a *App) SetDatetime(datetime string) {
	timex.SetDatetime(datetime)
}

// GetUUID .
func (a *App) GetUUID(isHyphen bool) string {
	return uuid.GenUUID(isHyphen)
}
