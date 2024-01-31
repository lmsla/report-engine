package controller

// import (
// 	"net/http"
// 	"strconv"

// 	"report-backend-golang/services"
// 	// "report-backend-golang/handler"

// 	// "report-backend-golang/entities"
// 	"github.com/gin-gonic/gin"
// )


// // @Summary ScreenShot
// // @Tags ScreenShot
// // @Accept  json
// // @Produce  json
// // @Param id path int true "id"
// // @Success 200 {object} models.Response
// // @Router /Screenshot/Create/{id} [post]
// // @Security ApiKeyAuth
// func GetScreenShot(c *gin.Context) {

// 	ReportID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "report ID should be int")
// 		// handler.WriteErrorLog(c, "report ID should be integer")
// 		return
// 	}

// 	services.ScreenshotbyReport(ReportID)
// 	// if err != nil {
// 	// 	c.JSON(http.StatusBadRequest, err.Error())
// 	// 	handler.WriteErrorLog(c, err.Error())
// 	// 	return
// 	// }
// 	// c.JSON(http.StatusOK, r)
// }
