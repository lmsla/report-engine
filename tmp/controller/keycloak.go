package controller

import (
	"context"
	"crypto/tls"
	"fmt"
	"net/http"
	"report-backend-golang/global"
	"report-backend-golang/models"
	"github.com/Nerzal/gocloak/v13"
	"github.com/gin-gonic/gin"
	// "go.uber.org/zap"
	// "golang.org/x/exp/slices"
)

// Keycloak [Token 驗證]
func GetUserInfo(c *gin.Context) {

	// 取 token 驗證
	tokens := c.Request.Header["Authorization"]

	if len(tokens) == 0 {
		c.AbortWithStatusJSON(http.StatusUnauthorized, "authorization token is required")
		return
	}
	token := tokens[0]

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
		
		fmt.Println("GetUserInfo: "+ err.Error())
		fmt.Printf("realm: %v\n", realm)

		return
	}
	user := models.SSOUser{}
	user.IsAdmin = false
	user.ID = *(userInfo.Sub)
	user.Name = *(userInfo.PreferredUsername)

	fmt.Println("用戶資訊 ===>")
	fmt.Printf("UserID: 	  %v\n", user.ID)
	fmt.Printf("Realm: 	          %v\n", realm)
	fmt.Printf("Name: 		  %v\n", user.Name)
	fmt.Printf("IsAdmin: 	  %v\n", user.IsAdmin)
	fmt.Printf("Groups: 	  %+v\n", user.Groups)
	c.Set("user", user)
	c.Next()
}
