package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"report-backend-golang/global"
	"report-backend-golang/log"
	"report-backend-golang/tools"
	"time"

	"github.com/chromedp/cdproto/emulation"

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
	report_data,err := GetReportByReportID(reportID)
	if err != nil {
		fmt.Println(err)
	}
	element_data, err := GetElementsByReportID(reportID)
	if err != nil {
		fmt.Println(err)
	}
	timefrom := tools.Timeconverter(report_data.TimeUnit,report_data.TimePeriod)
	for _, data := range element_data {
		log.Logrecord("截圖", "element name: "+data.Name+"開始執行截圖")
		Screenshot_element(data.Type, data.Instance.URL, data.SpaceName, data.UID, timefrom,data.Instance.User, data.Instance.Password)
		time.Sleep(3 * time.Second)
		log.Logrecord("截圖", "element name: "+data.Name+"完成截圖")
	}

}

func Screenshot_element(element_type string, url string, space string, uid string, timefrom string, user string, password string) {

	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()
	// 用 report 中 instance 的 instance ID 去 instance 資料庫中撈出 instance 的 IP,type,user,password
	// auth, err := GetInstanceTypeAndIP(data1.InstanceID)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// capture screenshot of an element 截圖程式碼，擺著就好勿動
	var buf []byte
	// 將取出來的個參數帶入網址中以便截圖
	var url1 string
	now := time.Now().Format("2006-01-02")
	switch element_type {
	case "visualiztion":
		url1 = fmt.Sprintf("%s/s/%s/app/visualize#/edit/%s?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:%s,to:now))", url, space, uid, timefrom)
		if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, user, password, `div.css-zxsb69`, &buf)); err != nil {
			// log.Fatal(err)
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid,timefrom,now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			// log.Fatal(err)
		}
	case "dashboard":
		url1 = fmt.Sprintf("%s/s/%s/app/dashboards#/view/%s?_g=(time:(from:%s,to:now))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", url, space, uid,timefrom)
		if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, user, password, `div.dashboardViewport`, &buf)); err != nil {
			// log.Fatal(err)
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid,timefrom,now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			// log.Fatal(err)
		}

	}

	// if element_type == "visualiztion" {
	// 	url1 = fmt.Sprintf("%s/app/visualize#/edit/%s?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-1y/d,to:now))", url, uid)
	// } else if element_type == "dashboard" {
	// 	url1 = fmt.Sprintf("%s/app/dashboards#/view/%s?_g=(time:(from:now-1y/d,to:now))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", url, uid)
	// }

	// url := "http://10.99.1.110:5601/kibana_iframe/app/visualize#/edit/41945620-ef08-11ec-b4a6-b712e673cd3e?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-1y/d,to:now))"
	// url := "http://10.99.1.110:5601/kibana_iframe/s/dashboard/app/dashboards#/view/3e4f8080-8234-11eb-b8a7-6d0097974bf2?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))"
	// url := fmt.Sprintf("%s/app/dashboards#/view/%s?_g=(time:(from:%s,to:%s))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", auth["IP"], data2.UID, reportinfo.From, reportinfo.To)
	// 帶入 url,instance 的 username 、 password 帶入截圖的 function
	// if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, "elastic", "RnIv7YhigaVKS=l-*yz9", `div.css-zxsb69`, &buf)); err != nil {
	// 	// log.Fatal(err)
	// }
	// if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url, "elastic", "RnIv7YhigaVKS=l-*yz9", `div.dashboardViewport`, &buf)); err != nil {
	// 	// log.Fatal(err)
	// }
	// // 寫到指定路徑，暫定以 dashboard UID 命名
	// file := fmt.Sprintf("%s/%s.png", "/Users/chen/Downloads/BiMap/CodeBackup/screenshot_test", uid)
	// if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
	// 	// log.Fatal(err)
	// }

}

// kibana elementScreenshot with auth takes a screenshot of a specific element.
func kibanaElementScreenshotWithAuth(loginUrl, username, password, sel string, res *[]byte) chromedp.Tasks {

	var executed *runtime.RemoteObject
	// instancedata := GetInstanceTypeAndIP(id)
	// username:="elastic"
	// password:="12345678"
	// width, height := 1920, 1080
	width, height := 1240, 1754

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
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		// chromedp.WaitVisible(`div.dashboardViewport`),
		// chromedp.WaitVisible(`div.css-zxsb69`),
		chromedp.Sleep(10 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
		// chromedp.Emulate(device.Reset),
	}
}
