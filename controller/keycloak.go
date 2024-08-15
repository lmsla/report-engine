package controller

import (
	"context"
	"crypto/tls"
	// "encoding/json"
	"fmt"
	"net/http"
	// "os"
	"report-backend-golang/global"
	"report-backend-golang/models"
	// "strings"

	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
	// "golang.org/x/exp/slices"
)

// Keycloak [Token 驗證]
func GetUserInfo(c *gin.Context) {

	// 取 token 驗證
	tokens := c.Request.Header["Authorization"]

	fmt.Println("tokens",tokens)

	if len(tokens) == 0 {
		// c.JSON(http.StatusUnauthorized, "auhorization token is required")
		c.AbortWithStatusJSON(http.StatusUnauthorized, "authorization token is required")
		return
	}
	token := tokens[0]

	fmt.Println("token",tokens[0])

	// // 取 realm 驗證
	// realms := "master"
	// // realms := c.Request.Header["Realm"]
	// if len(realms) == 0 {
	// 	c.JSON(http.StatusNotFound, "realm is empty")
	// 	return
	// }
	realm := "master"

	// 設定參數
	ctx := context.Background()
	url := global.EnvConfig.SSO.Url
	client := gocloak.NewClient(url)
	restyClient := client.RestyClient()
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	// 取 UserInfo
	userInfo, err := client.GetUserInfo(ctx, token, realm)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "GetUserInfo: "+ err.Error())
		return
	}
	user := models.SSOUser{}
	user.IsAdmin = false
	user.ID = *(userInfo.Sub)
	user.Name = *(userInfo.PreferredUsername)
	// user.Realm = realm

	// // 取 user group + role
	// params := gocloak.GetGroupsParams{
	// 	Full: gocloak.BoolP(true),
	// }
	// userGroups, err := client.GetUserGroups(ctx, token, realm, *userInfo.Sub, params)
	// if err != nil {
	// 	// global.Logger.Error(
	// 	// 	err.Error(),
	// 	// 	zap.String(global.LogEvent.Tag.Service, global.LogEvent.Keycloak.Query),
	// 	// )
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	return
	// }
	// fmt.Println("userGroups",userGroups)
	// var role_list []string
	// for _, group := range userGroups {
	// 	var g models.RealmGroupRes
	// 	g.ID = *group.ID
	// 	g.Name = *group.Name
	// 	fmt.Println("role_list",role_list)
	// 	fmt.Println("group",group)
	// 	role_list = append(role_list, *group.RealmRoles...)
	// 	user.Groups = append(user.Groups, g)

	// }

	// // 按 group mapping roles 判斷有沒有 admin 權限
	// user.Roles = role_list
	// if slices.Contains(role_list, global.EnvConfig.SSO.AdminRole) {
	// 	user.IsAdmin = true
	// }

	fmt.Println("用戶資訊 ===>")
	fmt.Printf("UserID: 	  %v\n", user.ID)
	fmt.Printf("Realm: 	          %v\n", realm)
	fmt.Printf("Name: 		  %v\n", user.Name)
	fmt.Printf("IsAdmin: 	  %v\n", user.IsAdmin)
	// fmt.Printf("Roles: 		  %v\n", strings.Join(role_list, ", "))
	fmt.Printf("Groups: 	  %+v\n", user.Groups)

	c.Set("user", user)
	c.Next()
}
