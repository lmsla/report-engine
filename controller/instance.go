package controller

import (
	"net/http"
	"strconv"
	"report-backend-golang/handler"
	"report-backend-golang/models"
	"report-backend-golang/services"
	"report-backend-golang/entities"
	"github.com/gin-gonic/gin"
)

// @Summary Get Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Instance/GetAll [get]
// @Security ApiKeyAuth
func GetAllInstances(c *gin.Context) {

	res := services.GetAllInstances()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

// @Summary Create Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param Instance body models.Instance true "instance"
// @Success 200 {object} models.Response
// @Router /Instance/Create [post]
// @Security ApiKeyAuth
func CreateInstance(c *gin.Context) {

	// body := new(models.Instance)
	body := new(entities.Instance)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.CreateInstance(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}

// @Summary Update Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param Instacne body models.Instance true "instance"
// @Success 200 {object} string
// @Router /Instance/Update [put]
// @Security ApiKeyAuth
func UpdateInstance(c *gin.Context) {

	body := new(models.Instance)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.UpdateInstance(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}

// @Summary Delete Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /Instance/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteInstance(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.DeleteInstance(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}


// @Summary Get Instance by Instance ID
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Instance
// @Router /Instance/GetInstance/{id} [get]
// @Security ApiKeyAuth
func GetInstanceByInstanceID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetInstanceByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get instance details")
		handler.WriteErrorLog(c, "error to get instance details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// @Summary Get Dashboard of server by Instance ID
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "instance id"
// @Success 200 {object} models.Element
// @Router /Instance/GetDashboards/{id} [get]
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


// @Summary Get Visualization of server by Instance ID
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "instance id"
// @Success 200 {object} models.Element
// @Router /Instance/GetVisualizations/{id} [get]
// @Security ApiKeyAuth
func GetDBVisualizationByInstanceID(c *gin.Context) {

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
		r, err := services.GetALLKibanaVisualizationTitle(inventory)
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



// @Summary check Instance by Instance ID
// @Tags Instance
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Instance
// @Router /Instance/CheckInstance/{id} [get]
// @Security ApiKeyAuth
func CheckInstanceByInstanceID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	res := services.CheckInstanceByID(id)
	c.JSON(http.StatusOK, res)
}