package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"
	"strconv"
)

// @Summary Get DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Table/GetAll [get]
// @Security ApiKeyAuth
func GetTables(c *gin.Context) {

	res := services.GetTables()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

// @Summary Get table by DataTable ID
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Table
// @Router /Table/GetTable/{id} [get]
// @Security ApiKeyAuth
func GetTableByID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	inventory, err := services.GetTableByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// @Summary Create DataTable
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param Table body entities.Table true "Table"
// @Success 200 {object} models.Response
// @Router /Table/Create [post]
// @Security ApiKeyAuth
func CreateTable(c *gin.Context) {

	body := new(entities.Table)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, "error to Create DataTable")
		return
	}

	res := services.CreateTable(*body)

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
// @Param Table body entities.Table true "Table"
// @Success 200 {object} string
// @Router /Table/Update [put]
// @Security ApiKeyAuth
func UpdateTable(c *gin.Context) {

	body := new(entities.Table)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, "error to Update DataTable")
		return
	}

	res := services.UpdateTable(*body)

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
// @Router /Table/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteTable(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, "error to Delete DataTable")
		return
	}
	res := services.DeleteTable(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Get DataTable by Report ID
// @Tags DataTable
// @Accept  json
// @Produce  json
// @Param id path int true "report id"
// @Success 200 {object} entities.Table
// @Router /Table/GetTableByReportID/{id} [get]
// @Security ApiKeyAuth
func GetTableByReportID(c *gin.Context) {

	ReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Report ID should be int")
		handler.WriteErrorLog(c, "Report ID should be in")
		return
	}

	inventory, err := services.GetTableByReportID(ReportID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)

}