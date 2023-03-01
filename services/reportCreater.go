package services

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	log1 "log"
	"os"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"report-backend-golang/tools"
	"time"

	pdf "github.com/adrg/go-wkhtmltopdf"
)

func subtr(a, b float64) float64 {
	return a - b
}

func list(e ...float64) []float64 {
	return e
}

type Report struct {
	Name     string
	Elements []Element
}

type Element struct {
	Img   string
	Name  string
	Period string
}

type title struct {
	Name string
}

type dashboard struct {
	Img   string
	Name  string
	Price string
}

func CreateHtmlbySchedule(ScheduleID int) {
	scheduleData, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
	}
	for _, reports := range scheduleData {
		// log.Logrecord("排程","report name: "+reports.Name+" 開始產出")
		CreateHtml(reports.ID)
		// log.Logrecord("排程","report name: "+reports.Name+" 完成產出")
	}

}

func CreatePDFbySchedule(ScheduleID int) {
	scheduleData, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)
	for _, reports := range scheduleData {
		fmt.Println(reports.Name)
		timefrom := tools.Timeconverter(reports.TimeUnit, reports.TimePeriod)
		now := time.Now().Format("2006-01-02")
		outputPath := fmt.Sprintf("%s/%s_%s_%s.html", global.EnvConfig.Files.HtmlFile, reports.Name, timefrom, now)
		pdfname := fmt.Sprintf("%s/%s_%s_%s.pdf", global.EnvConfig.Files.ReportFile, reports.Name, timefrom, now)
		// log.Logrecord("排程","report name: "+reports.Name+" 開始產出")
		// defer pdf.Destroy()
		GeneratePDF(outputPath, pdfname)
		// log.Logrecord("排程","report name: "+reports.Name+" 完成產出")
		// time.Sleep(5 * time.Second)
		// pdf.Destroy()
		// defer pdf.Destroy()
	}
	pdf.Destroy()
	// defer pdf.Destroy()

}

func CreateHtml(ReportId int) {

	// ScreenshotbyReport(ReportId)

	//	用 ReportID 取出 Report 的相關資料
	report_data, err := GetReportByReportID(ReportId)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(report_data.Name)

	element_data, err := GetElementsByReportID(ReportId)
	if err != nil {
		fmt.Println(err)
	}

	now := time.Now().Format("2006-01-02")
	timefrom := tools.Timeconverter(report_data.TimeUnit, report_data.TimePeriod)
	//----------------------------------------
	// data1 := []dashboard{}
	// for _, elements := range elementData {
	// 	uu := new(dashboard)
	// 	uu.Img = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID,timefrom,now)
	// 	uu.Name = fmt.Sprintf(elements.Name)
	// 	data1 = append(data1, *uu)
	// }
	// fmt.Println(data1)
	//----------------------------------------

	data1 := Report{}
	uu := new(Report)
	gg := new(Element)
	uu.Name = fmt.Sprintf(report_data.Name)
	for _, elements := range element_data {
		// uu := new(Report)
		gg.Img = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID, timefrom, now)
		gg.Name = fmt.Sprintf(elements.Name)
		gg.Period = fmt.Sprintf(timefrom+"~"+now)
		// uu.Elements = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID,timefrom,now)
		data1.Elements = append(data1.Elements, *gg)
	}

	data1 = Report{
		Name: uu.Name,
		Elements:data1.Elements ,
	}
	log.Logrecord("test", "testlog")
	fmt.Println(data1.Name)
	fmt.Println(data1.Elements)

	// fromtimeconverted := Timeconverter(inventory.From + "+8h")
	// totimeconverted := Timeconverter(inventory.To + "+8h")
	// str1 := fromtimeconverted[0:16]
	// str2 := totimeconverted[0:16]

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "/Users/chen/Documents/gitlab/git-out/product/report-backend/files/template/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	if err := templates.ExecuteTemplate(&processed, "page", data1); err != nil {
		fmt.Println(err.Error())
	}
	// outputPath := fmt.Sprintf("%s/%s_%s~%s.html", global.EnvConfig.Files.HtmlFile, inventory.Name, str1, str2)
	outputPath := fmt.Sprintf("%s/%s_%s_%s.html", global.EnvConfig.Files.HtmlFile, report_data.Name, timefrom, now)
	// htmlName := fmt.Sprintf("%s_%s~%s", inventory.Name, str1, str2)
	// htmlName := fmt.Sprintf("%s_%s~%s", inventory.Name)
	// outputPath := "/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/index.html"

	// 新增 html history
	// Htmlfile := entities.FileHistory{ScheduleID:ScheduleID, HtmlName: htmlName,Filetype: "html"}
	// global.Mysql.Create(&Htmlfile)

	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

	// pdf的名稱
	// pdfname := fmt.Sprintf("%s/%s_%s~%s.pdf", global.EnvConfig.Files.ReportFile, inventory.Name, str1, str2)

	// pdfname := fmt.Sprintf("%s/%s.pdf", global.EnvConfig.Files.ReportFile, inventory.Name)

	// pdfshortname := fmt.Sprintf("%s",htmlName)
	// time.Sleep(5 * time.Second)
	// 執行產出pdf的程式

	// GeneratePDF(outputPath, pdfname)

	// var wg sync.WaitGroup
	// go func() {
	// 	wg.Add(1)//計數器+1
	// 	defer wg.Done()

	// }()

	// time.Sleep(time.Millisecond * 30)//休息30 ms
	// log.Println("wait a goroutine")
	// wg.Wait()//等待計數器歸0

	// // 新增 pdf history
	// file := entities.FileHistory{ScheduleID:ScheduleID,PdfName: pdfshortname,Filetype: "pdf" }
	// global.Mysql.Create(&file)

}

func GeneratePDF(htmlpath string, pdfname string) {

	pdf.Init()
	// defer pdf.Destroy()

	// Create object from file.
	// object, err := pdf.NewObject("/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/index.html")
	object, err := pdf.NewObject(htmlpath)
	if err != nil {
		log1.Fatal(err)
	}
	object.Header.ContentCenter = "[title]"
	// object.Header.DisplaySeparator = true
	object.Header.DisplaySeparator = false

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log1.Fatal(err)
	}
	// defer converter.Destroy()

	// Add created objects to the converter.
	converter.Add(object)
	// converter.Add(object2)
	// converter.Add(object3)

	// Set converter options.
	converter.Title = "Sample document"
	converter.PaperSize = pdf.A4
	//橫向展示
	// converter.Orientation = pdf.Landscape
	converter.MarginTop = "1cm"
	converter.MarginBottom = "1cm"
	converter.MarginLeft = "10mm"
	converter.MarginRight = "10mm"

	// Convert objects and save the output PDF document.
	outFile, err := os.Create(pdfname)
	if err != nil {
		log1.Fatal(err)
	}
	// defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		log1.Fatal(err)
	}
	converter.Destroy()
	outFile.Close()
}
