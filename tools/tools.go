package tools

import (
	"fmt"
	"time"
)

func Timeconverter(time_unit string, time_period int) string {
	now := time.Now()
	var timeconverted string
	switch time_unit {
	case "日":
		timeconverted = now.AddDate(0, 0, -time_period).Format("2006-01-02")
	case "週":
		timeconverted = now.AddDate(0, 0, -time_period*7).Format("2006-01-02")
	case "月":
		timeconverted = now.AddDate(0, -time_period, 0).Format("2006-01-02")
	case "年":
		timeconverted = now.AddDate(-time_period, 0, 0).Format("2006-01-02")
	}
	// fmt.Println(timeconverted)
	return timeconverted
}

func Toools() {
	now := time.Now().Format("2006-01-02")
	fmt.Println(now)
}
