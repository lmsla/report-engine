package controller

import (
	"net/http"
	// "report-backend-golang/models"
	"report-backend-golang/services"
	// "report-backend-golang/handler"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)

// @Summary Get history
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /History/GetAll [get]
func GetAllHitory(c *gin.Context) {

	res := services.GetAllHistory()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}


// @Summary Get Old history
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /History/GetOldHistory [get]
func GetOldHitory(c *gin.Context) {

	res := services.GetOldHistory()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}