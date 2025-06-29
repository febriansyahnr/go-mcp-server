package util

import (
	"log"
	"time"
	_ "time/tzdata"

	"github.com/paper-indonesia/pg-mcp-server/constant"
)

type LocationLoader func(name string) (*time.Location, error)

func GetJakartaTimeWithLoader(loader LocationLoader) (time.Time, error) {
	t, err := loader("Asia/Jakarta")
	if err != nil {
		return time.Time{}, err
	}

	return time.Now().In(t), nil
}

func GetJakartaTime() (time.Time, error) {
	return GetJakartaTimeWithLoader(time.LoadLocation)
}

func SnapFormat(oldTime time.Time) string {

	t, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("error when load location", err)
		return ""
	}

	convertedTime := oldTime.In(t)

	parsedTime, err := time.Parse(constant.SnapDateFormatLayout, convertedTime.Format(constant.SnapDateFormatLayout))
	if err != nil {
		log.Println("error when parse time", err)
		return ""
	}

	return parsedTime.Format(constant.SnapDateFormatLayout)
}

func CustomFormat(oldTime time.Time, layout string) string {
	t, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("error when load location", err)
		return ""
	}

	convertedTime := oldTime.In(t)

	parsedTime, err := time.Parse(layout, convertedTime.Format(layout))
	if err != nil {
		log.Println("error when parse time", err)
		return ""
	}

	return parsedTime.Format(layout)
}

// C2aDateFormat is a function that convert time to c2a format
// Result: YYYYMMDD, HHMMSS
func C2aDateFormat(oldTime time.Time) (string, string) {
	dateLayout := "20060102"
	timeLayout := "150405"
	t, err := time.LoadLocation("Asia/Jakarta")
	if err != nil {
		log.Println("error when load location", err)
		return "", ""
	}

	convertedTime := oldTime.In(t)
	return convertedTime.Format(dateLayout), convertedTime.Format(timeLayout)
}

// CalculateToMidnight calculate duration from `now` to midnight in `location`
func CalculateToMidnight(now *time.Time, location string) time.Duration {
	tzLoc, _ := time.LoadLocation(location)
	nowLoc := now.In(tzLoc)
	midnight := time.Date(nowLoc.Year(), nowLoc.Month(), nowLoc.Day(), 0, 0, 0, 0, tzLoc)

	return midnight.Add(time.Hour * 24).Sub(nowLoc)
}

// TimeFormat is convert time to string with custom layout to Time.time format
func TimeFormat(dateStr string, layout string) time.Time {
	// Parse time using given layout, if error try to parse using another layout
	// 1. Parse time using given layout
	// 2. Parse time using snap date format layout
	// 3. Parse time using date format without time layout

	layouts := []string{layout, constant.SnapDateFormatLayout, time.DateTime, time.DateOnly}

	var parsedTime time.Time
	var err error
	for _, l := range layouts {
		parsedTime, err = time.Parse(l, dateStr)
		if err == nil {
			break
		}
	}

	return parsedTime
}
