package timex

import (
	"embed"
	"log"
	"os"
)

var (
	zoneInfoDataTempFile string
	//go:embed zoneinfo.zip
	zoneInfoDataFS embed.FS
)

func init() {
	// read zoneinfo.zip from embed.FS
	zoneInfoData, err := zoneInfoDataFS.ReadFile("zoneinfo.zip")
	if err != nil {
		log.Fatal(err.Error())
	}

	// mkdir temp dir
	zoneInfoDataTempDir, err = os.MkdirTemp(os.TempDir(), "zoneinfo")
	if err != nil {
		log.Fatal(err)
	}

	// write zoneinfo.zip to temp dir
	zoneInfoDataTempFile = filepath.Join(zoneInfoDataTempDir, "zoneinfo.zip")
	err = os.WriteFile(zoneInfoDataTempFile, zoneInfoData, 0666)
	if err != nil {
		log.Fatal(err)
	}

	// set ZONEINFO env
	_ = os.Setenv("ZONEINFO", zoneInfoDataTempFile)
}
