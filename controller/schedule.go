package controller

import (
	// "fmt"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	// "report-backend-golang/schedule"
	"report-backend-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param instance body entities.Schedule true "new schedule"
// @Success 200 {object} models.Response
// @Router /api/v1/Schedule/Create [post]
func CreateSchedule(c *gin.Context) {

	body := new(entities.Schedule)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateSchedule(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}


// @Summary Get All Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Schedule
// @Router /api/v1/Schedule/GetAll [get]
// @Security ApiKeyAuth
func GetAllSchedule(c *gin.Context) {

	// Find DB
	r, err := services.GetAllSchedule()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}


// @Summary Update Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param schedule body entities.Schedule true "instance"
// @Param id path int true "schedule id"
// @Success 200 {object} models.Response
// @Router /api/v1/Schedule/Update/{id} [put]
// @Security ApiKeyAuth
func UpdateSchedule(c *gin.Context) {

	scheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "schedule id should be int")
		handler.WriteErrorLog(c, "schedule id should be integer")
		return
	}

	body := new(entities.Schedule)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		// handler.WriteErrorLog(c, err.Error())
	}

	// Update DB
	r := services.UpdateSchedule(scheduleID, *body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}



// @Summary Get schedule by Schedule ID
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Schedule
// @Router /api/v1/Schedule/GetSchedule/{id} [get]
// @Security ApiKeyAuth
func GetScheduleByScheduleID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetScheduleBysSheduleID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}



// @Summary Create Report by ScheduleID
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Schedule/ReportCreate/{id} [post]
// @Security ApiKeyAuth
func CreateReportbySchedule(c *gin.Context) {

	ScheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Schedule ID should be int")
		handler.WriteErrorLog(c, "Schedule ID should be integer")
		return
	}

	services.ExecuteShedulePDF(ScheduleID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}
