package services

import (
	"context"
	// "io/ioutil"
	// "log"
	"time"
	// "fmt"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/page"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
)

// elementScreenshot takes a screenshot of a specific element.
func elementScreenshot(urlstr, sel string, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.WaitVisible(`div.dashboardViewport`),
		chromedp.Sleep(20 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Use
// device.Reset to reset the emulation and viewport settings.
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}


// kibana elementScreenshot with auth takes a screenshot of a specific element.
func kibanaElementScreenshotWithAuth(loginUrl, username,password,sel string, res *[]byte) chromedp.Tasks {

	var executed *runtime.RemoteObject
	// instancedata := GetInstanceTypeAndIP(id)
	// username:="elastic"
	// password:="12345678"
	// width, height := 1920, 1080
	width, height := 1240, 1754

	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.Sleep(5 * time.Second),
		chromedp.Evaluate(`var jq = document.createElement('script'); jq.src = "https://cdn.bootcss.com/jquery/1.4.2/jquery.js"; document.getElementsByTagName('head')[0].appendChild(jq);`,&executed),
		chromedp.Sleep(5 * time.Second),
		// chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys (`input[name="username"]`, username,chromedp.NodeVisible),
		chromedp.SendKeys (`input[name="password"]`, password,chromedp.NodeVisible),
		// chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`.euiButton`),
		chromedp.Sleep(5 * time.Second),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.WaitVisible(`div.dashboardViewport`),
		chromedp.Sleep(20 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}

// grafana elementScreenshot with auth takes a screenshot of a specific element.
func grafanaElementScreenshotWithAuth(loginUrl, username,password,sel string, res *[]byte) chromedp.Tasks {

	var executed *runtime.RemoteObject
	// instancedata := GetInstanceTypeAndIP(id)
	// username:="admin"
	// password:="12345678"
	// width, height := 1920, 1080
	width, height := 1240, 1754

	return chromedp.Tasks{
		chromedp.Navigate(loginUrl),
		chromedp.Sleep(5 * time.Second),
		chromedp.Evaluate(`var jq = document.createElement('script'); jq.src = "https://cdn.bootcss.com/jquery/1.4.2/jquery.js"; document.getElementsByTagName('head')[0].appendChild(jq);`,&executed),
		chromedp.Sleep(5 * time.Second),
		// chromedp.WaitVisible(`#password`, chromedp.ByID),
		chromedp.SendKeys (`input[name="user"]`, username,chromedp.NodeVisible),
		chromedp.SendKeys (`input[name="password"]`, password,chromedp.NodeVisible),
		// chromedp.SendKeys(`#password`, password, chromedp.ByID),
		chromedp.Sleep(2 * time.Second),
		chromedp.Click(`.css-14g7ilz-button`),
		chromedp.Sleep(5 * time.Second),
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
		chromedp.WaitVisible(`div.scrollbar-view`),
		chromedp.Sleep(20 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}


// 导出指定元素为PDF
func elementPDFPrint(urlstr, sel string, res *[]byte) chromedp.Tasks {
	var err error
	return chromedp.Tasks{
	  chromedp.Navigate(urlstr),
	  chromedp.WaitVisible(`div.dashboardViewport`),
	  chromedp.Sleep(time.Duration(25) * time.Second),
	  chromedp.ActionFunc(func(ctx context.Context) error {
		// 获取pdf数据
		*res, _, err = page.PrintToPDF().Do(ctx)
		if err != nil {
		  return err
		}
		//*res = buf
		return nil
	  }),
	}
  }