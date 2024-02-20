package controller

import (
	"fmt"
	"net/http"
	// "strconv"

	// "report-backend-golang/models"
	// "report-backend-golang/entities"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)

// @Summary Get Space of Instance by Instance ID
// @Tags Dropdown
// @Accept  json
// @Produce  json
// @Param source_type query string true "space,dashboard,visualization"
// @Param space_name query string false "space_name"
// @Param instance_id query int true "instance_id"
// @Success 200 {object}  entities.Dropdown
// @Security ApiKeyAuth
// @Router /Dropdown [get]
func GetDropdownSource(c *gin.Context) {


	body := new(entities.DropdownBody)
	c.Bind(&body)

	fmt.Println(body.InstanceID)
	fmt.Println(body.SourceType)
	fmt.Println(body.SpaceName)
	fmt.Println(body.InstanceID)

	// if err := services.GetDropdownSource(body); err != nil {

	// }
	switch body.SourceType {
	case "space":
		inventory, err := services.GetInstanceByID(body.InstanceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get inventory details")
			handler.WriteErrorLog(c, "error to get inventory details")
			return
		}
		spaces, err := services.GetKibanaSpaces1(inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get space details")
			handler.WriteErrorLog(c, "error to get space details")
			return
		}
		fmt.Println(spaces)
		c.JSON(http.StatusOK, spaces)

	case "dashboard":
		inventory, err := services.GetInstanceByID(body.InstanceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get inventory details")
			handler.WriteErrorLog(c, "error to get inventory details")
			return
		}
		dashboarddata, err := services.GetALLKibanaDashboardTitle1(body.SpaceName, inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get dashboard details")
			handler.WriteErrorLog(c, "error to get dashboard details")
			return
		}
		c.JSON(http.StatusOK, dashboarddata)

	case "visualization":
		inventory, err := services.GetInstanceByID(body.InstanceID)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get inventory details")
			handler.WriteErrorLog(c, "error to get inventory details")
			return
		}
		visualdata, err := services.GetALLKibanaVisualizationTitle1(body.SpaceName, inventory)
		if err != nil {
			c.JSON(http.StatusBadRequest, "error to get visualization details")
			handler.WriteErrorLog(c, "error to get visualization details")
			return
		}
		c.JSON(http.StatusOK, visualdata)

	}

}
