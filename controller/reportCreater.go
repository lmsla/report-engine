package controller

import (
	"net/http"
	"strconv"
	"report-backend-golang/services"
	// "report-backend-golang/handler"

	// "report-backend-golang/entities"
	"github.com/gin-gonic/gin"
)


// @Summary Html
// @Tags Html
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /Html/Create/{id} [post]
// @Security ApiKeyAuth
func CreateHtml(c *gin.Context) {

	ReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "report ID should be int")
		// handler.WriteErrorLog(c, "report ID should be integer")
		return
	}

	services.CreateHtml(ReportID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}


// @Summary PDF
// @Tags PDF
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /PDF/Create/{id} [post]
// @Security ApiKeyAuth
func CreatePDF(c *gin.Context) {

	ScheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "report ID should be int")
		// handler.WriteErrorLog(c, "report ID should be integer")
		return
	}

	services.CreatePDFbySchedule(ScheduleID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}
