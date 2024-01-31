package tools

import (
	"fmt"
	"image/png"
	"os"
	"time"

)

func Timeconverter(nowtime int64,time_unit string, time_period int) string {
	// now := time.Now()
	t := time.Unix(nowtime, 0)
	var timeconverted string
	switch time_unit {
	case "日":
		timeconverted = t.AddDate(0, 0, -time_period).Format("2006-01-02")
	case "週":
		timeconverted = t.AddDate(0, 0, -time_period*7).Format("2006-01-02")
	case "月":
		timeconverted = t.AddDate(0, -time_period, 0).Format("2006-01-02")
	case "年":
		timeconverted = t.AddDate(-time_period, 0, 0).Format("2006-01-02")
	}
	// fmt.Println(timeconverted)
	return timeconverted
}

func Toools() {
	now := time.Now().Format("2006-01-02")
	fmt.Println(now)
}



func GetImageHW(image string) (width int ,height int){
	// 打开 PNG 文件
	file, err := os.Open(image)
	if err != nil {
		fmt.Println("无法打开文件:", err)
		return
	}
	defer file.Close()

	// 解码 PNG 图像
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println("无法解码图像:", err)
		return
	}

	// 获取图像的长宽
	width = img.Bounds().Dx()
	height = img.Bounds().Dy()

	// fmt.Println("宽度:", width)
	// fmt.Println("高度:", height)

	return width , height
}