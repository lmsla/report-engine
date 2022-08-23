package services

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
	"report-backend-golang/entities"
	"report-backend-golang/global"
)

func GetUser(token string) (map[string]interface{}, error) {
	// use token to get user info from sso
	url := fmt.Sprintf("http://%v/api/v1/user/get", global.EnvConfig.Other.Backend)
	fmt.Printf("url: %v\n", url)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var user map[string]interface{}
	json.Unmarshal(body, &user)

	return user, nil
}

func GetUserAllFromSSO(token string) ([]entities.User, error) {
	// use token to get user info from sso
	url := fmt.Sprintf("http://%v/api/v1/user/get-all", global.EnvConfig.Other.Backend)
	fmt.Printf("url: %v\n", url)
	req, err := http.NewRequest("POST", url, nil)	
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", token)

	client := &http.Client{Timeout: time.Second * 10}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	
	var data []map[string]interface{}
	var res entities.User
	var user_list  []entities.User
	json.Unmarshal(body, &data)

	for _, user := range data {		
		res.UserID =  int(user["id"].(float64))
		res.UserName = user["nickname"].(string)
		user_list = append(user_list, res)
	}
	return user_list, nil
}