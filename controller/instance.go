package controller

import (
	"fmt"
	"net/http"
	"strconv"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param instance body entities.Instance true "new instance"
// @Success 200 {object} models.Response
// @Router /api/v1/Instance/Create [post]
// @Security ApiKeyAuth
func CreateInstance(c *gin.Context) {

	body := new(entities.Instance)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateInstance(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get All Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Instance
// @Router /api/v1/Instance/GetAll [get]
// @Security ApiKeyAuth
func GetAllInstance(c *gin.Context) {

	// Find DB
	r, err := services.GetAllInstance()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get Dashboard of server by Instance ID
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "instance id"
// @Success 200 {object} entities.Dashboard
// @Router /api/v1/Instance/GetDashboards/{id} [get]
// @Security ApiKeyAuth
func GetDBDashboardByInstanceID(c *gin.Context) {

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

// @Summary Update Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param instance body entities.Instance true "instance"
// @Param id path int true "instance id"
// @Success 200 {object} models.Response
// @Router /api/v1/Instance/Update/{id} [put]
// @Security ApiKeyAuth
func UpdateInstance(c *gin.Context) {

	instanceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "instance id should be int")
		handler.WriteErrorLog(c, "instance id should be integer")
		return
	}

	body := new(entities.Instance)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		// handler.WriteErrorLog(c, err.Error())
	}

	// Update DB
	r := services.UpdateInstance(instanceID, *body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Delete Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "instance id"
// @Success 200 {object} models.Response
// @Router /api/v1/Instance/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteInstance(c *gin.Context) {

	instanceID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "instance id should be int")
		handler.WriteErrorLog(c, "instance id should be integer")
		return
	}

	// Update DB
	r := services.DeleteInstance(instanceID)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Verify Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param instance body entities.Instance true "new instance"
// @Success 200 {object} models.Response
// @Router /api/v1/Instance/Verify [post]
// @Security ApiKeyAuth
func VerifyInstance(c *gin.Context) {

	instance := new(entities.Instance)
	err := c.Bind(&instance)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	fmt.Printf("instance.Type: %v\n", instance.Type)
	switch instance.Type{
	case "grafana":
		res := services.VerifyGrafanaInstance(instance)
		if !res.Success {
			c.JSON(http.StatusBadRequest, res.Msg)
			handler.WriteErrorLog(c, res.Msg)
			return
		}
		
		c.JSON(http.StatusOK, res)
	case "kibana":
		res := services.VerifyKibanaInstance(instance)
		if !res.Success {
			c.JSON(http.StatusBadRequest, res.Msg)
			handler.WriteErrorLog(c, res.Msg)
			return
		}
		c.JSON(http.StatusOK, res)
	default:
		c.JSON(http.StatusBadRequest, "instance type unknown")
	}
}



	
