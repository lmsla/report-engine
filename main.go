package main

import (
	"fmt"
	"report-backend-golang/clients"
	"report-backend-golang/global"
	"report-backend-golang/migrate"
	"report-backend-golang/router"
	"report-backend-golang/services"
	pdf "github.com/adrg/go-wkhtmltopdf"
	// "report-backend-golang/log"
	// "report-backend-golang/tools"
	// "report-backend-golang/services"
	"report-backend-golang/utils"
	// "report-backend-golang/entities"
	"io/ioutil"
	"encoding/base64"
)

// @title Report Engine Golang API
// @version 1.0
// @description Golang API 專案描述
// @termsOfService http://swagger.io/terms/

// @contact.name Winston
// @contact.email support@swagger.io

//// @host 10.99.1.133:8005
// @host localhost:8005
// @BasePath  /api/v1

// @query.collection.format multi

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

// @schemes http
func main() {

	utils.LoadEnvironment()

	// es.LoadElasticsearch()

	clients.LoadDatabase()
	mysql, _ := global.Mysql.DB()
	defer mysql.Close()

	// if global.EnvConfig.Database.Migration {
	// 	entities.InitTable()
	// }

	migrate.Run()

	clients.LoadRedis()
	defer global.Redis.Close()

	utils.LoadCrontab()
	// // authorize.LoadCasbin()

	r := router.LoadRouter()
	r.Run(global.EnvConfig.Server.Port)

	// log.Logrecord("環境參數", "截圖檔案位置 : "+global.EnvConfig.Files.ScreenshotFile)
	// log.Logrecord("環境參數", "Html檔案位置 : "+global.EnvConfig.Files.HtmlFile)
	// log.Logrecord("環境參數", "PDF報表檔案位置 : "+global.EnvConfig.Files.ReportFile)
}

func main3() {
	utils.LoadEnvironment()
	// utils.LoadCrontab()
	htmlPaths := []string{"esinsight_report_part1_2024-01-18_2024-01-25", "esinsight_report_part2_2024-01-20_2024-01-25", "esinsight_report_part3_2024-01-15_2024-01-25"}


	// Initialize library.
	if err := pdf.Init(); err != nil {
		fmt.Println("init error")
	}

	defer pdf.Destroy()

	for _, path := range htmlPaths {

		htmlpath := fmt.Sprintf("/Users/chen/Downloads/personal_re/test/html_files/%s.html",path)
		pdfpath :=  fmt.Sprintf("/Users/chen/Downloads/personal_re/test/report_files/%s.pdf",path)
		// fmt.Println(htmlpath)
		// services.GeneratePDF_new(htmlpath,path)

		// services.GeneratePDF1(2000,htmlpath,pdfpath)

		services.GeneratePDF_by_Chromedp(htmlpath,pdfpath)
	}
	

}


func main2() {
		// 讀取圖片文件
		imagePath := "/Users/chen/Downloads/personal_re/test/html_files/3001c560-7949-11ee-992a-e1aa9b0ae3ae_2024-01-14_2024-01-24.png"
		imageData, err := ioutil.ReadFile(imagePath)
		if err != nil {
			fmt.Println("Error reading image file:", err)
			return
		}
	
		// 將圖片轉換為 base64 編碼
		base64Encoded := base64.StdEncoding.EncodeToString(imageData)
	
		// 將 base64 字符串輸出或用於其他地方
		fmt.Println("Base64 Encoded Image:")
		fmt.Println(base64Encoded)
}


func main1() {
	// services.GeneratePDF_by_gofpdf_No_seprate()
}