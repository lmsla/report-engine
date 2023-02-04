package controller

import (
	"net/http"
	"strconv"

	// "report-backend-golang/models"
	"report-backend-golang/services"
	"report-backend-golang/handler"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)




// @Summary Get Space of Instance by Instance ID
// @Tags DropDown
// @Accept  json
// @Produce  json
// @Param id path int true "instance id"
// @Success 200 {object}  entities.Dropdown
// @Router /DropDown/Get/{id} [get]
// @Security ApiKeyAuth
func GetSpaceByInstanceID(c *gin.Context) {

	instanceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "instance ID should be int")
		handler.WriteErrorLog(c, "instance ID should be integer")
		return
	}

	inventory, err := services.GetInstanceByID(instanceID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}

	switch inventory.Type {
	case "grafana":
		r, err := services.GetAllGrafanaDashboardTitle(inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			handler.WriteErrorLog(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, r)
	case "kibana":
		r, err := services.GetALLKibanaDashboardTitle(inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			// handler.WriteErrorLog(c, err.Error())
			return
		}
		c.JSON(http.StatusOK, r)
	default:
		c.JSON(http.StatusBadRequest, "instance type unknown")
	}
}