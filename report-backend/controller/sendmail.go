package controller


import (
	"net/http"
	"strconv"
	"report-backend-golang/services"
	// "report-backend-golang/handler"
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
		// handler.WriteErrorLog(c, "report ID should be integer")
		return
	}

	services.SendEmailBySchedule(ScheduleID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}