package lib

import (
	"errors"
	"fmt"
	"time"
)

func addHour(t time.Time) time.Time {
	newtm := time.Date(t.Year(), t.Month(), t.Day(), t.Hour()+1, 0, 0, 0, t.Location())
	if t.IsDST() && !newtm.IsDST() {
		newtm = newtm.Add(time.Hour)
	}
	if !t.IsDST() && newtm.IsDST() {
		newtm = newtm.Add(-time.Hour)
	}
	return newtm
}

func addDay(t time.Time) time.Time {
	newtm := t.AddDate(0, 0, 1)
	if t.IsDST() && !newtm.IsDST() {
		newtm = newtm.Add(time.Hour)
	}
	if !t.IsDST() && newtm.IsDST() {
		newtm = newtm.Add(-time.Hour)
	}
	return newtm

}

func addMonth(t time.Time) time.Time {
	newtm := time.Date(t.Year(), t.Month()+1, 1, t.Hour(), 0, 0, 0, t.Location())
	newtm = newtm.AddDate(0, 1, -1)

	if t.IsDST() && !newtm.IsDST() {
		newtm = newtm.Add(time.Hour)
	}

	if !t.IsDST() && newtm.IsDST() {
		newtm = newtm.Add(-time.Hour)
	}

	return newtm
}

func addYear(t time.Time) time.Time {
	return t.AddDate(1, 0, 0)
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
			return time.Time{}, errors.New("Unsupported period")
		}
	default:
		return time.Time{}, errors.New("Unsupported period")
	}
}

func truncateInit(tm time.Time, dur interface{}) time.Time {
	var newtm time.Time
	switch v := dur.(type) {
	case string:
		if v == "1h" {
			newtm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour()+1, 0, 0, 0, tm.Location())
			return newtm
		} else if v == "1d" {
			newtm = time.Date(tm.Year(), tm.Month(), tm.Day(), tm.Hour()+1, 0, 0, 0, tm.Location())
			return newtm
		} else if v == "1mo" {
			fmt.Println("month is daylight sacing time : ", tm.IsDST())
			newtm = time.Date(tm.Year(), tm.Month(), 1, tm.Hour()+1, 0, 0, 0, tm.Location())
			newtm = newtm.AddDate(0, 1, -1)
			if !tm.IsDST() && newtm.IsDST() {
				newtm = newtm.Add(time.Hour * 2)
			}
			return newtm
		} else if v == "1y" {
			if tm.IsDST() {
				newtm = time.Date(tm.Year(), 12, 31, tm.Hour()+1, 0, 0, 0, tm.Location())
			} else {
				newtm = time.Date(tm.Year(), 12, 31, tm.Hour()+2, 0, 0, 0, tm.Location())
			}
			return newtm
		} else {
			return tm
		}
	default:
		return tm
	}
}

func fromTimeToTime(from time.Time, to time.Time, dur interface{}, loc *time.Location) ([]time.Time, error) {
	var err error
	timeSlice := make([]time.Time, 0)
	initialT := truncateInit(from, dur)
	curT := initialT
	for !curT.After(to) {
		timeSlice = append(timeSlice, curT)
		curT, err = AddTime(curT, dur)
		if err != nil {
			return []time.Time{}, err
		}
	}
	return timeSlice, err
}
