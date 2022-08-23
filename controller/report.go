package controller

import (
	// "fmt"
	"net/http"
	// "strconv"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)



// @Summary Create Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param instance body entities.Report true "new report"
// @Success 200 {object} models.Response
// @Router /api/v1/Report/Create [post]
// @Security ApiKeyAuth
func CreateReport(c *gin.Context) {

	body := new(entities.Report)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateReport(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get All Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Report
// @Router /api/v1/Report/GetAll [get]
// @Security ApiKeyAuth
func GetAllReport(c *gin.Context) {

	// Find DB
	r, err := services.GetAllReport()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}