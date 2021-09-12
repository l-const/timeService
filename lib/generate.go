package lib

func GenerateTimeStamps(t1Str string, t2Str string, locStr string, period interface{}) ([]string, error) {
	var err error
	t1, err := intoTimeFormat(t1Str, locStr)
	if err != nil {
		return []string{}, err
	}
	t2, err := intoTimeFormat(t2Str, locStr)
	if err != nil {
		return []string{}, err
	}
	times, err := fromTimeToTime(t1, t2, period)
	if err != nil {
		return []string{}, err
	}
	timesStr := timeSliceToStrings(times)
	return timesStr, err
}
