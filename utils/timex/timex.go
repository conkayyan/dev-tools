package timex

import (
	"fmt"
	"os"
	"time"
)

type (
	Format struct {
		Label string `json:"label"`
		Value string `json:"value"`
	}
	TimeOption struct {
		Timestamp int64    `json:"timestamp"`
		Location  string   `json:"location"`
		Layout    string   `json:"layout"`
		Datetime  string   `json:"datetime"`
		Formats   []Format `json:"formats"`
	}
)

var (
	zoneInfoDataTmpFile string
	timeOption          TimeOption
)

func TimeInit() TimeOption {
	timeOption.Location = time.Now().Location().String()
	timeOption.Layout = time.DateTime
	timeOption.Formats = append(
		timeOption.Formats,
		Format{
			Label: fmt.Sprintf("%s(%s)", "Layout", time.Layout),
			Value: time.Layout,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "ANSIC", time.ANSIC),
			Value: time.ANSIC,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "UnixDate", time.UnixDate),
			Value: time.UnixDate,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RubyDate", time.RubyDate),
			Value: time.RubyDate,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC822", time.RFC822),
			Value: time.RFC822,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC822Z", time.RFC822Z),
			Value: time.RFC822Z,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC850", time.RFC850),
			Value: time.RFC850,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC1123", time.RFC1123),
			Value: time.RFC1123,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC1123Z", time.RFC1123Z),
			Value: time.RFC1123Z,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC3339", time.RFC3339),
			Value: time.RFC3339,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "RFC3339Nano", time.RFC3339Nano),
			Value: time.RFC3339Nano,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "Kitchen", time.Kitchen),
			Value: time.Kitchen,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "Stamp", time.Stamp),
			Value: time.Stamp,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "StampMilli", time.StampMilli),
			Value: time.StampMilli,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "StampMicro", time.StampMicro),
			Value: time.StampMicro,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "StampNano", time.StampNano),
			Value: time.StampNano,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "DateTime", time.DateTime),
			Value: time.DateTime,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "DateOnly", time.DateOnly),
			Value: time.DateOnly,
		},
		Format{
			Label: fmt.Sprintf("%s(%s)", "TimeOnly", time.TimeOnly),
			Value: time.TimeOnly,
		},
	)
	timeOption.Timestamp = time.Now().Unix()

	tm := time.Unix(timeOption.Timestamp, 0)
	location, err := time.LoadLocation(timeOption.Location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return timeOption
	}
	tm = tm.In(location)
	timeOption.Datetime = tm.Format(timeOption.Layout)
	return timeOption
}

func GetTimeOption() TimeOption {
	return timeOption
}

func SetUnix(timestamp int64) {
	timeOption.Timestamp = timestamp

	tm := time.Unix(timeOption.Timestamp, 0)
	location, err := time.LoadLocation(timeOption.Location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return
	}
	tm = tm.In(location)
	timeOption.Datetime = tm.Format(timeOption.Layout)
}

func SetLocation(timezone string) {
	timeOption.Location = timezone

	tm := time.Unix(timeOption.Timestamp, 0)
	location, err := time.LoadLocation(timeOption.Location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return
	}
	tm = tm.In(location)
	timeOption.Datetime = tm.Format(timeOption.Layout)
}

func SetFormat(layout string) {
	timeOption.Layout = layout

	tm := time.Unix(timeOption.Timestamp, 0)
	location, err := time.LoadLocation(timeOption.Location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return
	}
	tm = tm.In(location)
	timeOption.Datetime = tm.Format(timeOption.Layout)
}

func SetDatetime(datetime string) {
	timeOption.Datetime = datetime

	location, err := time.LoadLocation(timeOption.Location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return
	}

	tm, err := time.ParseInLocation(timeOption.Layout, timeOption.Datetime, location)
	if err != nil {
		timeOption.Datetime = err.Error()
		return
	}

	timeOption.Timestamp = tm.Unix()
}

func DeleteTmpFile() {
	if zoneInfoDataTmpFile != "" {
		_ = os.Remove(zoneInfoDataTmpFile)
	}
}
