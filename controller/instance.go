package controller

import (
	"net/http"
	"strconv"

	"report-backend-golang/models"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get Instance
// @Tags Instance
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Instance/GetAll [get]
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
func CreateInstance(c *gin.Context) {

	body := new(models.Instance)

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
