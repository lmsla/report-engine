package controller

import (
	"net/http"
	// "strconv"

	"report-backend-golang/services"
	// "report-backend-golang/handler"

	"report-backend-golang/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Get Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Schedule/GetAll [get]
func GetAllSchedule(c *gin.Context) {

	res := services.GetAllSchedule()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}


// @Summary Create Schedule
// @Tags Schedule
// @Accept  json
// @Produce  json
// @Param Schedule body entities.Schedule true "schedule"
// @Success 200 {object} models.Response
// @Router /Schedule/Create [post]
func CreateSchedule(c *gin.Context) {

	// body := new(models.Instance)
	body := new(entities.Schedule)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.CreateSchedule(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}