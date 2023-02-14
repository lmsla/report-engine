package services

import (
	//"fmt"
	"fmt"
	"report-backend-golang/global"

	gr "github.com/mikeshimura/goreport"
	//"io/ioutil"
	// "strconv"
	//"strings"
)


func ReportCreater() {
	r := gr.CreateGoReport()
	//Page Total Function
	r.PageTotal = true
	r.SumWork["g1amtcum"] = 0.0
	r.SumWork["g2amtcum"] = 0.0
	r.SumWork["g1hrcum"] = 0.0
	r.SumWork["g2hrcum"] = 0.0
	r.SumWork["g2item"] = 0.0
	fileIpaexg := fmt.Sprintf("%s/ipaexg.ttf",global.EnvConfig.Files.FontFile)
	fileMpbold := fmt.Sprintf("%s/mplus-1p-bold.ttf",global.EnvConfig.Files.FontFile)

	font1 := gr.FontMap{
		FontName: "IPAexG",
		FileName: fileIpaexg,
	}
	font2 := gr.FontMap{
		FontName: "MPBOLD",
		FileName: fileMpbold,
	}
	fonts := []*gr.FontMap{&font1, &font2}
	r.SetFonts(fonts)
	// d := new(C1Detail)
	// r.RegisterBand(gr.Band(*d), gr.Detail)
	h := new(C1Header)
	r.RegisterBand(gr.Band(*h), gr.PageHeader)
	// f := new(C1Footer)
	// r.RegisterBand(gr.Band(*f), gr.PageFooter)
	// s1h := new(C1G1Header)
	// r.RegisterGroupBand(gr.Band(*s1h), gr.GroupHeader, 1)
	// s1 := new(C1G1Summary)
	// r.RegisterGroupBand(gr.Band(*s1), gr.GroupSummary, 1)
	// s2 := new(C1G2Summary)
	// r.RegisterGroupBand(gr.Band(*s2), gr.GroupSummary, 2)
	r.Records = gr.ReadTextFile("/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test/invoice.txt", 12)
	//fmt.Printf("Records %v \n", r.Records)
	r.SetPage("A4", "mm", "P")
	r.SetFooterY(265)
	r.Execute("complex1.pdf")
	r.SaveText("complex1.txt")
}

type C1Header struct {
}


func (h C1Header) GetHeight(report gr.GoReport) float64 {
	if report.SumWork["g2item"] == 0.0 {
		return 116
	}
	return 38
}

func (h C1Header) Execute(report gr.GoReport) {
	// cols := report.Records[report.DataPos].([]string)
	// y := 32.0
	if report.SumWork["g2item"] == 0.0 {

		rowIncreaseStep := 80.0
		rowStartPosition := 20.0
		picstart := 80.0
		picheight := 80.0
		for i := 0; i < 3; i++ {
			report.Image("/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test/testpic1.jpg", 10, rowStartPosition, 200, picstart)
			fmt.Println(picheight)
			rowStartPosition += rowIncreaseStep
			picstart += picheight
			
		}
		// report.Image("/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test/testpic1.jpg", 10, 20, 200, 80)
		// report.Image("/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test/testpic1.jpg", 10, 100, 200, 160)
		// report.Image("/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test/testpic1.jpg", 10, 180, 200, 240)

		report.Font("MPBOLD", 18, "")
		// report.LineType("straight", 1)
		// report.GrayStroke(0.9)
		// report.LineV(49, 72, 90)
		// report.LineV(150, 43, 67)
		// report.LineV(150, 71, 95)
		// report.GrayStroke(0)
		//	report.LineType("straight", 0.5)
		//	report.Rect(48, 13, 81, 21)
		report.Cell(10, 10, "報表001")
		// report.Font("MPBOLD", 9, "")
		// report.Cell(139, 45, "From")
		// x := 153.0
		// report.Cell(x, 45, "Test Consulting Corp.")
		// report.Cell(x, 51, "123 Hyde Street")
		// report.Cell(x, 57, "San Francisco, Calfornia")
		// report.Cell(x, 63, "USA")

		// report.Cell(139, 74, "To")
		// report.Cell(x, 74, cols[0])
		// report.Cell(x, 80, cols[1])
		// report.Cell(x, 86, cols[2])
		// report.Cell(x, 92, cols[3])

		// x = 14.0
		// report.Cell(x, 73, "Tax Invoice No:")
		// report.Cell(x, 79, "Tax Invoice Date:")
		// report.Cell(x, 85, "Payment Due Date:")

		// x = 52
		// report.Cell(x, 73, cols[9])
		// report.Cell(x, 79, cols[10])
		// report.Cell(x, 85, cols[11])

		// y = 110
		// y = y
	}
	// report.LineType("straight", 7)
	// report.GrayStroke(0.9)
	// report.LineH(11, y-2, 199)
	// report.GrayStroke(0)
	// report.Cell(14, y, "Type")
	// report.Cell(40, y, "Description")
	// report.Cell(161, y, "Hours")
	// report.Cell(184, y, "Amount")
	report.SumWork["g2item"] = 1.0
}
