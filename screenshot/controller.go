package screenshot

import (
	"context"
	// "io/ioutil"
	// "log"
	"time"
	"github.com/chromedp/chromedp"
	"github.com/chromedp/cdproto/page"

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