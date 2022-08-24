package controller

import (
	// "fmt"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)



// @Summary Create Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param instance body entities.Schedule true "new schedule"
// @Success 200 {object} models.Response
// @Router /api/v1/Schedule/Create [post]
// @Security ApiKeyAuth
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