package main

import (
	"encoding/json"
	"os"
	"timeservice/lib"
)

func main() {
	locStr := "Europe/Athens"
	var err error
	time1, time2, dur := "20211030T204603Z", "20211031T123456Z", "1h"
	response, err := lib.GenerateTimeStamps(time1, time2, locStr, dur)
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")

	if err != nil {
		panic(err)
	}
	enc.Encode(response)

	time1, time2, durD := "20211010T204603Z", "20211115T123456Z", "1d"
	response, err = lib.GenerateTimeStamps(time1, time2, locStr, durD)
	if err != nil {
		panic(err)
	}
	enc.Encode(response)

	time1, time2, durM := "20210214T204603Z", "20211115T123456Z", "1mo"
	response, err = lib.GenerateTimeStamps(time1, time2, locStr, durM)
	if err != nil {
		panic(err)
	}
	enc.Encode(response)

	time1, time2, durY := "20180214T204603Z", "20211115T123456Z", "1y"
	response, err = lib.GenerateTimeStamps(time1, time2, locStr, durY)
	if err != nil {
		panic(err)
	}
	enc.Encode(response)

}
