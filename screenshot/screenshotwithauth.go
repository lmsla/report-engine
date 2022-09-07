package screenshot

import (
	"context"
	"github.com/chromedp/cdproto/emulation"
	"github.com/chromedp/cdproto/runtime"
	"github.com/chromedp/chromedp"
	"io/ioutil"
	"log"
	"strings"
	"time"
	"fmt"
)

func Screenshot1()  {
	var buf []byte

	// create chrome instance
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// create a timeout
	ctx, cancel = context.WithTimeout(ctx, 50*time.Second)
	defer cancel()

	// run task list
	var res string
	var err error

	width, height := 1920, 1080
	err=chromedp.Run(ctx, chromedp.Tasks{
		emulation.SetDeviceMetricsOverride(int64(width), int64(height), 1.0, false),
	})

	loginUrl:=`http://10.99.1.117:5601/login?next=%2Fapp%2Fdashboards#/view/cee6d680-16c0-11ed-8f9a-d7ff1d0d43e3?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))`

	var executed *runtime.RemoteObject
	username:="elastic"
	password:="12345678"


	err = chromedp.Run(ctx, chromedp.Tasks{
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
		chromedp.CaptureScreenshot(&buf),
	})

	if err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("1.png", buf, 0644); err != nil {
		log.Fatal(err)
	}

	indexPageUrl:=`http://10.99.1.117:5601/app/dashboards#/view/cee6d680-16c0-11ed-8f9a-d7ff1d0d43e3?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))`

	
	// url := fmt.Sprintf("http://10.99.1.138:5601/app/dashboards#/view/%s?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)",uid)
	// fmt.Printf("uid is %s\n",uid)
	if err := chromedp.Run(ctx,elementScreenshotwithauth1(indexPageUrl, `div.dashboardViewport`,&buf)); err != nil {
		log.Fatal(err)
	}
	file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/test1.png")
	if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
		log.Fatal(err)
	}
	log.Printf("got: `%s`", strings.TrimSpace(res))

}

// func ScreenshotWithAuth() {
// 	// create context
// 	ctx, cancel := chromedp.NewContext(
// 		context.Background(),
// 		// chromedp.WithDebugf(log.Printf),
// 	)
// 	defer cancel()

// 	// capture screenshot of an element
// 	var buf []byte

// 	// url := fmt.Sprintf("http://10.99.1.138:5601/app/dashboards#/view/%s?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)",uid)
// 	url := `http://10.99.1.117:5601/login?next=%2Fapp%2Fdashboards#/view/cee6d680-16c0-11ed-8f9a-d7ff1d0d43e3?_g=(filters:!(),refreshInterval:(pause:!t,value:0),time:(from:now-15m,to:now))`
// 	// fmt.Printf("uid is %s\n",uid)
// 	if err := chromedp.Run(ctx,kibanaElementScreenshotWithAuth(url, `div.dashboardViewport`,&buf)); err != nil {
// 		log.Fatal(err)
// 	}
// 	file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/test.png")
// 	if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
// 		log.Fatal(err)
// 	}

	
// }




////--------sample------------------////
// elementScreenshot takes a screenshot of a specific element.
func elementScreenshotwithauth1(loginUrl, sel string, res *[]byte) chromedp.Tasks {

	var executed *runtime.RemoteObject
	username:="elastic"
	password:="12345678"

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
		chromedp.WaitVisible(`div.dashboardViewport`),
		chromedp.Sleep(20 * time.Second),
		chromedp.Screenshot(sel, res, chromedp.NodeVisible),
	}
}
////--------sample------------------////