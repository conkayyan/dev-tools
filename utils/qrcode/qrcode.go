package qrcode

import (
	"encoding/base64"
	"fmt"
	qrcodex "github.com/skip2/go-qrcode"
)

func Generate(txt string) string {
	png, err := qrcodex.Encode(txt, qrcodex.Medium, 256)
	if err != nil {
		fmt.Println(err.Error())
		return ""
	}

	return fmt.Sprintf("data:image/png;base64,%s", base64.StdEncoding.EncodeToString(png))
}
