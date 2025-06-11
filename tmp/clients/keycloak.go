package clients

import (
	"context"
	"crypto/tls"
	// "encoding/json"
	"fmt"
	"report-backend-golang/global"
	// "report-backend-golang/models"
	// "os"
	// "strings"

	"github.com/Nerzal/gocloak/v13"
	// "go.uber.org/zap"
)


// Keycloak
func LoadKeycloak() {
	url := global.EnvConfig.SSO.Url
	client := gocloak.NewClient(url)
	restyClient := client.RestyClient()
	restyClient.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})

	ctx := context.Background()
	user := global.EnvConfig.SSO.User
	password := global.EnvConfig.SSO.Password
	realm := global.EnvConfig.SSO.Realm


	_, err := client.LoginAdmin(ctx, user, password, realm)
	if err != nil {
		// global.Logger.Error(
		// 	fmt.Sprintf("keycloak [LoginAdmin] error: [%v]", err.Error()),
		// 	zap.String(global.LogEvent.Tag.Service, global.LogEvent.Keycloak.Query),
		// )

		fmt.Println("keycloak error",err.Error())
	}

	// global.Logger.Info(
	// 	"keycloak connection success",
	// 	zap.String(global.LogEvent.Tag.Service, global.LogEvent.Keycloak.Connection),
	// )
}
