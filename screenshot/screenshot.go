package screenshot

import (
	"context"
	"io/ioutil"
	"log"
	"fmt"
	// "time"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/cdproto/page"

)

func Screenshot() {
	// create context
	ctx, cancel := chromedp.NewContext(
		context.Background(),
		// chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	UID,err := GetUID()
	if err != nil {
		fmt.Println(err)
	}
	
	for _,uid := range UID {
		url := fmt.Sprintf("http://10.99.1.138:5601/app/dashboards#/view/%s?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)",uid)
		// fmt.Printf("uid is %s\n",uid)
		if err := chromedp.Run(ctx,elementScreenshot(url, `div.dashboardViewport`,&buf)); err != nil {
			log.Fatal(err)
		}
		file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/dashboardshot/%s.png",uid)
		if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
			log.Fatal(err)
		}

	}




	// if err := chromedp.Run(ctx,elementScreenshot(`http://10.99.1.138:5601/app/dashboards#/view/1a93dd50-8863-11eb-beca-e57e3c411975?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:%22warning%22'),tags:!(),timeRestore:!t,viewMode:view)`, `div.dashboardViewport`,&buf)); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := ioutil.WriteFile("elementScreenshot.png", buf, 0o644); err != nil {
	// 	log.Fatal(err)
	// }


//-------------------------備份---------------------------------

// capture entire browser viewport, returning png with quality=90
//	if err := chromedp.Run(ctx, fullScreenshot(`https://brank.as/`, 90, &buf)); err != nil {
//		log.Fatal(err)
//	}
//	if err := ioutil.WriteFile("fullScreenshot.png", buf, 0o644); err != nil {
//		log.Fatal(err)
//	}

	//log.Printf("wrote elementScreenshot.png and fullScreenshot.png")
	
	// // 导出指定元素为PDF
	// if err := chromedp.Run(ctx,elementPDFPrint(`http://10.99.1.138:5601/app/dashboards#/view/d909e710-bb64-11ea-902b-f929e5152feb?_g=(filters:!())&_a=(description:'',filters:!(),fullScreenMode:!t,options:(darkTheme:!f,hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:taipei),tags:!(),timeRestore:!t,title:Apache,viewMode:view)`,`div.dashboardViewport`,&buf)); err != nil {
	// 	log.Fatal(err)
	// }
	// if err := ioutil.WriteFile("elementScreenshot.pdf", buf, 0644); err != nil {
	// 	log.Fatal(err)
	// }
	// log.Printf("wrote elementScreenshot.png and fullScreenshot.png and elementScreenshot.pdf")

}

