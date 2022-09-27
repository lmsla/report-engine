package services

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"report-backend-golang/entities"
	"report-backend-golang/global"

	// "report-backend-golang/services"

	// "time"
	"github.com/chromedp/chromedp"
	// "github.com/chromedp/cdproto/page"
)


func ScreenshotDocker(ReportID int) {


	ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://10.99.1.138:3003/devtools/page/BROWSERLESSGDV8E5H4386JYCYGVRBAQ")
	// ctx, cancel := chromedp.NewRemoteAllocator(context.Background(), "ws://10.99.1.138:9222/devtools/page")
	defer cancel()

	ctx, cancel = chromedp.NewContext(ctx,
		chromedp.WithDebugf(log.Printf),
	)
	defer cancel()

	// go get the current targets in the Electron app.
	ts, err := chromedp.Targets(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if len(ts) == 0 {
		log.Fatal("targets not found.")
	}

	ctx, cancel = chromedp.NewContext(ctx,
		// use the first target
		chromedp.WithTargetID(ts[0].TargetID),
	)

	// // create context
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()

	// 先取得 report 中包含的 instance
	data,err := GetInstanceInReportbyReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	reportinfo,err := GetReportByReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	
	//判斷 report 中 instance 的 type 是 kibana or grafana
	for _,data1 := range data{
		// Kibana 的區塊
		if data1.Type == "kibana" {
			// 取得屬於此 instance 的 dashboards
			dashboard,err := GetDashboardOfInstanceByInstanceID(data1.InstanceID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data1.InstanceID)
			for _,data2 := range dashboard {
				// create context 截圖程式碼，擺著就好勿動
				// ctx, cancel := chromedp.NewContext(
				// 	context.Background(),
				// 	// chromedp.WithDebugf(log.Printf),
				// )
				// defer cancel()
				// 用 report 中 instance 的 instance ID 去 instance 資料庫中撈出 instance 的 IP,type,user,password
				auth,err := GetInstanceTypeAndIP(data1.InstanceID)
				if err != nil {
					fmt.Println(err)
				}
				// capture screenshot of an element 截圖程式碼，擺著就好勿動
				var buf []byte
				// 將取出來的個參數帶入網址中以便截圖
				url := fmt.Sprintf("%s/app/dashboards#/view/%s?_g=(time:(from:%s,to:%s))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)",auth["IP"],data2.UID,reportinfo.From,reportinfo.To)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,kibanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.dashboardViewport`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("%s/%s.png",global.EnvConfig.Reportengine.PicturePath,data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
			}	
		}else if data1.Type == "grafana" {
			dashboard,err := GetDashboardOfInstanceByInstanceID(data1.InstanceID)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(data1.InstanceID)
			for _,data2 := range dashboard {
				// create context 截圖程式碼，擺著就好勿動
				// ctx, cancel := chromedp.NewContext(
				// 	context.Background(),
				// 	// chromedp.WithDebugf(log.Printf),
				// )
				// defer cancel()
				// 用 report 中 instance 的 instance ID 去 instance 資料庫中撈出 instance 的 IP,type,user,password
				auth,err := GetInstanceTypeAndIP(data1.InstanceID)
				if err != nil {
					fmt.Println(err)
				}
				// capture screenshot of an element 截圖程式碼，擺著就好勿動
				var buf []byte
				// 將取出來的個參數帶入網址中以便截圖
				url := fmt.Sprintf("%s/grafana_iframe/d/%s/?from=%s/d&to=%s&orgId=1",auth["IP"],data2.UID,reportinfo.From,reportinfo.To)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,grafanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.scrollbar-view`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("%s/%s.png",global.EnvConfig.Reportengine.PicturePath,data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
				
			}
			
		}

	}
	
}




func Screenshot(ReportID int) {

	// // create context
	// ctx, cancel := chromedp.NewContext(
	// 	context.Background(),
	// 	// chromedp.WithDebugf(log.Printf),
	// )
	// defer cancel()

	// 先取得 report 中包含的 instance
	data,err := GetInstanceInReportbyReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	reportinfo,err := GetReportByReportID(ReportID)
	if err != nil {
		fmt.Println(err)
	}
	
	//判斷 report 中 instance 的 type 是 kibana or grafana
	for _,data1 := range data{
		// Kibana 的區塊
		if data1.Type == "kibana" {
			// 取得屬於此 instance 的 dashboards
			dashboard,err := GetDashboardOfInstanceByInstanceID(data1.InstanceID)
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
				url := fmt.Sprintf("%s/app/dashboards#/view/%s?_g=(time:(from:%s,to:%s))&_a=(fullScreenMode:!f,options:(hidePanelTitles:!f,useMargins:!t),query:(language:lucene,query:''),tags:!(),timeRestore:!t,viewMode:view)",auth["IP"],data2.UID,reportinfo.From,reportinfo.To)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,kibanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.dashboardViewport`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("%s/%s.png",global.EnvConfig.Reportengine.PicturePath,data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
			}	
		}else if data1.Type == "grafana" {
			dashboard,err := GetDashboardOfInstanceByInstanceID(data1.InstanceID)
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
				url := fmt.Sprintf("%s/grafana_iframe/d/%s/?from=%s/d&to=%s&orgId=1",auth["IP"],data2.UID,reportinfo.From,reportinfo.To)
				// 帶入 url,instance 的 username 、 password 帶入截圖的 function
				if err := chromedp.Run(ctx,grafanaElementScreenshotWithAuth(url, auth["user"],auth["password"],`div.scrollbar-view`,&buf)); err != nil {
					log.Fatal(err)
				}
				// 寫到指定路徑，暫定以 dashboard UID 命名
				file := fmt.Sprintf("%s/%s.png",global.EnvConfig.Reportengine.PicturePath,data2.UID)
				if err := ioutil.WriteFile(file, buf, 0o644); err != nil {
					log.Fatal(err)
				}
				
			}
			
		}

	}
	
}




// 查 instance 的資訊
func GetInstanceInfoByInstanceID(InstanceID int) (entities.Instance, error){

	var instance entities.Instance
	instance.InstanceID = InstanceID

	err := global.Mysql.Where("instance_id= ?",InstanceID).First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil

}


// // 查單一Instance
// func GetInstanceByID(instanceID int) (entities.Instance, error) {

// 	var instance entities.Instance
// 	instance.InstanceID = instanceID
// 	err := global.Mysql.First(&instance).Error
// 	if err != nil {
// 		return instance, err
// 	}
// 	return instance, nil
// }


// 只取出 Instance 中需要的元素
func GetInstanceTypeAndIP(id int) (map[string]string, error){
	// var IPType []map[string]string
	data,err := GetInstanceByID(id)
	if err != nil {
		fmt.Println(err)
		// log.Logrecord("ERROR",err.Error())
		//return
	}
	// var UID1 []string
	IPType := make(map[string]string)
	IPType["IP"] = data.IP
	IPType["type"] = data.Type
	IPType["user"] = data.User
	IPType["password"] = data.Pass
	return IPType,err
}
// 查 Dashboard uid by report ID
func GetDashboardByReportID(reportID int) (entities.Report, error) {

	var instance entities.Report
	instance.ReportID = reportID
	err := global.Mysql.First(&instance).Error
	if err != nil {
		return instance, err
	}
	return instance, nil
}

// 查詢All dashboard
func GetAllDashboard() ([]entities.RDashboard, error) {

	var instancies []entities.RDashboard
	err := global.Mysql.Find(&instancies).Error
	if err != nil {
		return nil, err
	}

	return instancies, nil
}

func GetUID() ([]string,error){
	uid,err := GetAllDashboard()
	if err != nil {
		fmt.Println(err)
		// log.Logrecord("ERROR",err.Error())
		//return
	}
	var UID1 []string
	for _,uids := range uid {
		// var UID1 []string
		UID1 = append(UID1,uids.UID)
	}	
	fmt.Println(UID1)
	if err != nil {
		return nil, err
	}
	return UID1,nil

}