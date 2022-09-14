package screenshot

import (
	"context"
	"io/ioutil"
	"log"
	// "time"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/cdproto/page"

)

func Pdfcreate() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte
	
	// 导出指定元素为PDF
	if err := chromedp.Run(ctx,elementPDFPrint(`http://10.99.1.138:5601/app/dashboards#/view/d909e710-bb64-11ea-902b-f929e5152feb?_g=(filters:!())&_a=(description:'',filters:!(),fullScreenMode:!t,options:(darkTheme:!f,hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:taipei),tags:!(),timeRestore:!t,title:Apache,viewMode:view)`,`div.dashboardViewport`,&buf)); err != nil {
		log.Fatal(err)
	}
	if err := ioutil.WriteFile("elementScreenshot.pdf", buf, 0644); err != nil {
		log.Fatal(err)
	}
	log.Printf("wrote elementScreenshot.png and fullScreenshot.png and elementScreenshot.pdf")

}
