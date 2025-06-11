package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"report-backend-golang/entities"
	"report-backend-golang/models"
	"strings"
)

func VerifyKibanaInstance(instance *entities.Instance) models.Response {
	var res models.Response
	client := &http.Client{}
	req, _ := http.NewRequest("GET", instance.URL+"/api/features", nil)
	req.SetBasicAuth(instance.User, instance.Password)
	resq, err := client.Do(req)
	if err != nil {
		res.Msg = fmt.Sprintf("%v connection refused", instance.URL)
		res.Success = false
		return res
	}

	if resq.StatusCode == 200 {
		res.Msg = "Connection Successfully"
		res.Success = true
		return res

	} else if resq.StatusCode == 401 {
		res.Msg = "invalid username or password"
		res.Success = false
		return res

	} else {
		res.Msg = fmt.Sprintf("%v connection refused", instance.URL)
		res.Success = false
		return res
	}

}

func GetKibanaSpaces1(instance models.Instance) ([]entities.Dropdown, error) {

	var curl *exec.Cmd
	if instance.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", instance.URL+"/api/spaces/space")
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", instance.User+":"+instance.Password, "-s", instance.URL+"/api/spaces/space")
	}

	out, err := curl.Output()
	if err != nil {
		return nil, err
	}

	var mapResult []map[string]interface{}
	err = json.Unmarshal([]byte(string(out)), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	var dropdownDatas []entities.Dropdown
	for i := 0; i < len(mapResult); i++ {
		dropdownData := new(entities.Dropdown)
		dropdownData.Text = mapResult[i]["id"].(string)
		dropdownData.Value = mapResult[i]["id"].(string)
		dropdownDatas = append(dropdownDatas, *dropdownData)
	}

	return dropdownDatas, nil
}

func GetKibanaSpaces(instance models.Instance) ([]string, error) {

	var curl *exec.Cmd
	if instance.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", instance.URL+"/api/spaces/space")
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", instance.User+":"+instance.Password, "-s", instance.URL+"/api/spaces/space")
	}

	out, err := curl.Output()
	if err != nil {
		return nil, err
	}

	var mapResult []map[string]interface{}
	err = json.Unmarshal([]byte(string(out)), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	var spaces []string
	for i := 0; i < len(mapResult); i++ {
		spaces = append(spaces, mapResult[i]["id"].(string))
	}

	return spaces, nil
}

func GetKibanaDashboardData(inventory models.Instance, space string) (string, error) {

	var curl *exec.Cmd
	if inventory.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", inventory.URL+"/s/"+space+"/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", inventory.User+":"+inventory.Password, "-s", inventory.URL+"/s/"+space+"/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	}
	// curl -k -u elastic:12345678 -X GET 10.99.1.242:5601/kibana_iframe/s/default/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000

	out, err := curl.Output()
	if err != nil {
		return "error", err
	}
	return string(out), nil

}

func GetKibanaVisualizationData(inventory models.Instance, space string) (string, error) {

	var curl *exec.Cmd
	if inventory.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", inventory.URL+"/s/"+space+"/api/saved_objects/_find?type=visualization&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", inventory.User+":"+inventory.Password, "-s", inventory.URL+"/s/"+space+"/api/saved_objects/_find?type=visualization&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	}

	out, err := curl.Output()
	if err != nil {
		return "error", err
	}
	return string(out), nil

}

func GetALLKibanaDashboardTitle(instance models.Instance) ([]entities.Dashboard, error) {

	var dashboards []entities.Dashboard
	spaces, err := GetKibanaSpaces(instance)
	if err != nil {
		return nil, err
	}
	for _, space := range spaces {
		var mapResult map[string]interface{}
		res, _ := GetKibanaDashboardData(instance, space)
		err := json.Unmarshal([]byte(res), &mapResult)
		if err != nil {
			fmt.Println("JsonToMapDemo err: ", err)
			return nil, err
		}

		data := mapResult["saved_objects"]

		for i := range mapResult["saved_objects"].([]interface{}) {
			var dashboard entities.Dashboard
			name := data.([]interface{})[i].(map[string]interface{})["attributes"].(map[string]interface{})["title"]
			uid := data.([]interface{})[i].(map[string]interface{})["id"]
			dashboard.Name = name.(string)
			dashboard.UID = uid.(string)
			dashboard.InstanceID = instance.ID
			dashboards = append(dashboards, dashboard)
		}

	}
	return dashboards, nil
}

func GetALLKibanaVisualizationTitle(instance models.Instance) ([]entities.Visualization, error) {

	var dashboards []entities.Visualization
	spaces, err := GetKibanaSpaces(instance)
	if err != nil {
		return nil, err
	}
	for _, space := range spaces {
		var mapResult map[string]interface{}
		res, _ := GetKibanaVisualizationData(instance, space)
		err := json.Unmarshal([]byte(res), &mapResult)
		if err != nil {
			fmt.Println("JsonToMapDemo err: ", err)
			return nil, err
		}

		data := mapResult["saved_objects"]

		for i := range mapResult["saved_objects"].([]interface{}) {
			var dashboard entities.Visualization
			name := data.([]interface{})[i].(map[string]interface{})["attributes"].(map[string]interface{})["title"]
			uid := data.([]interface{})[i].(map[string]interface{})["id"]
			dashboard.Name = name.(string)
			dashboard.UID = uid.(string)
			dashboard.InstanceID = instance.ID
			dashboards = append(dashboards, dashboard)
		}

	}
	return dashboards, nil
}

func GetALLKibanaDashboardTitle1(space string, instance models.Instance) ([]entities.Dropdown, error) {

	var dropdownDatas []entities.Dropdown
	var mapResult map[string]interface{}
	res, _ := GetKibanaDashboardData(instance, space)
	err := json.Unmarshal([]byte(res), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	data := mapResult["saved_objects"]

	for i := range mapResult["saved_objects"].([]interface{}) {
		name := data.([]interface{})[i].(map[string]interface{})["attributes"].(map[string]interface{})["title"]
		uid := data.([]interface{})[i].(map[string]interface{})["id"]
		dropdownData := new(entities.Dropdown)
		dropdownData.Text = name.(string)
		dropdownData.Value = uid.(string)
		dropdownDatas = append(dropdownDatas, *dropdownData)
	}
	return dropdownDatas, nil

}

func GetALLKibanaVisualizationTitle1(space string, instance models.Instance) ([]entities.Dropdown, error) {

	var dropdownDatas []entities.Dropdown
	var mapResult map[string]interface{}
	res, _ := GetKibanaVisualizationData(instance, space)
	err := json.Unmarshal([]byte(res), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	data := mapResult["saved_objects"]

	for i := range mapResult["saved_objects"].([]interface{}) {
		// var dashboard entities.Visualization
		dropdownData := new(entities.Dropdown)
		name := data.([]interface{})[i].(map[string]interface{})["attributes"].(map[string]interface{})["title"]
		uid := data.([]interface{})[i].(map[string]interface{})["id"]
		dropdownData.Text = name.(string)
		dropdownData.Value = uid.(string)
		dropdownDatas = append(dropdownDatas, *dropdownData)
	}
	return dropdownDatas, nil
}

func GetKibanaDataViews(space string, inventory models.Instance) ([]entities.Dropdown, error) {

	var dropdownDatas []entities.Dropdown

	var curl *exec.Cmd
	if inventory.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", inventory.URL+"/s/"+space+"/api/data_views")
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", inventory.User+":"+inventory.Password, "-s", inventory.URL+"/s/"+space+"/api/data_views")
	}
	fmt.Println(curl)

	out, err := curl.Output()
	if err != nil {
		return nil, err
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(string(out)), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	data := mapResult["data_view"]

	for i := range mapResult["data_view"].([]interface{}) {
		// var dashboard entities.Visualization
		dropdownData := new(entities.Dropdown)
		name := data.([]interface{})[i].(map[string]interface{})["title"]
		uid := data.([]interface{})[i].(map[string]interface{})["id"]
		dropdownData.Text = name.(string)
		dropdownData.Value = uid.(string)
		dropdownDatas = append(dropdownDatas, *dropdownData)
	}
	return dropdownDatas, nil
}

/// <kibana host>:<port>/s/<space_id>/api/data_views/data_view/<id>

func GetDataViewData(space string, inventory models.Instance, uid string) ([]entities.Dropdown, error) {

	var curl *exec.Cmd
	if inventory.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", inventory.URL+"/s/"+space+"/api/data_views/data_view/"+uid)
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", inventory.User+":"+inventory.Password, "-s", inventory.URL+"/s/"+space+"/api/data_views/data_view/"+uid)
	}

	out, err := curl.Output()
	if err != nil {
		return nil, err
	}

	var mapResult map[string]interface{}
	err = json.Unmarshal([]byte(string(out)), &mapResult)
	if err != nil {
		fmt.Println("JsonToMapDemo err: ", err)
		return nil, err
	}

	// fmt.Println(mapResult)

	data := mapResult["data_view"].(map[string]interface{})["fields"]

	var dropdownDatas []entities.Dropdown
	// 	遞迴 data_view 中的 fields 部分
	for _, field := range data.(map[string]interface{}) {
		dropdownData := new(entities.Dropdown)
		// 獲取每個 field 的 name 和 type
		name := field.(map[string]interface{})["name"]
		uid := field.(map[string]interface{})["type"]
		if strings.Contains(name.(string), ".keyword") {
			dropdownData.Text = uid.(string)
			dropdownData.Value = name.(string)
			dropdownDatas = append(dropdownDatas, *dropdownData)
		}

	}

	// fmt.Println("Dropdown Data:", dropdownDatas)
	return dropdownDatas, nil
}
