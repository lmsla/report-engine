package services

import (
	"encoding/json"
	"fmt"
	"io"
	// "os"
	"net/http"
	"net/url"
	"report-backend-golang/entities"
	"report-backend-golang/models"

	gapi "github.com/grafana/grafana-api-golang-client"
)

func VerifyGrafanaInstance(instance *entities.Instance) (models.Response) {
	var res models.Response

	client := &http.Client{}
	req, _ := http.NewRequest("GET", instance.URL+"/api/users", nil)
	req.SetBasicAuth(instance.User, instance.Password)
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")
	resq, err := client.Do(req)
	if err != nil {
		res.Msg = fmt.Sprintf("%v connection refused", instance.URL)
		res.Success = false
		return res
	}

	// invalid username or password
	var load_json map[string]interface{}
	if resq.StatusCode != 200 {
		byte_json, _ := io.ReadAll(resq.Body)
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
func GetAllGrafanaDashboardTitle(instance models.Instance) ([]entities.Element, error) {

	conn, err := ConnGapi(instance.URL, instance.User, instance.Password)
	if err != nil {
		return nil, err
	}

	res, _ := conn.Dashboards()
	var dashboards []entities.Element
	for _, fdsr := range res {
		var dashboard entities.Element
		dashboard.Name = fdsr.Title
		dashboard.UID = fdsr.UID

		dashboards = append(dashboards, dashboard)
	}

	return dashboards, nil
}