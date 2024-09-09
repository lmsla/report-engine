package controller

import (
	"net/http"
	"strconv"
	"report-backend-golang/entities"
	"report-backend-golang/services"
	// "report-backend-golang/handler"
	"github.com/gin-gonic/gin"
)


// @Summary Get DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /DataTable/GetAll [get]
// @Security ApiKeyAuth
func GetDataTables(c *gin.Context) { 

	res := services.GetDataTables()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}



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


// @Summary Update DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param DataTable body entities.DataTable true "DataTable"
// @Success 200 {object} string
// @Router /DataTable/Update [put]
// @Security ApiKeyAuth
func UpdateDataTable(c *gin.Context) {

	body := new(entities.DataTable)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.UpdateDataTable(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}


// @Summary Delete DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /DataTable/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteDataTable(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.DeleteDataTable(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}