package services

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	"log"
	"os"
	"report-backend-golang/global"
	// "report-backend-golang/services"
	// "report-backend-golang/models"
	"report-backend-golang/entities"

	// "strings"
	// "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	pdf "github.com/adrg/go-wkhtmltopdf"
)

type Product struct {
	Img         string
	Name        string
	Price       string
	Stars       float64
	Reviews     int
	Description string
}

type Dashboard struct {
	Img   string
	Name  string
	Price string
}

func subtr(a, b float64) float64 {
	return a - b
}

func list(e ...float64) []float64 {
	return e
}

func CreateHtml(ScheduleID int, ReportID int) {

	// data := []Product{
	// 	{"pics/RlgkWFg4k.png", "strawberries", "$2.00", 4.0, 251, "Lorem ipsum dolor sit amet, consectetur adipiscing elit."},
	// 	{"pics/AuBDTJ6nz.png", "onions", "$2.80", 5.0, 123, "Morbi sit amet erat vitae purus consequat vehicula nec sit amet purus."},
	// 	{"pics/7adfa750-4c81-11e8-b3d7-01146121b73d.png", "tomatoes", "$3.10", 4.5, 235, "Curabitur tristique odio et nibh auctor, ut sollicitudin justo condimentum."},
	// }
	// fmt.Println(data)

	dashboardData, err := GetDashboardInReport(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	data1 := []Dashboard{}

	for _, dashboards := range dashboardData {
		uu := new(Dashboard)
		uu.Img = fmt.Sprintf("%s/%s.png", global.EnvConfig.Reportengine.PicturePath, dashboards.UID)
		uu.Name = fmt.Sprintf(dashboards.DashboardName)
		data1 = append(data1, *uu)
	}
	fmt.Println(data1)

	//	用 ReportID 取出 Report 的相關資料
	inventory, err := GetReportByReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(inventory.Name)

	fromtimeconverted := Timeconverter(inventory.From + "+8h")
	totimeconverted := Timeconverter(inventory.To + "+8h")
	str1 := fromtimeconverted[0:16]
	str2 := totimeconverted[0:16]

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, "/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/template/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	if err := templates.ExecuteTemplate(&processed, "page", data1); err != nil {
		fmt.Println(err.Error())
	}
	outputPath := fmt.Sprintf("%s/%s_%s~%s.html", global.EnvConfig.Reportengine.HtmlPath, inventory.Name, str1, str2)
	htmlName := fmt.Sprintf("%s_%s~%s", inventory.Name, str1, str2)
	// outputPath := "/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/index.html"

	// 新增 html history
	Htmlfile := entities.FileHistory{ScheduleID:ScheduleID, HtmlName: htmlName,Filetype: "html"}
	global.Mysql.Create(&Htmlfile)

	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

	// pdf的名稱
	pdfname := fmt.Sprintf("%s/%s_%s~%s.pdf", global.EnvConfig.Reportengine.PdfPath, inventory.Name, str1, str2)
	pdfshortname := fmt.Sprintf("%s",htmlName)

	// 執行產出pdf的程式
	GeneratePDF(outputPath, pdfname)

	// 新增 pdf history
	file := entities.FileHistory{ScheduleID:ScheduleID,PdfName: pdfshortname,Filetype: "pdf" }
	global.Mysql.Create(&file)

}

// 新增 history
func CreateFileHistory() {
	// res := models.Response{}
	file := entities.FileHistory{ScheduleID: 1}
	global.Mysql.Create(&file)
}

func GeneratePDF(htmlpath string, pdfname string) {

	pdf.Init()
	defer pdf.Destroy()

	// Create object from file.
	// object, err := pdf.NewObject("/Users/chen/Documents/gitlab/git-out/product/report-backend/pdf/index.html")
	object, err := pdf.NewObject(htmlpath)
	if err != nil {
		log.Fatal(err)
	}
	object.Header.ContentCenter = "[title]"
	// object.Header.DisplaySeparator = true
	object.Header.DisplaySeparator = false

	// Create converter.
	converter, err := pdf.NewConverter()
	if err != nil {
		log.Fatal(err)
	}
	defer converter.Destroy()

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
		log.Fatal(err)
	}
	defer outFile.Close()

	if err := converter.Run(outFile); err != nil {
		log.Fatal(err)
	}
}
