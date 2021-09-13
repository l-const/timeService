package lib

import "time"

func GenerateTimeStamps(t1Str string, t2Str string, locStr string, period interface{}) ([]string, error) {
	var err error
	location, err := time.LoadLocation(locStr)
	t1, err := intoTimeFormat(t1Str, locStr)
	if err != nil {
		return []string{}, err
	}
	t2, err := intoTimeFormat(t2Str, locStr)
	if err != nil {
		return []string{}, err
	}
	times, err := fromTimeToTime(t1, t2, period, location)
	if err != nil {
		return []string{}, err
	}
	timesStr := timeSliceToStrings(times)
	return timesStr, err
}
