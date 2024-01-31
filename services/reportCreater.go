package services

import (
	"bufio"
	"bytes"
	"fmt"
	"html/template"
	log1 "log"
	"os"
	"report-backend-golang/entities"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"report-backend-golang/tools"
	"strings"
	"time"
	// "github.com/SebastiaanKlippert/go-wkhtmltopdf"
	// pdf "github.com/adrg/go-wkhtmltopdf"
	"github.com/jung-kurt/gofpdf"
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
	Img    string
	Name   string
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

func CreateHtmlbySchedule(nowtime int64, ScheduleID int) (err error) {

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			log.Logrecord("ERROR", "func CreateHtmlbySchedule error")
			// fmt.Println("func CreateHtmlbySchedule  發生錯誤：", err)
		}
	}()

	scheduleData, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
		log.Logrecord("ERROR ", "Get Report by Schedule ID error"+err.Error())
	}
	for _, reports := range scheduleData {
		// log.Logrecord("排程","report name: "+reports.Name+" 開始產出")
		// CreateHtml(reports.ID)
		// log.Logrecord("排程","report name: "+reports.Name+" 完成產出")
		err := CreateHtml(nowtime, reports.ID)
		if err != nil {
			fmt.Println("CreateHtml - line 78", err)
			log.Logrecord("ERROR", "CreateHtml error "+err.Error())
			return err
		}

	}
	return err
}

func CreateHtml(nowtime int64, ReportId int) (err error) {

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			log.Logrecord("ERROR", "func CreatePDFbySchedule error")
			// fmt.Println("func CreateHtmlbySchedule  發生錯誤：", err)
		}
	}()

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
	t := time.Unix(nowtime, 0)
	now := t.Format("2006-01-02")
	timefrom := tools.Timeconverter(nowtime, report_data.TimeUnit, report_data.TimePeriod, report_data.Alias)

	data1 := Report{}
	uu := new(Report)
	gg := new(Element)
	uu.Name = fmt.Sprintf(report_data.Name)
	for _, elements := range element_data {
		// uu := new(Report)
		gg.Img = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID, timefrom, now)
		gg.Name = fmt.Sprintf(elements.Name)
		gg.Period = fmt.Sprintf(timefrom + "~" + now)
		// uu.Elements = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID,timefrom,now)
		data1.Elements = append(data1.Elements, *gg)
	}

	data1 = Report{
		Name:     uu.Name,
		Elements: data1.Elements,
	}
	fmt.Println("CreateHtml-169")
	fmt.Println(data1.Name)
	fmt.Println(data1.Elements)

	allFiles := []string{"content.tmpl", "footer.tmpl", "header.tmpl", "page.tmpl"}

	var allPaths []string
	for _, tmpl := range allFiles {
		allPaths = append(allPaths, global.EnvConfig.Files.TemplateFile+"/"+tmpl)
	}

	templates := template.Must(template.New("").Funcs(template.FuncMap{"subtr": subtr, "list": list}).ParseFiles(allPaths...))

	var processed bytes.Buffer
	if err := templates.ExecuteTemplate(&processed, "page", data1); err != nil {
		fmt.Println(err.Error())
	}
	// outputPath := fmt.Sprintf("%s/%s_%s~%s.html", global.EnvConfig.Files.HtmlFile, inventory.Name, str1, str2)
	outputPath := fmt.Sprintf("%s/%s_%s_%s.html", global.EnvConfig.Files.HtmlFile, report_data.Name, timefrom, now)

	f, _ := os.Create(outputPath)
	w := bufio.NewWriter(f)
	w.WriteString(string(processed.Bytes()))
	w.Flush()

	return err
}

func CreatePDFbySchedule(nowtime int64, ScheduleID int) (err error) {

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			log.Logrecord("ERROR", "func CreatePDFbySchedule error")

		}
	}()
	scheduleData, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
	}
	time.Sleep(10 * time.Second)

	for _, reports := range scheduleData {
		// defer pdf.Destroy()
		fmt.Println(reports.Name)
		timefrom := tools.Timeconverter(nowtime, reports.TimeUnit, reports.TimePeriod, reports.Alias)
		t := time.Unix(nowtime, 0)
		now := t.Format("2006-01-02")

		// outputPath := fmt.Sprintf("%s/%s_%s_%s.html", global.EnvConfig.Files.HtmlFile, reports.Name, timefrom, now)

		log.Logrecord("排程", "report name: "+reports.Name+" 開始產出")

		//	用 ReportID 取出 Report 的相關資料
		report_data, err := GetReportByReportID(reports.ID)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(report_data.Name)
		element_data, err := GetElementsByReportID(reports.ID)
		if err != nil {
			fmt.Println(err)
		}

		timefrom = tools.Timeconverter(nowtime, report_data.TimeUnit, report_data.TimePeriod, report_data.Alias)

		GeneratePDF_by_gofpdf_No_seprate(element_data, reports.Name, timefrom, now)

		log.Logrecord("排程", "report name: "+reports.Name+" 完成產出")
	}

	return err
}

// 不分頁
func GeneratePDF_by_gofpdf_No_seprate(elementData []entities.Element, report_name string, timefrom string, now string) {
	pdfname := fmt.Sprintf("%s/%s_%s_%s.pdf", global.EnvConfig.Files.ReportFile, report_name, timefrom, now)

	fontSize := float64(14)
	// 創建 PDF 文件
	pdf := gofpdf.New("L", "mm", "A4", "")
	var width1, total_height float64
	for _, element := range elementData {
		image_name := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, element.UID, timefrom, now)
		width, height := tools.GetImageHW(image_name)
		width1 = (float64(width) / 40 * 4) + 20
		total_height += (float64(height) / 40 * 4) + 20
	}
	total_height = total_height + 27
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: width1, Ht: total_height})
	pdf.SetFontLocation("/src")
	pdf.AddUTF8Font("Taipei Sans TC Beta", "", "TaipeiSansTCBeta-Regular.ttf")

	// 報表名稱
	pdf.SetFont("Taipei Sans TC Beta", "", 20)
	pdf.CellFormat(0, 0, report_name, "", 0, "C", false, 0, "")

	// set logo
	pdf.Image(global.EnvConfig.Files.LogoFile, 5, 5, 20, 20, false, "", 0, "")

	// //给个空字符串就会去替换默认的 "{nb}"。
	// //如果这里指定了特别的字符串，那么SetFooterFunc() 中的 "nb" 也必须换成这个特别的字符串
	// pdf.AliasNbPages("")
	pdf.SetTopMargin(25)
	// Page properties
	pageWidth, pageHeight := pdf.GetPageSize()
	fmt.Println("pageWidth: ", pageWidth, "pageHeight: ", pageHeight)

	_, topMargin, _, _ := pdf.GetMargins()

	//Start from top of the page
	var y float64

	fmt.Println("topMargin", topMargin)
	y = topMargin
	for i, element := range elementData {

		fmt.Println("image:", i)

		image_name := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, element.UID, timefrom, now)
		width, height := tools.GetImageHW(image_name)

		image_width_cm := (float64(width) / 40) * 4

		image_height_cm := (float64(height) / 40) * 4

		fmt.Println("width", image_width_cm, "height", image_height_cm)

		// Add 圖表標題
		fmt.Println("text Y :", y)
		pdf.SetY(y)

		pdf.SetFont("Taipei Sans TC Beta", "", fontSize)
		pdf.CellFormat(0, 0, element.Name, "", 0, "C", false, 0, "")
		y = y + 5
		pdf.SetY(y)
		period := fmt.Sprintf("報表期間:%s~%s", timefrom, now)
		pdf.SetFont("Taipei Sans TC Beta", "", 10)
		pdf.CellFormat(0, 0, period, "", 0, "R", false, 0, "")

		y = y + 5
		pdf.Image(image_name, 10, y, image_width_cm, image_height_cm, false, "", 0, "")

		fmt.Println("Image Y", y)

		// Update Y coordinate
		y = y + image_height_cm + 15 // Adjust spacing as needed

		fmt.Println("y += imageHeight", y)

	}

	err := pdf.OutputFileAndClose(pdfname)
	if err != nil {
		log1.Fatal(err)
	}

	fmt.Println("PDF generated successfully.")

}

// 分頁
func GeneratePDF_by_gofpdf() {

	// 創建 PDF 文件
	pdf := gofpdf.New("P", "mm", "A4", "")

	//写文字内容之前，必须先要设置好字体
	pdf.SetFont("Arial", "B", 14)

	//设置页眉
	pdf.SetHeaderFuncMode(func() {
		pdf.Image("/Users/chen/Downloads/personal_re/test/screenshot_files/bimap.png", 0, 0, 20, 20, false, "", 0, "")
		pdf.SetY(5)
		// pdf.Ln(10)
	}, true)

	//设置页脚
	pdf.SetFooterFunc(func() {
		pdf.SetY(-10)
		pdf.CellFormat(
			0, 10,
			fmt.Sprintf("page %d , total page {nb}", pdf.PageNo()), //字符串中的 {nb}。大括号是可以省的，但不建议这么做
			"", 0, "C", false, 0, "",
		)
	})
	//给个空字符串就会去替换默认的 "{nb}"。
	//如果这里指定了特别的字符串，那么SetFooterFunc() 中的 "nb" 也必须换成这个特别的字符串
	pdf.AliasNbPages("")
	pdf.SetTopMargin(20)
	// Page properties
	pageWidth, pageHeight := pdf.GetPageSize()
	fmt.Println("pageWidth: ", pageWidth, "pageHeight: ", pageHeight)
	_, topMargin, _, _ := pdf.GetMargins()
	maxY := pageHeight - topMargin

	//添加多張圖片到 PDF
	imagePaths := []string{
		"/Users/chen/Downloads/personal_re/test/screenshot_files/3001c560-7949-11ee-992a-e1aa9b0ae3ae_2024-01-15_2024-01-25.png",
		"/Users/chen/Downloads/personal_re/test/screenshot_files/rocket.png",
		"/Users/chen/Downloads/personal_re/test/screenshot_files/2a97f7c0-7949-11ee-992a-e1aa9b0ae3ae_2024-01-15_2024-01-25.png",
		"/Users/chen/Downloads/personal_re/test/screenshot_files/3001c560-7949-11ee-992a-e1aa9b0ae3ae_2024-01-20_2024-01-25.png",
		"/Users/chen/Downloads/personal_re/test/screenshot_files/18c9dcc0-7949-11ee-992a-e1aa9b0ae3ae_2024-01-18_2024-01-25.png",
		// 添加更多圖片路徑...
	}

	//添加一页
	pdf.AddPage()
	//Start from top of the page
	y := topMargin
	fmt.Println("topMargin", y)

	var image1_width_cm, image1_height_cm float64
	for i := range imagePaths {
		fmt.Println(i)
		width, height := tools.GetImageHW(imagePaths[i])

		image_width_cm := (float64(width) / 40) * 3.9
		// image_width_string := fmt.Sprintf("%.1fcm", image_width_cm)

		image_height_cm := (float64(height) / 40) * 3.9
		// image_height_string := fmt.Sprintf("%.1fcm", image_height_cm)
		fmt.Println("width", image_width_cm, "height", image_height_cm)

		if i+1 < len(imagePaths) {
			width1, height1 := tools.GetImageHW(imagePaths[i+1])

			image1_width_cm = (float64(width1) / 40) * 3.9
			// image_width_string := fmt.Sprintf("%.1fcm", image_width_cm)

			image1_height_cm = (float64(height1) / 40) * 3.9
			// image_height_string := fmt.Sprintf("%.1fcm", image_height_cm)
			fmt.Println("width1", image1_width_cm, "height1", image1_height_cm)
		} else {
			image1_width_cm = 0
			image1_height_cm = 0
		}

		// Add title
		title := strings.TrimSuffix(strings.TrimPrefix(imagePaths[i], "/Users/chen/Downloads/personal_re/test/screenshot_files/"), ".png")

		fmt.Println("text Y :", y)
		pdf.Text(50, y, title)

		y = y + 5
		pdf.Image(imagePaths[i], 10, y, image_width_cm, image_height_cm, false, "", 0, "")

		fmt.Println("Image Y", y)

		// Update Y coordinate
		y += image_height_cm + 10 // Adjust spacing as needed

		fmt.Println("y += imageHeight", y)
		// Check if there is enough space for the image on the current page
		if y+image1_height_cm > maxY-20 {
			fmt.Println("y += imageHeight", y+image_height_cm, "maxY", maxY)
			// Add new page
			// pdf.AddPageFormat("P", gofpdf.SizeType{Wd: pageWidth, Ht: pageHeight})
			pdf.AddPage()
			// Reset Y coordinate
			y = topMargin
			fmt.Println("y_top", y)
		}

	}
	pdf.Text(50, y, "end")

	// Output PDF to file
	err := pdf.OutputFileAndClose("output.pdf")
	if err != nil {
		log1.Fatal(err)
	}

	fmt.Println("PDF generated successfully.")

}

// func CreatePDFbySchedule_old_v(ScheduleID int) (err error) {

// 	defer func() {
// 		if err != nil {
// 			// 进行错误处理，例如记录日志或返回错误信息给调用方
// 			log.Logrecord("ERROR", "func CreatePDFbySchedule error")

// 		}
// 	}()
// 	scheduleData, err := GetReportByScheduleID(ScheduleID)
// 	if err != nil {
// 		fmt.Println(err)
// 	}
// 	time.Sleep(10 * time.Second)

// 	// Initialize library.
// 	if err := pdf.Init(); err != nil {
// 		fmt.Println("init error")
// 	}
// 	// defer pdf.Destroy()

// 	// // Create converter.
// 	// converter, err := pdf.NewConverter()
// 	// if err != nil {
// 	// 	log1.Fatal(err)
// 	// }
// 	// defer converter.Destroy()

// 	for _, reports := range scheduleData {
// 		// defer pdf.Destroy()
// 		fmt.Println(reports.Name)
// 		timefrom := tools.Timeconverter(reports.TimeUnit, reports.TimePeriod)
// 		now := time.Now().Format("2006-01-02")
// 		outputPath := fmt.Sprintf("%s/%s_%s_%s.html", global.EnvConfig.Files.HtmlFile, reports.Name, timefrom, now)
// 		pdfname := fmt.Sprintf("%s/%s_%s_%s.pdf", global.EnvConfig.Files.ReportFile, reports.Name, timefrom, now)
// 		log.Logrecord("排程", "report name: "+reports.Name+" 開始產出")

// 		//	用 ReportID 取出 Report 的相關資料
// 		report_data, err := GetReportByReportID(reports.ID)
// 		if err != nil {
// 			fmt.Println(err)
// 		}
// 		fmt.Println(report_data.Name)
// 		element_data, err := GetElementsByReportID(reports.ID)
// 		if err != nil {
// 			fmt.Println(err)
// 		}

// 		now = time.Now().Format("2006-01-02")
// 		timefrom = tools.Timeconverter(report_data.TimeUnit, report_data.TimePeriod)
// 		var images []string
// 		for _, element := range element_data {
// 			fmt.Println(element.UID)
// 			img := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, element.UID, timefrom, now)
// 			fmt.Println(img)
// 			images = append(images, img)
// 		}
// 		total_height := 0
// 		// var width float64
// 		for _, image := range images {
// 			_, height := tools.GetImageHW(image)
// 			// width = width
// 			total_height += (height + 98)
// 		}
// 		total_height = total_height + 126
// 		fmt.Println("total_height", total_height)

// 		GeneratePDF(total_height, outputPath, pdfname)
// 		log.Logrecord("排程", "report name: "+reports.Name+" 完成產出")
// 	}
// 	pdf.Destroy()
// 	return err
// }

// func GeneratePDF(total_height int, htmlpath string, pdfname string) {

// 	// Create object from file.
// 	object, err := pdf.NewObject(htmlpath)
// 	if err != nil {
// 		log1.Fatal(err)
// 	}
// 	object.Header.ContentCenter = "[title]"
// 	// object.Header.DisplaySeparator = true
// 	object.Header.DisplaySeparator = false

// 	// Create converter.
// 	converter, err := pdf.NewConverter()
// 	if err != nil {
// 		log1.Fatal(err)
// 	}
// 	// defer converter.Destroy()

// 	// Add created objects to the converter.
// 	converter.Add(object)
// 	// converter.Add(object2)
// 	// converter.Add(object3)
// 	//1cm = 38.34px
// 	// Set converter options.
// 	// var total_height_cm float64

// 	total_height_cm := float64(total_height) / 40
// 	total_height_string := fmt.Sprintf("%.1fcm", total_height_cm)
// 	fmt.Println("total_height_string", total_height_string)
// 	converter.Title = "Sample document"
// 	converter.PaperSize = pdf.A4
// 	converter.Width = "48cm"
// 	converter.Height = total_height_string
// 	//橫向展示
// 	// converter.Orientation = pdf.Landscape
// 	converter.MarginTop = "10mm"
// 	converter.MarginBottom = "10mm"
// 	converter.MarginLeft = "10mm"
// 	converter.MarginRight = "10mm"

// 	// Convert objects and save the output PDF document.
// 	outFile, err := os.Create(pdfname)
// 	if err != nil {
// 		log1.Fatal(err)
// 	}
// 	defer outFile.Close()
// 	fmt.Println("401")

// 	if err := converter.Run(outFile); err != nil {
// 		fmt.Println("converter.Run error")
// 	}
// 	fmt.Println("407")

// 	converter.Destroy()
// }
