package controller

import (
	"net/http"
	// "strconv"
	"report-backend-golang/entities"
	"report-backend-golang/services"
	// "report-backend-golang/handler"
	"github.com/gin-gonic/gin"
)


// @Summary Create DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param DataTable body entities.DataTable true "DataTable"
// @Success 200 {object} models.Response
// @Router /DataTable/Create [post]
// @Security ApiKeyAuth
func CreateDataTable(c *gin.Context) {

	body := new(entities.DataTable)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.CreateDataTable(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}