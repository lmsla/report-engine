package services

import (
	"context"
	"errors"
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

func ScreenshotbySchedule(scheduleID int) (err error) {

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			fmt.Println("ScreenshotbySchedule 發生錯誤：", err)
		}
	}()

	schedule_data, err := GetReportByScheduleID(scheduleID)
	if err != nil {
		fmt.Println(err)
	}
	for _, data := range schedule_data {

		err := ScreenshotbyReport(data.ID)
		if err != nil {
			fmt.Println("ScreenshotbyReport - line 47", err.Error())
			log.Logrecord("ERROR", "ScreenshotbyReport error"+err.Error())
			return err
		}
		
	}
	return err
}

func ScreenshotbyReport(reportID int) (err error) {

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			fmt.Println("ScreenshotbyReport 發生錯誤：", err)
		}
	}()

	report_data, err := GetReportByReportID(reportID)
	if err != nil {
		fmt.Println("ScreenshotbyReport - line 36", err.Error())
		log.Logrecord("ERROR", "Get Report by Report ID error"+err.Error())
		// fmt.Println(err)
	}
	element_data, err := GetElementsByReportID(reportID)
	if err != nil {
		fmt.Println("ScreenshotbyReport - line 41", err.Error())
		log.Logrecord("ERROR", "Get Elements by Report ID error"+err.Error())
		// fmt.Println(err)
	}
	timefrom := tools.Timeconverter(report_data.TimeUnit, report_data.TimePeriod)

	for _, data := range element_data {

		log.Logrecord("截圖", "element name: "+data.Name+"開始執行截圖")

		err := Screenshot_element1(data.Type, data.Instance.URL, data.SpaceName, data.UID, timefrom, data.Instance.User, data.Instance.Password)
		if err != nil {
			fmt.Println("ScreenshotbyReport - line 88", err.Error())
			log.Logrecord("ERROR", "ScreenshotbyReport error"+err.Error())
		}

		time.Sleep(3 * time.Second)

		log.Logrecord("截圖", "element name: "+data.Name+"完成截圖")

		if err!= nil {
			return err
		}
	}
	if err != nil {
		// fmt.Println("在 ScreenshotbyReport 的錯誤")
		fmt.Println( "在 ScreenshotbyReport 的錯誤" + err.Error())
	}
	return err
}

func Screenshot_element1(element_type string, url string, space string, uid string, timefrom string, user string, password string) (err error) {

	// /// ----------------error handle----------------//////

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			fmt.Println("Screenshot_element1 發生錯誤：", err)
		}
	}()

	// capture screenshot of an element 截圖程式碼
	var buf []byte
	// 將取出來的個參數帶入網址中以便截圖
	var url1 string
	now := time.Now().Format("2006-01-02")
	switch element_type {
	case "visualiztion":
		url1 = fmt.Sprintf("%s/s/%s/app/visualize#/edit/%s?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:'%s',to:now))", url, space, uid, timefrom)
		if err := kibanaElementScreenshotWithAuth_timeout(url1, user, password, `div.css-zxsb69`, &buf); err != nil {
			fmt.Println("Screenshot_element - line 94", err.Error())
			log.Logrecord("ERROR", "Visualiztion Screenshot error"+err.Error())
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 100", err.Error())
			log.Logrecord("ERROR", "Write Visualiztion Screenshot file error"+err.Error())

		}
	case "dashboard":
		url1 = fmt.Sprintf("%s/s/%s/app/dashboards#/view/%s?_g=(time:(from:'%s',to:now))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", url, space, uid, timefrom)

		err := kibanaElementScreenshotWithAuth_timeout(url1, user, password, `div.dashboardViewport`, &buf)
		if err != nil {
			fmt.Println("Screenshot_element - line 155", err)
			log.Logrecord("ERROR", "Dashboard Screenshot error "+err.Error())

		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 161", err.Error())
			log.Logrecord("ERROR", "Write Dashboard Screenshot file error"+err.Error())

		}
		if err != nil {
			// fmt.Println("在 Screenshot_element1 ")
			fmt.Println("在 Screenshot_element1 發生錯誤 " + err.Error())
		}
		return err
	}
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	return err
}

func kibanaElementScreenshotWithAuth_timeout(loginUrl, username, password, sel string, res *[]byte) (err error) {
	var executed *runtime.RemoteObject
	// 自定義長寬
	// width, height := 1240, 1754

	defer func() {
		if err != nil {
			// 进行错误处理，例如记录日志或返回错误信息给调用方
			fmt.Println("发生错误：", err)
		}
	}()

	// 创建带超时的 context
	ctx, cancel := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancel()

	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		// chromedp.ExecPath("/usr/bin/google-chrome"),
		// chromedp.ExecPath(global.EnvConfig.Files.ChromePath),
	)

	allocCtx, cancelAlloc := chromedp.NewExecAllocator(ctx, opts...)
	defer cancelAlloc()

	ctx, cancel = chromedp.NewContext(allocCtx)
	defer cancel()

	err = chromedp.Run(ctx, chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.Sleep(3 * time.Second),
		chromedp.Evaluate(`var jq = document.createElement('script'); jq.src = "https://cdn.bootcss.com/jquery/1.4.2/jquery.js"; document.getElementsByTagName('head')[0].appendChild(jq);`, &executed),
		chromedp.Sleep(3 * time.Second),
		chromedp.SendKeys(`input[name="username"]`, username, chromedp.NodeVisible),
		chromedp.SendKeys(`input[name="password"]`, password, chromedp.NodeVisible),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`.euiButton`),
		chromedp.Sleep(3 * time.Second),
		emulation.SetDeviceMetricsOverride(0, 0, 1.0, false),
		chromedp.Sleep(15 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	})
	if err != nil {
		fmt.Println(err.Error())
	}
	return err
}

// func kibanaElementScreenshotWithAuth_old(loginUrl, username, password, sel string, res *[]byte) (err error) {
// 	fmt.Println("kibanaElementScreenshotWithAuth in")
// 	var executed *runtime.RemoteObject
// 	// 自定義長寬
// 	// width, height := 1240, 1754

// 	defer func() {
// 		if err != nil {
// 			// 进行错误处理，例如记录日志或返回错误信息给调用方
// 			fmt.Println("发生错误：", err)
// 		}
// 	}()
// 	opts := append(chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.Flag("headless", true),
// 		// chromedp.ExecPath("/usr/bin/google-chrome"),
// 		// chromedp.ExecPath(global.EnvConfig.Files.ChromePath),
// 	)
// 	fmt.Println("kibanaElementScreenshotWithAuth in 2")
// 	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
// 	ctx, cancel := chromedp.NewContext(allocCtx)
// 	defer cancel()
// 	fmt.Println("chrome error")
// 	err = chromedp.Run(ctx, chromedp.Tasks{
// 		chromedp.Navigate(loginUrl),
// 		chromedp.Sleep(3 * time.Second),
// 		chromedp.Evaluate(`var jq = document.createElement('script'); jq.src = "https://cdn.bootcss.com/jquery/1.4.2/jquery.js"; document.getElementsByTagName('head')[0].appendChild(jq);`, &executed),
// 		chromedp.Sleep(3 * time.Second),
// 		chromedp.SendKeys(`input[name="username"]`, username, chromedp.NodeVisible),
// 		chromedp.SendKeys(`input[name="password"]`, password, chromedp.NodeVisible),
// 		chromedp.Sleep(2 * time.Second),
// 		chromedp.Click(`.euiButton`),
// 		chromedp.Sleep(3 * time.Second),
// 		emulation.SetDeviceMetricsOverride(0, 0, 1.0, false),
// 		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
// 	})
// 	fmt.Println(err.Error())
// 	return err

// }

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

func Screenshot_element(element_type string, url string, space string, uid string, timefrom string, user string, password string) (err error) {

	defer func() {
		if panicError := recover(); panicError != nil {
			// 处理错误
			log.Logrecord("ERROR", "func Screenshot_element panic error")
			err = errors.New("panic error")
		} else {
			log.Logrecord("ERROR", "func Screenshot_element normal error")
			err = errors.New("normal error")
		}
	}()
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()
	opts := append(chromedp.DefaultExecAllocatorOptions[:],
		chromedp.Flag("headless", true),
		// chromedp.ExecPath("/usr/bin/google-chrome"),
		chromedp.ExecPath(global.EnvConfig.Files.ChromePath),
	)
	allocCtx, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
	ctx, cancel := chromedp.NewContext(allocCtx)
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
			fmt.Println("Screenshot_element - line 94", err.Error())
			log.Logrecord("ERROR", "Visualiztion Screenshot error"+err.Error())
		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 100", err.Error())
			log.Logrecord("ERROR", "Write Visualiztion Screenshot file error"+err.Error())

		}
	case "dashboard":
		url1 = fmt.Sprintf("%s/s/%s/app/dashboards#/view/%s?_g=(time:(from:'%s',to:now))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)", url, space, uid, timefrom)

		if err := chromedp.Run(ctx, kibanaElementScreenshotWithAuth(url1, user, password, `div.dashboardViewport`, &buf)); err != nil {
			fmt.Println("Screenshot_element - line 108", err.Error())
			log.Logrecord("ERROR", "Dashboard Screenshot error"+err.Error())

		}
		file := fmt.Sprintf("%s/%s_%s_%s.png", global.EnvConfig.Files.ScreenshotFile, uid, timefrom, now)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			fmt.Println("Screenshot_element - line 114", err.Error())
			log.Logrecord("ERROR", "Write Dashboard Screenshot file error"+err.Error())

		}

	}
	return nil
}
