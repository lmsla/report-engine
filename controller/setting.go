package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"report-backend-golang/services"
		"report-backend-golang/handler"
)

// @Summary Get Server Module
// @Tags    Env
// @Accept  json
// @Produce json
// @Success 200 {object} []entities.Module
// @Router /get-server-module [get]
func GetServerModule(c *gin.Context) {
	res, err := services.GetServerModule()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary  Get Server Menu
// @Tags     Env
// @Accept   json
// @Produce  json
// @Success  200 {object} []entities.MainMenu
// @Security ApiKeyAuth
// @Router   /user/get-server-menu [get]
func GetServerMenu(c *gin.Context) {

	// user := c.Keys["user"].(models.SSOUser)

	// var roleName string
	// // 根據 realm 跟 role 給 conf
	// if user.IsAdmin {
	// 	roleName = global.EnvConfig.SSO.AdminRole
	// } else {
	// 	roleName = global.EnvConfig.SSO.UserRole
	// }
	res, err := services.GetServerMenu()

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, res)
}
