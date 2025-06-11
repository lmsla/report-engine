package services

import (
	"bufio"
	"encoding/json"
	"report-backend-golang/entities"
	// "bytes"
	"fmt"
	// "html/template"
	// log1 "log"
	"os"
	// "report-backend-golang/entities"
	// "report-backend-golang/global"
	// "report-backend-golang/log"
	// "report-backend-golang/tools"
	"github.com/jung-kurt/gofpdf"
	"sort"
	"strconv"
	"strings"
)

// strDelimit converts 'ABCDEFG' to, for example, 'A,BCD,EFG'
// 協助函數用於格式化數字的函數
func strDelimit(str string, sepstr string, sepcount int) string {
	pos := len(str) - sepcount
	for pos > 0 {
		str = str[:pos] + sepstr + str[pos:]
		pos = pos - sepcount
	}
	return str
}

type Column struct {
	Name  string
	Alias string
	Order int
	Size  int
}

// 按照  Column.Order 從大到小排序
type ColumnSlice []Column

func (a ColumnSlice) Len() int { // 重寫 Len() method
	return len(a)
}
func (a ColumnSlice) Swap(i, j int) { // 重寫 Swap() method
	a[i], a[j] = a[j], a[i]
}
func (a ColumnSlice) Less(i, j int) bool { // 重寫 Less() method， 從大到小排序
	return a[j].Order < a[i].Order
}

func TableMeasurment(pdf *gofpdf.Fpdf, table entities.Table, table_data string) (EndY float64) {
	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 210, Ht: 10000})

	// 設置顏色、字體等屬性
	pdf.SetFillColor(126, 186, 181)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(255, 255, 255)
	pdf.SetLineWidth(.3)

	// 字體
	pdf.SetFontLocation("/Users/chen/Downloads/01BiMap/02gitlab/bimap-product/report-backend")
	pdf.AddUTF8Font("TaipeiSansTCBeta", "", "TaipeiSansTCBeta-Regular.ttf")
	pdf.SetFont("TaipeiSansTCBeta", "", 14)

	// 動態計算欄位寬度
	colWidths := make([]float64, len(table.Columns)+1)
	table.Columns = append(table.Columns, entities.Column{Name: "Count", Order: len(table.Columns) + 1})

	for i := range table.Columns {
		colWidths[i] = 200.0 / float64(len(table.Columns)) // 根據需要調整每列的初始寬度
	}

	// 計算總寬度和左邊距
	var wSum float64
	for _, w := range colWidths {
		wSum += w
	}

	var Columns []Column
	for _, data := range table.Columns {
		var column Column
		if data.Name == "Count" {
			column.Name = data.Name
			column.Order = data.Order
			Columns = append(Columns, column)
		} else {
			column.Name = data.Name[:len(data.Name)-8]
			column.Order = data.Order
			Columns = append(Columns, column)
		}

	}
	sort.Sort(sort.Reverse(ColumnSlice(Columns))) // 按照 Age 的升序排序

	left := (210 - wSum) / 2

	// 繪制表格標題
	pdf.SetX(left)
	pdf.SetFont("TaipeiSansTCBeta", "", 14)
	fmt.Println("table.Name", table.Name)
	pdf.CellFormat(0, 0, table.Name, "", 0, "C", false, 0, "")
	startX := pdf.GetX()

	fmt.Println("startX", startX)

	pdf.SetX(left)
	startX1 := pdf.GetX()
	fmt.Println("startX1", startX1)
	for i, data := range Columns {
		pdf.CellFormat(colWidths[i], 7, data.Name, "1", 0, "C", true, 0, "")
		// pdf.MultiCell(colWidths[i], 7, str, "LR", "L", true)
		// pdf.SetXY(startX+colWidths[i], startY)
		// startX = pdf.GetX()
	}
	pdf.Ln(-1)

	// 恢復顏色和字體設置
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)

	// 繪制表格數據

	// 用於存放解碼後的數據的變數
	var result []map[string]interface{}
	// 將JSON解碼為Go的map類型
	err := json.Unmarshal([]byte(table_data), &result)
	if err != nil {
		
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fill := false
	for _, c := range result {
		pdf.SetX(left)
		startX := pdf.GetX()
		startY := pdf.GetY()
		// startYY := pdf.GetY()
		// 計算需要的行高
		cellHeight := 6.0

		// 創建一個切片來保存所有的值
		// values := []string{}
		var ColumnsV []Column
		// 遍歷 map 並添加值到切片
		for k, v := range c {
			// val := fmt.Sprintf("%v", v)
			// values = append(values, val)
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

		// 繪製每個動態欄位
		// values := []string{c.nameStr, c.capitalStr, strDelimit(c.areaStr, ",", 3), strDelimit(c.popStr, ",", 3), "集先鋒", "This is a test message for report testing"}

		for i, val := range ColumnsV {
			if i == len(ColumnsV)-1 {
				// 動態處理最後一個欄位的多行文本
				pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "L", fill)
			} else {
				pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "L", fill)
				// pdf.CellFormat(colWidths[i], cellHeight, val, "LR", 0, "", fill, 0, "")
				// 更新 X 軸位置
				pdf.SetXY(startX+colWidths[i], startY)
				startX = pdf.GetX()
				// startYY = pdf.GetY()

			}
			// fmt.Printf("startYY: %v\n", startYY)
		}
		// 移動到下一行
		pdf.Ln(-1)
		fill = !fill
		EndY = pdf.GetY()

	}
	EndY = EndY + 20
	// pdf.SetY(EndY)
	fmt.Printf("Measurement EndY: %v\n", EndY)
	pdf.SetX(left)
	pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")

	return EndY
}

func TableCreater(table entities.Table, table_data string, EndY float64) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 210, Ht: EndY + 25})

	// 設置顏色、字體等屬性
	pdf.SetFillColor(126, 186, 181)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(255, 255, 255)
	pdf.SetLineWidth(.3)

	// 字體
	pdf.SetFontLocation("/Users/chen/Downloads/01BiMap/02gitlab/bimap-product/report-backend")
	pdf.AddUTF8Font("TaipeiSansTCBeta", "", "TaipeiSansTCBeta-Regular.ttf")
	pdf.SetFont("TaipeiSansTCBeta", "", 14)

	// 動態計算欄位寬度
	colWidths := make([]float64, len(table.Columns)+1)
	// 加入 count 欄位
	table.Columns = append(table.Columns, entities.Column{Name: "Count", Order: len(table.Columns) + 1})
	for i := range table.Columns {
		if i == len(table.Columns)-1 {
			colWidths[i] = 25
		} else {
			colWidths[i] = 175 / float64(len(table.Columns)-1)
		}
		// colWidths[i] = 200.0 / float64(len(table.Columns)) // 您可以根據需要調整每列的初始寬度
	}

	// 計算總寬度和左邊距
	var wSum float64
	for _, w := range colWidths {
		wSum += w

	}

	var Columns []Column
	for _, data := range table.Columns {
		var column Column
		if data.Name == "Count" {
			column.Name = data.Name
			column.Alias = "資料筆數"
			column.Order = data.Order
			Columns = append(Columns, column)
		} else {
			column.Name = data.Name[:len(data.Name)-8]
			column.Alias = data.Alias
			column.Order = data.Order
			Columns = append(Columns, column)
		}

	}
	sort.Sort(sort.Reverse(ColumnSlice(Columns))) // 按照 Age 的升序排序

	left := (210 - wSum) / 2
	pdf.SetX(left)

	pdf.SetFont("TaipeiSansTCBeta", "", 14)
	fmt.Println("table.Name", table.Name)
	pdf.CellFormat(0, 0, table.Name, "", 0, "C", false, 0, "")

	// 繪制表格標題
	pdf.SetX(left)

	for i, data := range Columns {
		pdf.CellFormat(colWidths[i], 7, data.Name, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// 恢復顏色和字體設置
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)

	// 繪制表格數據
	// 用於存放解碼後的數據的變數
	var result []map[string]interface{}
	// 將JSON解碼為Go的map類型
	err := json.Unmarshal([]byte(table_data), &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	fill := false
	for _, c := range result {
		pdf.SetX(left)
		startX := pdf.GetX()
		startY := pdf.GetY()
		// startYY := pdf.GetY()
		// 計算需要的行高
		cellHeight := 6.0

		// 創建一個切片來保存所有的值
		// values := []string{}
		var ColumnsV []Column
		// 遍歷 map 並添加值到切片
		for k, v := range c {
			// val := fmt.Sprintf("%v", v)
			// values = append(values, val)
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

		// 繪製每個動態欄位
		// values := []string{c.nameStr, c.capitalStr, strDelimit(c.areaStr, ",", 3), strDelimit(c.popStr, ",", 3), "集先鋒", "This is a test message for report testing"}

		for i, val := range ColumnsV {
			if i == len(ColumnsV)-1 {
				// 動態處理最後一個欄位的多行文本
				pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "R", fill)
			} else {
				pdf.MultiCell(colWidths[i], cellHeight, val.Name, "LR", "R", fill)
				// pdf.CellFormat(colWidths[i], cellHeight, val, "LR", 0, "", fill, 0, "")
				// 更新 X 軸位置
				pdf.SetXY(startX+colWidths[i], startY)
				startX = pdf.GetX()
				// startYY = pdf.GetY()

			}
			// fmt.Printf("startYY: %v\n", startYY)
		}
		// 移動到下一行
		pdf.Ln(-1)
		fill = !fill
		EndY = pdf.GetY()

	}
	fmt.Printf("EndY create: %v\n", EndY)
	pdf.SetX(left)
	pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")

	filename := fmt.Sprintf("%s.pdf", table.Name)
	if err := pdf.OutputFileAndClose(filename); err != nil {
		panic(err.Error())
	}

}

func TableCreaterAd(pdf *gofpdf.Fpdf, table entities.Table, table_data string, EndY float64) {
	// pdf := gofpdf.New("P", "mm", "A4", "")
	// pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 210, Ht: EndY + 25})

	// 設置顏色、字體等屬性
	pdf.SetFillColor(126, 186, 181)
	pdf.SetTextColor(255, 255, 255)
	pdf.SetDrawColor(255, 255, 255)
	pdf.SetLineWidth(.3)

	// 字體
	pdf.SetFontLocation("/Users/chen/Downloads/01BiMap/02gitlab/bimap-product/report-backend")
	pdf.AddUTF8Font("TaipeiSansTCBeta", "", "TaipeiSansTCBeta-Regular.ttf")
	pdf.SetFont("TaipeiSansTCBeta", "", 14)

	// 動態計算欄位寬度
	colWidths := make([]float64, len(table.Columns)+1)
	// 加入 count 欄位
	table.Columns = append(table.Columns, entities.Column{Name: "Count", Order: len(table.Columns) + 1})
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
		if data.Name == "Count" {
			column.Name = data.Name
			column.Order = data.Order
			Columns = append(Columns, column)
		} else {
			column.Name = data.Name[:len(data.Name)-8]
			column.Order = data.Order
			Columns = append(Columns, column)
		}

	}
	sort.Sort(sort.Reverse(ColumnSlice(Columns))) // 按照 Age 的升序排序

	left := (210 - wSum) / 2
	// 繪制表格標題
	pdf.SetX(left)

	pdf.SetFont("TaipeiSansTCBeta", "", 14)
	fmt.Println("table.Name", table.Name)
	pdf.CellFormat(0, 0, table.Name, "", 0, "C", false, 0, "")

	pdf.SetX(left)
	for i, data := range Columns {
		pdf.CellFormat(colWidths[i], 7, data.Name, "1", 0, "C", true, 0, "")
	}
	pdf.Ln(-1)

	// 恢復顏色和字體設置
	pdf.SetFillColor(224, 235, 255)
	pdf.SetTextColor(0, 0, 0)
	// 繪制表格數據
	// 用於存放解碼後的數據的變數
	var result []map[string]interface{}
	// 將JSON解碼為Go的map類型
	err := json.Unmarshal([]byte(table_data), &result)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}
	fill := false

	for _, c := range result {

		pdf.SetX(left)
		startX := pdf.GetX()
		startY := pdf.GetY()
		// 計算需要的行高
		cellHeight := 6.0
		// 創建一個切片來保存所有的值
		// values := []string{}
		var ColumnsV []Column
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

		// 繪製每個動態欄位
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
		// 移動到下一行
		pdf.Ln(-1)
		fill = !fill
		// EndY = pdf.GetY()

		// EndY = EndY + 20

	}
	fmt.Printf("EndY create: %v\n", EndY)
	pdf.SetX(left)
	pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")
	// filename := fmt.Sprintf("%s.pdf", table.Name)
	// if err := pdf.OutputFileAndClose(filename); err != nil {
	// 	panic(err.Error())
	// }
}

func ExampleFpdf_CellFormat_tables1() (EndY float64) {
	pdf := gofpdf.New("P", "mm", "A4", "")
	type countryType struct {
		nameStr, capitalStr, areaStr, popStr string
	}
	countryList := make([]countryType, 0, 8)

	// 動態 header，可以根據需要更改
	header := []string{"Country", "Capital", "Area(sq km)", "Pop.(thousands)", "ok", "yaya"}

	// 加載數據的函數
	loadData := func(fileStr string) {
		fl, err := os.Open(fileStr)
		if err == nil {
			scanner := bufio.NewScanner(fl)
			var c countryType
			for scanner.Scan() {
				lineStr := scanner.Text()
				list := strings.Split(lineStr, ";")
				if len(list) == 4 {
					c.nameStr = list[0]
					c.capitalStr = list[1]
					c.areaStr = list[2]
					c.popStr = list[3]
					countryList = append(countryList, c)
				} else {
					err = fmt.Errorf("error tokenizing %s", lineStr)
				}
			}
			fl.Close()
			if len(countryList) == 0 {
				err = fmt.Errorf("error loading data from %s", fileStr)
			}
		}
		if err != nil {
			pdf.SetError(err)
		}
	}
	// var EndY float64
	// 動態生成表格的函數
	fancyTable := func() {

		// 設置顏色、字體等屬性
		pdf.SetFillColor(126, 186, 181)
		pdf.SetTextColor(255, 255, 255)
		pdf.SetDrawColor(255, 255, 255)
		pdf.SetLineWidth(.3)
		// 字體
		pdf.SetFontLocation("/Users/chen/Downloads/01BiMap/02gitlab/bimap-product/report-backend")
		pdf.AddUTF8Font("TaipeiSansTCBeta", "", "TaipeiSansTCBeta-Regular.ttf")
		pdf.SetFont("TaipeiSansTCBeta", "", 14)
		// pdf.SetFont("", "B", 0)

		// 動態計算欄位寬度
		colWidths := make([]float64, len(header))
		for i := range header {
			colWidths[i] = 200.0 / float64(len(header)) // 您可以根據需要調整每列的初始寬度
		}

		// 計算總寬度和左邊距
		var wSum float64
		for _, w := range colWidths {
			wSum += w
		}
		left := (210 - wSum) / 2

		// 繪制表格標題
		pdf.SetX(left)
		// startX := pdf.GetX()
		// startY := pdf.GetY()
		for i, str := range header {
			pdf.CellFormat(colWidths[i], 7, str, "1", 0, "C", true, 0, "")
			// pdf.MultiCell(colWidths[i], 7, str, "LR", "L", true)
			// pdf.SetXY(startX+colWidths[i], startY)
			// startX = pdf.GetX()
		}
		pdf.Ln(-1)

		// 恢復顏色和字體設置
		pdf.SetFillColor(224, 235, 255)
		pdf.SetTextColor(0, 0, 0)
		// pdf.SetFont("", "", 0)

		// 繪制表格數據
		fill := false
		for _, c := range countryList {
			pdf.SetX(left)
			startX := pdf.GetX()
			startY := pdf.GetY()
			startYY := pdf.GetY()
			// 計算需要的行高
			cellHeight := 6.0

			// 繪製每個動態欄位
			values := []string{c.nameStr, c.capitalStr, strDelimit(c.areaStr, ",", 3), strDelimit(c.popStr, ",", 3), "集先鋒", "This is a test message for report testing"}
			for i, val := range values {
				if i == len(values)-1 {
					// 動態處理最後一個欄位的多行文本
					pdf.MultiCell(colWidths[i], cellHeight, val, "LR", "L", fill)
				} else {
					pdf.MultiCell(colWidths[i], cellHeight, val, "LR", "L", fill)
					// pdf.CellFormat(colWidths[i], cellHeight, val, "LR", 0, "", fill, 0, "")
					// 更新 X 軸位置
					pdf.SetXY(startX+colWidths[i], startY)
					startX = pdf.GetX()
					startYY = pdf.GetY()

				}
				fmt.Printf("startYY: %v\n", startYY)
			}
			// 移動到下一行
			pdf.Ln(3.0)
			fill = !fill
			EndY = pdf.GetY()

		}

		fmt.Printf("EndY: %v\n", EndY)
		pdf.SetX(left)
		pdf.CellFormat(wSum, 0, "", "T", 0, "", false, 0, "")

	}

	loadData("/Users/chen/Downloads/01BiMap/02gitlab/bimap-product/report-backend/services/countries.txt")
	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 210, Ht: 10000})

	fancyTable()

	fmt.Printf("EndY1: %v\n", EndY)

	pdf.AddPageFormat("P", gofpdf.SizeType{Wd: 210, Ht: EndY + 30})
	fancyTable()

	if err := pdf.OutputFileAndClose("test1.pdf"); err != nil {
		panic(err.Error())
	}
	return EndY

}
