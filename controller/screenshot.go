package controller

import (
	// "fmt"
	// "report-backend-golang/entities"
	"net/http"
	"report-backend-golang/handler"
	// "report-backend-golang/screenshot"
	"report-backend-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Get ScreenShot
// @Tags ScreenShot
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Screenshot/Create/{id} [post]
// @Security ApiKeyAuth
func GetScreenShot(c *gin.Context) {

	ReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "report ID should be int")
		handler.WriteErrorLog(c, "report ID should be integer")
		return
	}

	services.Screenshot(ReportID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}


// // @Summary Create Html
// // @Tags ScreenShot
// // @Accept  json
// // @Produce  json
// // @Param id path int true "id"
// // @Success 200 {object} models.Response
// // @Router /api/v1/Html/Create/{id} [post]
// // @Security ApiKeyAuth
// func CreateHtml(c *gin.Context) {

// 	ReportID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "report ID should be int")
// 		handler.WriteErrorLog(c, "report ID should be integer")
// 		return
// 	}

// 	services.CreateHtml(ReportID)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, err.Error())
// 	// 	handler.WriteErrorLog(c, err.Error())
// 	// 	return
// 	// }
// 	// c.JSON(http.StatusOK, r)
// }


