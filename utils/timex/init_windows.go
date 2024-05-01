package timex

import (
	"embed"
	"log"
	"os"
)

var (
	//go:embed tzdata.zip
	zoneInfoDataFS embed.FS
)

func init() {
	// read tzdata.zip from embed.FS
	zoneInfoData, err := zoneInfoDataFS.ReadFile("tzdata.zip")
	if err != nil {
		log.Fatal(err.Error())
	}

	// write tzdata.zip to temp file
	tempFile, err := os.CreateTemp("", "tzdata.zip")
	if err != nil {
		log.Fatal(err)
	}
	_, err = tempFile.Write(zoneInfoData)
	if err != nil {
		log.Fatal(err)
	}

	zoneInfoDataTmpFile = tempFile.Name()
	_ = os.Setenv("ZONEINFO", tempFile.Name())
}
