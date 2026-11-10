package main

import (
	"fmt"
	"time"
)

func main() {
	var now time.Time = time.Now()
	fmt.Println(now.Local())

	var utc time.Time = time.Date(2026, time.August, 17, 0,0,0,0, time.UTC)
	fmt.Println(utc.Local())

	// format golang time
	formatter := "2006-01-02 15:04:05"

	value := "2026-08-17 10:00:00"
	valueTime, err := time.Parse(formatter, value)
	if err != nil {
		fmt.Println(err.Error())
		} else {
		fmt.Println(valueTime)
	}

	fmt.Println(valueTime.Year())
	fmt.Println(valueTime.Month())
	fmt.Println(valueTime.Day())
}