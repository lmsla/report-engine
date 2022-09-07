package screenshot

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"report-backend-golang/services"
	// "time"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/cdproto/page"
)

func Screenshot(ReportID int) {

	// // create context
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()
	
	// 先取得 report 中包含的 instance
	data,err := services.GetInstanceInReportbyReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	//判斷 report 中 instance 的 type 是 kibana or grafana
	for _,data1 := range data{
		// Kibana 的區塊
		if data1.Type == "kibana" {
			// 取得屬於此 instance 的 dashboards
			dashboard,err := services.GetDashboardOfInstanceByInstanceID(data1.InstanceID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data1.InstanceID)
			for _,data2 := range dashboard {
				// create context 截圖程式碼，擺著就好勿動
				ctx, cancel := chromedp.NewContext(
					context.Background(),
					// chromedp.WithDebugf(log.Printf),
				)
				defer cancel()
				// 用 report 中 instance 的 instance ID 去 instance 資料庫中撈出 instance 的 IP,type,user,password
				auth,err := GetInstanceTypeAndIP(data1.InstanceID)
				if err != nil {
					fmt.Println(err)
				}
				// capture screenshot of an element 截圖程式碼，擺著就好勿動
				var buf []byte
				// 將取出來的個參數帶入網址中以便截圖
				url := fmt.Sprintf("%s/app/dashboards#/view/%s?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)",auth["IP"],data2.UID)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,kibanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.dashboardViewport`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/dashboardshot/%s.png",data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
			}	
		}else if data1.Type == "grafana" {
			dashboard,err := services.GetDashboardOfInstanceByInstanceID(data1.InstanceID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data1.InstanceID)
			for _,data2 := range dashboard {
				// create context 截圖程式碼，擺著就好勿動
				ctx, cancel := chromedp.NewContext(
					context.Background(),
					// chromedp.WithDebugf(log.Printf),
				)
				defer cancel()
				// 用 report 中 instance 的 instance ID 去 instance 資料庫中撈出 instance 的 IP,type,user,password
				auth,err := GetInstanceTypeAndIP(data1.InstanceID)
				if err != nil {
					fmt.Println(err)
				}
				// capture screenshot of an element 截圖程式碼，擺著就好勿動
				var buf []byte
				// 將取出來的個參數帶入網址中以便截圖
				url := fmt.Sprintf("%s/grafana_iframe/d/%s/?from=now-1h/d&to=now&orgId=1",auth["IP"],data2.UID)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,grafanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.scrollbar-view`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/dashboardshot/%s.png",data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
			}

		}

	}
	


	// capture screenshot of an element
	// var buf []byte

	// UID,err := GetUID()
	// if err != nil {
	// 	fmt.Println(err)
	// }
	
	// for _,uid := range UID {
	// 	url := fmt.Sprintf("http://10.99.1.138:5601/app/dashboards#/view/%s?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)",uid)
	// 	// fmt.Printf("uid is %s\n",uid)
	// 	if err := chromedp.Run(ctx,elementScreenshot(url, `div.dashboardViewport`,&buf)); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	file := fmt.Sprintf("/Users/chen/Documents/gitlab/git-out/product/report-backend/dashboardshot/%s.png",uid)
	// 	if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
	// 		log.Fatal(err)
	// 	}

	// }




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


// kibana url
// http://10.99.1.117:5601/app/dashboards#/view/7adfa750-4c81-11e8-b3d7-01146121b73d?_g=(time:(from:'2022-03-15T00:00:00.000Z',to:'2022-03-17T00:42:06.662Z'))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:'level:warning'),tags:!(),timeRestore:!t,viewMode:view)

// grafana url
//http://10.99.1.240:3000/grafana_iframe/d/mOTd6X67k/?from=now-1h%2Fd&to=now&orgId=1