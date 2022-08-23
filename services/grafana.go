package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"report-backend-golang/entities"
	"report-backend-golang/models"

	gapi "github.com/grafana/grafana-api-golang-client"
)

func VerifyGrafanaInstance(instance *entities.Instance) (models.Response) {
	var res models.Response

	client := &http.Client{}
	req, _ := http.NewRequest("GET", instance.IP+"/api/users", nil)
	req.SetBasicAuth(instance.User, instance.Pass)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resq, err := client.Do(req)
	if err != nil {
		res.Msg = fmt.Sprintf("%v connection refused", instance.IP)
		res.Success = false
		return res
	}

	// invalid username or password
	var load_json map[string]interface{}
	if resq.StatusCode != 200 {
		byte_json, _ := ioutil.ReadAll(resq.Body)
		json.Unmarshal(byte_json, &load_json)
		message := load_json["message"].(string)
		res.Msg = message
		res.Success = false
		return res
	}
	
	res.Msg = "Connection Successfully"
	res.Success = true
	return res
}

// 連線
func ConnGapi(baseURL, user, pass string) (*gapi.Client, error) {

	client, err := gapi.New(baseURL, gapi.Config{
		BasicAuth: url.UserPassword(user, pass),
	})

	if err != nil {
		return nil, err
	}

	return client, nil
}

// 取得所有 Dashboard 標題名稱
func GetAllGrafanaDashboardTitle(instance entities.Instance) ([]entities.Dashboard, error) {

	conn, err := ConnGapi(instance.IP, instance.User, instance.Pass)
	if err != nil {
		return nil, err
	}

	res, _ := conn.Dashboards()
	var dashboards []entities.Dashboard
	for _, fdsr := range res {
		var dashboard entities.Dashboard
		dashboard.DashboardName = fdsr.Title
		dashboard.UID = fdsr.UID
		dashboards = append(dashboards, dashboard)
	}

	return dashboards, nil
}
