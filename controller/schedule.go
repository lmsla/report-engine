package controller

import (
	"net/http"
	"strconv"

	"report-backend-golang/services"
	"report-backend-golang/handler"

	"report-backend-golang/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Get Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Schedule/GetAll [get]
// @Security ApiKeyAuth
func GetAllSchedule(c *gin.Context) {

	res := services.GetAllSchedule()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

// @Summary Get Schedule by Schedule ID
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /Schedule/GetSchedule/{id} [get]
// @Security ApiKeyAuth
func GetScheduleByScheduleID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetScheduleBySheduleID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}



// @Summary Create Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param Schedule body entities.Schedule true "schedule"
// @Success 200 {object} models.Response
// @Router /Schedule/Create [post]
// @Security ApiKeyAuth
func CreateSchedule(c *gin.Context) {

	// body := new(models.Instance)
	body := new(entities.Schedule)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.CreateSchedule(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}


// @Summary Update Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param Schedule body entities.Schedule true "schedule"
// @Success 200 {object} string
// @Router /Schedule/Update [put]
// @Security ApiKeyAuth
func UpdateSchedule(c *gin.Context) {

	body := new(entities.Schedule)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.UpdateSchedule(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}




// @Summary Delete Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /Schedule/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteSchedule(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.DeleteSchedule(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}