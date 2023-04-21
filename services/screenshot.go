package services

import (
	"context"
	"fmt"
	"github.com/chromedp/cdproto/emulation"
	"io/ioutil"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"report-backend-golang/tools"
	"time"
	// "github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/chromedp/device"
)

func ScreenshotbySchedule(scheduleID int) {
	schedule_data, err := GetReportByScheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range schedule_data {

		ScreenshotbyReport(data.ID)
	}

}

func ScreenshotbyReport(reportID int) {
	log.Logrecord("debug", "screenshot by report in:")
	fmt.Println("1111")
	report_data, err := GetReportByReportID(reportID)
	if err != nil {
		fmt.Println("ScreenshotbyReport - line 33")
		fmt.Println(err)
	}
	element_data, err := GetElementsByReportID(reportID)
	if err != nil {
		fmt.Println("ScreenshotbyReport - line 38")
		fmt.Println(err)
	}
	timefrom := tools.Timeconverter(report_data.TimeUnit, report_data.TimePeriod)
	for _, data := range element_data {
		fmt.Println("2222")
		log.Logrecord("截圖", "data.Type: "+data.Name+", data.Instance.URL: "+data.Instance.URL+", SpaceName: "+data.SpaceName+", UID: "+data.UID+" ,timefrom: "+timefrom+" , Instance.User: "+ data.Instance.User+" , data.Instance.Password: "+data.Instance.Password)
		log.Logrecord("截圖", "element name: "+data.Name+"開始執行截圖1")
		Screenshot_element(data.Type, data.Instance.URL, data.SpaceName, data.UID, timefrom, data.Instance.User, data.Instance.Password)
		time.Sleep(3 * time.Second)
		log.Logrecord("截圖", "element name: "+data.Name+"完成截圖1")
	}

}

func Screenshot_element(element_type string, url string, space string, uid string, timefrom string, user string, password string) {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// // 創建超時上下文
	// ctx, cancel = context.WithTimeout(ctx, 20*time.Second)
	// defer cancel()

	// capture screenshot of an element 截圖程式碼
	var buf []byte
	// 將取出來的個參數帶入網址中以便截圖
	var url1 string
	now := time.Now().Format("2006-01-02")
	switch element_type {
	case "visualiztion":
		url1 = fmt.Sprintf("%s/s/%s/app/visualize#/edit/%s?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:'%s',to:now))", url, space, uid, timefrom)
		if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, user, password, `div.css-zxsb69`, &buf)); err != nil {
			// log.Fatal(err)
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 76")
			// log.Fatal(err)
		}
	case "dashboard":
		url1 = fmt.Sprintf("%s/s/%s/app/dashboards#/view/%s?_g=(time:(from:'%s',to:now))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", url, space, uid, timefrom)
		log.Logrecord("截圖url", url1)
		if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, user, password, `div.dashboardViewport`, &buf)); err != nil {
			fmt.Println("Screenshot_element - line 82")
			// log.Fatal(err)
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 87")
			// log.Fatal(err)
		}

	}

}

// kibana elementScreenshot with auth takes a screenshot of a specific element.
func kibanaElementScreenshotWithAuth(loginUrl, username, password, sel string, res *[]byte) chromedp.Tasks {

	var executed *runtime.RemoteObject
	// 自定義長寬
	// width, height := 1240, 1754

	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.Sleep(3 * time.Second),
		chromedp.Evaluate(`var jq = document.createElement('script'); jq.src = "https://cdn.bootcss.com/jquery/1.4.2/jquery.js"; document.getElementsByTagName('head')[0].appendChild(jq);`, &executed),
		chromedp.Sleep(3 * time.Second),
		// chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys(`input[name="username"]`, username, chromedp.NodeVisible),
		chromedp.SendKeys(`input[name="password"]`, password, chromedp.NodeVisible),
		// chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`.euiButton`),
		chromedp.Sleep(3 * time.Second),
		//如果要自訂長寬
		// emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		// 使用原圖的長寬比
		emulation.SetDeviceMetricsOverride(0, 0, 1.0, false),
		// chromedp.WaitVisible(`div.dashboardViewpxort`),
		// chromedp.WaitVisible(`div.css-zxsb69`),
		chromedp.Sleep(15 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
		// chromedp.Emulate(device.Reset),
	}

}
