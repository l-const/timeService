package lib

import (
	"errors"
	"time"
)

func addDay(t time.Time) time.Time {
	return t.AddDate(0, 0, 1)
}

func addYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
}

func addHour(t time.Time) time.Time {
	return t.Add(time.Hour)
}

func addMonth(t time.Time) time.Time {
	return t.AddDate(0, 1, 0)
}

func AddTime(t time.Time, dur interface{}) (time.Time, error) {
	switch v := dur.(type) {
	case time.Duration:
		return t.Add(v), nil
	case string:
		if dur == "1h" {
			return addHour(t), nil
		} else if dur == "1d" {
			return addDay(t), nil
		} else if dur == "1mo" {
			return addMonth(t), nil
		} else if dur == "1y" {
			return addYear(t), nil
		} else {
			timeDur, err := time.ParseDuration(v)
			if err != nil {
				return time.Time{}, errors.New("Unsupported period")
			}
			return t.Add(timeDur), nil
		}
	default:
		return time.Time{}, errors.New("Unsupported period")
	}
}

func truncateOriginal(tm time.Time, dur interface{}) time.Time {
	var newtm time.Time
	switch v := dur.(type) {
	case string:
		if v == "1h" {
			newtm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour()+1, 0, 0, 0, time.UTC)
			return newtm
		} else if v == "1d" {
			return tm
		} else if v == "1mo" {
			newtm = tm.AddDate(0, 1, -tm.Day())
			return newtm
		} else if v == "1y" {
			newtm = time.Date(tm.Year(), 12, 31, 22, 0, 0, 0, time.UTC)
			return newtm
		} else {
			return tm
		}
	default:
		return tm
	}
}

func fromTimeToTime(from time.Time, to time.Time, dur interface{}) ([]time.Time, error) {
	var err error
	timeSlice := make([]time.Time, 0)
	initialT := truncateOriginal(from, dur)
	curT := initialT
	for !curT.After(to) {
		timeSlice = append(timeSlice, curT)
		curT, err = AddTime(curT, dur)
	}

	return timeSlice, err
}
