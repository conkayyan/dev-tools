package main

import (
	"context"
	"dev-tools/utils/qrcode"
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
