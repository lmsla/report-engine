package tools

import (
	"fmt"
	"image/png"
	"os"
	// "report-backend-golang/entities"
	"time"
)

func Timeconverter(nowtime int64, time_unit string, time_period int, alias string) string {
	// now := time.Now()
	t := time.Unix(nowtime, 0)
	var timeconverted string
	if alias == "" {
		switch time_unit {
		case "day":
			timeconverted = t.AddDate(0, 0, -time_period).Format("2006-01-02")
		case "week":
			timeconverted = t.AddDate(0, 0, -time_period*7).Format("2006-01-02")
		case "month":
			timeconverted = t.AddDate(0, -time_period, 0).Format("2006-01-02")
		case "year":
			timeconverted = t.AddDate(-time_period, 0, 0).Format("2006-01-02")
		}

	} else if alias != "" {
		switch alias {
		case "last_one_day":
			timeconverted = t.AddDate(0, 0, -1).Format("2006-01-02")
		case "last_seven_day":
			timeconverted = t.AddDate(0, 0, -7).Format("2006-01-02")
		case "last_fifteen_day":
			timeconverted = t.AddDate(0, 0, -15).Format("2006-01-02")
		case "last_month":
			timeconverted = t.AddDate(0, -1, 0).Format("2006-01-02")
		case "last_quarter":
			timeconverted = t.AddDate(0, -3, 0).Format("2006-01-02")
		case "last_six_month":
			timeconverted = t.AddDate(0, -6, 0).Format("2006-01-02")
		case "last_one_year":
			timeconverted = t.AddDate(-1, 0, 0).Format("2006-01-02")
		}

	}

	// fmt.Println(timeconverted)
	return timeconverted
}

func Toools() {
	now := time.Now().Format("2006-01-02")
	fmt.Println(now)
}

func GetImageHW(image string) (width int, height int) {
	// Open PNG 
	file, err := os.Open(image)
	if err != nil {
		fmt.Println("無法打開文件:", err)
		return
	}
	defer file.Close()

	// 解碼 PNG 圖片
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("無法解碼圖片:", err)
		return
	}

	// 取得圖片的長寬
	width = img.Bounds().Dx()
	height = img.Bounds().Dy()

	// fmt.Println("宽度:", width)
	// fmt.Println("高度:", height)

	return width, height
}
