package lib

import (
	"errors"
	"strconv"
	"strings"
	"time"
)

func timeSliceToStrings(times []time.Time) []string {
	strSlice := make([]string, 0)
	for _, v := range times {
		strSlice = append(strSlice, intoString(v))
	}
	return strSlice
}

func intoString(tm time.Time) string {
	tmFor := tm.Format(time.RFC3339)
	tmFor = strings.ReplaceAll(tmFor, "-", "")
	tmFor = strings.ReplaceAll(tmFor, ":", "")[:15] + "Z"
	return tmFor
}

func intoTimeFormat(timeString string, locStr string) (time.Time, error) {
	location, err := time.LoadLocation(locStr)
	if err != nil {
		return time.Time{}, err
	}
	if len(timeString) != 16 {
		return time.Time{}, errors.New("wrong format for string timestamp")
	}
	year, err := strconv.Atoi(timeString[0:4])
	if err != nil {
		return time.Time{}, err
	}
	month, err := strconv.Atoi(timeString[4:6])
	if err != nil {
		return time.Time{}, err
	}
	day, err := strconv.Atoi(timeString[6:8])
	if err != nil {
		return time.Time{}, err
	}
	hour, err := strconv.Atoi(timeString[9:11])
	if err != nil {
		return time.Time{}, err
	}
	minutes, err := strconv.Atoi(timeString[11:13])
	if err != nil {
		return time.Time{}, err
	}
	seconds, err := strconv.Atoi(timeString[13:15])
	date := time.Date(year, time.Month(month), day, hour, minutes, seconds, 0, location)
	return date, err

}
