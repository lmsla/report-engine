package controller

import (
	"net/http"
	"report-backend-golang/services"
	"strconv"
	"time"
	"fmt"
	"report-backend-golang/handler"
	"report-backend-golang/log"
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
		handler.WriteErrorLog(c, "report ID should be integer")
		return
	}
	nowtime := time.Now().Unix()
	services.CreateHtml(nowtime, ReportID)
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
		c.JSON(http.StatusBadRequest, "Schedule ID should be int")
		handler.WriteErrorLog(c, "report ID should be integer")
		return
	}
	c.JSON(http.StatusOK, "報表試寄中，請稍候")
	
	log.Logrecord("排程",fmt.Sprintf("報表試寄, Schedule ID : %d",ScheduleID))

	go services.FuncAddToCron(ScheduleID)

	log.Logrecord("排程",fmt.Sprintf("報表試寄完成, Schedule ID : %d",ScheduleID))

}
