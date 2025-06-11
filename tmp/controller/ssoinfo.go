package controller

import (

	"github.com/gin-gonic/gin"
	"net/http"
	"report-backend-golang/global"
	"report-backend-golang/models"
)

// @Summary Get SSOUrl
// @Tags SSO
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /get-sso-url [get]
func GetSsoURL(c *gin.Context) {

	res := GetSsoUrl()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

func GetSsoUrl() models.Response {
	res := models.Response{}
	res.Success = false
	res.Body = global.EnvConfig.SSO.Url
	// fmt.Println(global.EnvConfig.SSO.SsoUrl)
	// fmt.Println(res)
	res.Success = true
	res.Msg = "Get All SSO URL Success"
	return res
}
