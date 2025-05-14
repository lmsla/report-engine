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
	"encoding/json"
	"github.com/jung-kurt/gofpdf"
	"sort"
	"strconv"
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


// 不分頁
func GeneratePDF_by_gofpdf_No_seprate(elementData []entities.Element, tableData []entities.Table, report_name string, timefrom string, now string) {
	pdfname := fmt.Sprintf("%s/%s_%s_%s.pdf", global.EnvConfig.Files.ReportFile, report_name, timefrom, now)

	fontSize := float64(14)
	//// time 格式為 2024-10-09

	//// 創建 PDF 文件
	pdf := gofpdf.New("L", "mm", "A4", "")
	// add page
	var width1, total_height float64
	fmt.Println("len", len(elementData))
	if len(elementData) != 0 {
		for _, element := range elementData {
			image_name := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, element.UID, timefrom, now)
			width, height := tools.GetImageHW(image_name)
			width1 = (float64(width) / 40 * 4) + 20
			total_height += (float64(height) / 40 * 4) + 20
		}
		total_height = total_height + 27
	} else {
		total_height = 27
		width1 = 210
	}
	fmt.Println("total_height 219", total_height)
	var total_table_height float64

	for _, table := range tableData {

		_, data_len := DataDealing(table, timefrom, now)
		table_height := float64(data_len)*6 + 7
		total_table_height += table_height + 70

	}
	fmt.Println("total_table_height", total_table_height)
	total_height = total_height + total_table_height
	fmt.Println("pdf_height", total_height)

	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: width1, Ht: total_height})

	// pdf.SetFontLocation("/src")
	pdf.SetFontLocation(global.EnvConfig.Files.FontFile)
	pdf.AddUTF8Font("Taipei Sans TC Beta", "", "TaipeiSansTCBeta-Regular.ttf")

	// 報表名稱
	pdf.SetFont("Taipei Sans TC Beta", "", 20)
	pdf.CellFormat(0, 0, report_name, "", 0, "C", false, 0, "")

	// set logo
	pdf.Image(global.EnvConfig.Files.LogoFile, 5, 5, 20, 20, false, "", 0, "")

	//// 给個空字符串就会去替换默认的 "{nb}"。
	//// 如果这里指定了特别的字符串，那么SetFooterFunc() 中的 "nb" 也必须换成这个特别的字符串
	// pdf.AliasNbPages("")
	pdf.SetTopMargin(25)
	// Page properties
	pageWidth, pageHeight := pdf.GetPageSize()
	fmt.Println("pageWidth: ", pageWidth, "pageHeight: ", pageHeight)

	_, topMargin, _, _ := pdf.GetMargins()

	//Start from top of the page
	var y float64

	// fmt.Println("topMargin", topMargin)
	y = topMargin
	if len(elementData) != 0 {

		for _, element := range elementData {

			image_name := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, element.UID, timefrom, now)
			width, height := tools.GetImageHW(image_name)

			image_width_cm := (float64(width) / 40) * 4

			image_height_cm := (float64(height) / 40) * 4

			// fmt.Println("width", image_width_cm, "height", image_height_cm)

			// Add 圖表標題
			// fmt.Println("text Y :", y)
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

			// fmt.Println("Image Y", y)

			// Update Y coordinate
			y = y + image_height_cm + 15 // Adjust spacing as needed
			// fmt.Println("y += imageHeight", y)

		}
	}
	//// 開始產出 table
	pdf.SetY(y)
	for _, table := range tableData {
		y = pdf.GetY()
		pdf.SetY(y)
		pdf.SetFont("Taipei Sans TC Beta", "", 14)
		pdf.CellFormat(0, 0, table.Name, "", 0, "C", false, 0, "")

		period := fmt.Sprintf("報表期間:%s~%s", timefrom, now)
		pdf.SetFont("Taipei Sans TC Beta", "", 10)
		pdf.CellFormat(0, 0, period, "", 0, "R", false, 0, "")

		table_data, _ := DataDealing(table, timefrom, now)
		// 設置顏色、字體等屬性
		pdf.SetFillColor(126, 186, 181)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(0.3)

		// 字體

		pdf.SetFontLocation(global.EnvConfig.Files.FontFile)
		pdf.AddUTF8Font("Taipei Sans TC Beta", "", "TaipeiSansTCBeta-Regular.ttf")
		pdf.SetFont("Taipei Sans TC Beta", "", 14)

		// 動態計算欄位寬度
		colWidths := make([]float64, len(table.Columns)+1)
		// 加入 count 欄位 (可調整欄位顯示)
		table.Columns = append(table.Columns, entities.Column{Name: "資料筆數", Order: len(table.Columns) + 1})
		for i := range table.Columns {
			if i == len(table.Columns)-1 {
				colWidths[i] = 25
			} else {
				colWidths[i] = 175 / float64(len(table.Columns)-1)
			}
		}

		// 計算總寬度和左邊距
		var wSum float64
		for _, w := range colWidths {
			wSum += w
		}

		var Columns []Column
		for _, data := range table.Columns {
			var column Column
			//// 如果選擇 .keyword 欄位，移除掉 .keyword 字樣(現階段也只提供 .keyword 欄位選取)
			if strings.Contains(data.Name, ".keyword") {
				column.Name = data.Name[:len(data.Name)-8]
				column.Order = data.Order
				column.Alias = data.Alias
				Columns = append(Columns, column)
			} else {
				column.Name = data.Name
				column.Order = data.Order
				column.Alias = data.Alias
				Columns = append(Columns, column)
			}

		}
		sort.Sort(sort.Reverse(ColumnSlice(Columns))) // 按照 Age 的升序排序

		left := (210 - wSum) / 2
		// 繪制表格標題
		y = y + 5
		pdf.SetY(y)

		pdf.SetX(left)
		// pdf.SetDrawColor(0, 0, 0)
		// pdf.SetLineWidth(0.1)

		for i, data := range Columns {
			if data.Alias != "" {
				pdf.CellFormat(colWidths[i], 7, data.Alias, "1", 0, "C", true, 0, "")
			} else {
				pdf.CellFormat(colWidths[i], 7, data.Name, "1", 0, "C", true, 0, "")
			}

		}
		pdf.Ln(-1)

		// 恢復顏色和字體設置
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		pdf.SetDrawColor(0, 0, 0)
		pdf.SetLineWidth(0.3)
		pdf.SetFont("Taipei Sans TC Beta", "", 12)
		// 繪制表格數據
		// 用於存放解碼後的數據的變數
		var result []map[string]interface{}
		// 將JSON解碼為Go的map類型
		err := json.Unmarshal([]byte(table_data), &result)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}
		// // var prevRow map[string]interface{} // 用於存儲上一行的數據
		// var prevRow []Column
		fill := false
		for _, c := range result {

			pdf.SetX(left)
			startX := pdf.GetX()
			startY := pdf.GetY()
			// 計算需要的行高
			cellHeight := 6.0
			// 創建一個切片來保存所有的值

			var ColumnsV []Column
			// c = table 的欄位值
			// 遍歷 map 並添加值到切片
			for k, v := range c {
				var columnV Column
				if k == "Count" {
					columnV.Name = fmt.Sprintf("%v", v)
					columnV.Order = len(c) + 1
				} else {
					columnV.Name = v.(string)
					i, _ := strconv.Atoi(k)
					columnV.Order = i
				}
				ColumnsV = append(ColumnsV, columnV)
			}
			sort.Sort(sort.Reverse(ColumnSlice(ColumnsV))) // 按照 Age 的升序排序

			// 繪製每個動態欄位(儲存格不合併)-------------
			for i, val := range ColumnsV {
				if i == len(ColumnsV)-1 {
					// 動態處理最後一個欄位的多行文本
					pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "R", fill)
				} else {
					pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "R", fill)
					// 更新 X 軸位置
					pdf.SetXY(startX+colWidths[i], startY)
					startX = pdf.GetX()
				}
			}
			// -------------------------------------------

			// pdf.SetFont("Taipei Sans TC Beta", "", 12)

			// // 繪製每個動態欄位(儲存格合併)
			// for i, val := range ColumnsV {

			// 	// 檢查該欄位值是否與上一行相同，若相同則合併儲存格
			// 	if i == len(ColumnsV)-1 {
			// 		// 動態處理最後一個欄位的多行文本
			// 		pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LRB", "C", fill)
			// 	} else {
			// 		if prevRow != nil && prevRow[i].Name == val.Name {

			// 			pdf.MultiCell(colWidths[i], cellHeight, "", "LR", "C", fill)
			// 			pdf.SetXY(startX+colWidths[i], startY)
			// 			startX = pdf.GetX()

			// 		} else {

			// 			pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LRT", "C", fill)
			// 			// 更新 X 軸位置
			// 			pdf.SetXY(startX+colWidths[i], startY)
			// 			startX = pdf.GetX()
			// 		}
			// 	}
			// }
			// // 將當前行設置為上一行，供下一行比較使用
			// prevRow = ColumnsV
			// // -------------------------------------------

			// pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
			// 移動到下一行
			// pdf.Ln(-1)
			fill = !fill

		}

		//// 繪製表格底邊
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "B", 0, "", false, 0, "")
		y = pdf.GetY() + 5
		pdf.SetY(y)
	}

	err := pdf.OutputFileAndClose(pdfname)
	if err != nil {
		log1.Fatal(err)
	}

	fmt.Println("PDF generated successfully.")

}

///////////////-------------棄用--------------///////////////



func CreateHtmlbySchedule(nowtime int64, ScheduleID int) (err error) {

	defer func() {
		if err != nil {
			// Error handling
			log.Logrecord("ERROR", "func CreateHtmlbySchedule error")
			// fmt.Println("func CreateHtmlbySchedule  發生錯誤：", err)
		}
	}()

	scheduleData, err := GetReportByScheduleID(ScheduleID)
	if err != nil {
		fmt.Println(err)
		log.Logrecord("ERROR", "Get Report by Schedule ID error"+err.Error())
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
			// Error handling
			log.Logrecord("ERROR", "func CreatePDFbySchedule error")
			// fmt.Println("func CreateHtmlbySchedule  發生錯誤：", err)
		}
	}()

	//	用 ReportID 取出 Report 的相關資料
	report_data, err := GetReportByReportID(ReportId)
	if err != nil {
		log.Logrecord("ERROR", "Get Report By ReportID error at CreateHtml stage: "+err.Error())
		fmt.Println(err)
	}
	fmt.Println(report_data.Name)

	element_data, err := GetElementsByReportID(ReportId)
	if err != nil {
		log.Logrecord("ERROR", "Get Elements By ReportID error at CreateHtml stage: "+err.Error())
		fmt.Println(err)
	}
	t := time.Unix(nowtime, 0)
	now := t.Format("2006-01-02")
	timefrom := tools.Timeconverter(nowtime, report_data.TimeUnit, report_data.TimePeriod, report_data.Alias)

	data1 := Report{}
	uu := new(Report)
	gg := new(Element)
	uu.Name = report_data.Name
	for _, elements := range element_data {
		// uu := new(Report)
		gg.Img = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID, timefrom, now)
		gg.Name = elements.Name
		gg.Period = timefrom+"~"+now
		// uu.Elements = fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, elements.UID,timefrom,now)
		data1.Elements = append(data1.Elements, *gg)
	}

	data1 = Report{
		Name:     uu.Name,
		Elements: data1.Elements,
	}

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
	// 1.18 時的舊寫法
	// w.WriteString(string(processed.Bytes()))
	w.WriteString(processed.String())
	w.Flush()

	return err
}

func CreatePDFbySchedule(nowtime int64, ScheduleID int) (err error) {
	// nowtime = time_execute
	defer func() {
		if err != nil {
			// Error handling
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
		// timefrom := tools.Timeconverter(nowtime, reports.TimeUnit, reports.TimePeriod, reports.Alias)
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

		table_data, err := GetTableByReportID(reports.ID)
		if err != nil {
			fmt.Println(err)
		}

		timefrom := tools.Timeconverter(nowtime, report_data.TimeUnit, report_data.TimePeriod, report_data.Alias)

		GeneratePDF_by_gofpdf_No_seprate(element_data, table_data, reports.Name, timefrom, now)

		log.Logrecord("排程", "report name: "+reports.Name+" 完成產出")
	}

	return err
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


