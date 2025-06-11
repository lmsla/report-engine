package controller

import (
	"net/http"
	"report-backend-golang/services"
	"strconv"
	"time"
	"report-backend-golang/handler"
	// "report-backend-golang/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Mail
// @Tags Mail
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /Mail/Send/{id} [post]
// @Security ApiKeyAuth
func SendEmailBySchedule(c *gin.Context) {

	ScheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Schedule ID should be int")
		handler.WriteErrorLog(c, "Schedule ID should be integer")
		return
	}
	nowtime := time.Now().Unix()
	services.SendEmailBySchedule(nowtime,ScheduleID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}