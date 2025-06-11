package controller

import (
	"net/http"
	"report-backend-golang/entities"
	"strconv"
	// "report-backend-golang/models"
	"github.com/gin-gonic/gin"
	"report-backend-golang/handler"
	"report-backend-golang/services"
)

// @Summary Get Element
// @Tags Element
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Element/GetAll [get]
// @Security ApiKeyAuth
func GetAllElements(c *gin.Context) {

	res := services.GetAllElements()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

// @Summary Create Element
// @Tags Element
// @Accept  json
// @Produce  json
// @Param Element body entities.Element true "element"
// @Success 200 {object} models.Response
// @Router /Element/Create [post]
// @Security ApiKeyAuth
func CreateElement(c *gin.Context) {

	body := new(entities.Element)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.CreateElement(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}

// @Summary Update Element
// @Tags Element
// @Accept  json
// @Produce  json
// @Param Instacne body entities.Element true "instance"
// @Success 200 {object} string
// @Router /Element/Update [put]
// @Security ApiKeyAuth
func UpdateElement(c *gin.Context) {

	body := new(entities.Element)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.UpdateElement(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}

// @Summary Delete Element
// @Tags Element
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /Element/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteElement(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	res := services.DeleteElement(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Get Element by Report ID
// @Tags Element
// @Accept  json
// @Produce  json
// @Param id path int true "report id"
// @Success 200 {object} models.Element
// @Router /Element/GetElementByReportID/{id} [get]
// @Security ApiKeyAuth
func GetElementByReportID(c *gin.Context) {

	ReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "instance ID should be int")
		handler.WriteErrorLog(c, "instance ID should be integer")
		return
	}

	inventory, err := services.GetElementsByReportID(ReportID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)

	//-----------------
	// res := services.GetElementsByReportID(ReportID)

	// if !res.Success {
	// 	c.JSON(http.StatusBadRequest, res.Msg)
	// 	return
	// }

	// c.JSON(http.StatusOK, res.Body)
	//-----------------

	// inventory, err := services.GetElementsByReportID(ReportID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "error to get inventory details")
	// 	handler.WriteErrorLog(c, "error to get inventory details")
	// 	return
	// }
	// c.JSON(http.StatusOK, inventory)

}
