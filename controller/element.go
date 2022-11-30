package controller

import (
	"net/http"
	"strconv"

	"report-backend-golang/entities"
	// "report-backend-golang/models"
	"report-backend-golang/services"

	// "report-backend-golang/handler"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)



// @Summary Get Element
// @Tags Element
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Element/GetAll [get]
func GetAllElements(c *gin.Context) {

	res := services.GetAllElements()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
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
func CreateElement(c *gin.Context) {

	body := new(entities.Element)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
func UpdateElement(c *gin.Context) {

	body := new(entities.Element)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
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
func DeleteElement(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.DeleteElement(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}