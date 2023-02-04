package services

import (
	"encoding/json"
	"fmt"
	// "net/http"
	// "os/exec"
	"report-backend-golang/entities"
	"report-backend-golang/models"
)




func GetSpace(instance models.Instance) ([]entities.Dashboard, error) {

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
			dashboards = append(dashboards, dashboard)
		}
		
	}
	return dashboards, nil
}