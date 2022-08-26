package controller

import (
	// "fmt"
	// "report-backend-golang/entities"
	"report-backend-golang/screenshot"
	// "strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get ScreenShot
// @Tags ScreenShot
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /api/v1/Screenshot/Create [post]
// @Security ApiKeyAuth
func GetScreenShot(c *gin.Context) {

	screenshot.Screenshot()
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}


