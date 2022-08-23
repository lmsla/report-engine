package services

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os/exec"
	"report-backend-golang/entities"
	"report-backend-golang/models"
)

func VerifyKibanaInstance(instance *entities.Instance) (models.Response) {
	var res models.Response
	client := &http.Client{}
	req, _ := http.NewRequest("GET", instance.IP+"/api/features", nil)
	req.SetBasicAuth(instance.User, instance.Pass)
	resq, err := client.Do(req)
	if err != nil {
		res.Msg = fmt.Sprintf("%v connection refused", instance.IP)
		res.Success = false
		return res
	}

	if 	resq.StatusCode == 200 {
		res.Msg = "Connection Successfully"
		res.Success = true
		return res

	} else if resq.StatusCode == 401 {
		res.Msg = "invalid username or password"
		res.Success = false
		return res

	} else {
		res.Msg = fmt.Sprintf("%v connection refused", instance.IP)
		res.Success = false
		return res
	}
		
}

func GetKibanaSpaces(instance entities.Instance) ([]string, error) {

	var curl *exec.Cmd
	if instance.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", instance.IP+"/api/spaces/space")
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", instance.User+":"+instance.Pass, "-s", instance.IP+"/api/spaces/space")
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

func GetKibanaDashboardData(inventory entities.Instance, space string) (string, error) {

	var curl *exec.Cmd
	if inventory.User == "" {
		curl = exec.Command("curl", "-XGET", "-k", "-s", inventory.IP+"/s/"+space+"/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	} else {
		curl = exec.Command("curl", "-XGET", "-k", "-u", inventory.User+":"+inventory.Pass, "-s", inventory.IP+"/s/"+space+"/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000") // 修改了此行
	}

	out, err := curl.Output()
	if err != nil {
		return "error", err
	}
	return string(out), nil

}

// func GetKibanaDashboardData() string {
//     curl := exec.Command("curl","-XGET","-k","-u","elastic:RnIv7YhigaVKS=l-*yz9","-s","http://10.99.1.110:5601/api/saved_objects/_find?type=dashboard&fields=id&fields=title&fields=description&per_page=10000")  // 修改了此行
//     out, err := curl.Output()
//     if err != nil {
//         fmt.Println("erorr", err)

//     }
// 	return string(out)

// }

func GetALLKibanaDashboardTitle(instance entities.Instance) ([]entities.Dashboard, error) {

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
			dashboard.DashboardName = name.(string)
			dashboard.UID = uid.(string)
			dashboards = append(dashboards, dashboard)
		}
	}
	return dashboards, nil
}

