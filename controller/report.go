package controller

import (
	"net/http"
	"strconv"

	"report-backend-golang/models"
	"report-backend-golang/services"
	// "report-backend-golang/handler"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)


// @Summary Get Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Report/GetAll [get]
func GetAllReports(c *gin.Context) {

	res := services.GetAllReports()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}



// @Summary Create Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param Report body models.Report true "report"
// @Success 200 {object} models.Response
// @Router /Report/Create [post]
func CreateReport(c *gin.Context) {

	// body := new(models.Instance)
	body := new(models.Report)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.CreateReport(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}

// @Summary Update Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param Instacne body models.Report true "report"
// @Success 200 {object} string
// @Router /Report/Update [put]
func UpdateReport(c *gin.Context) {

	body := new(models.Report)

	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.UpdateReport(*body)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Body)
}


// @Summary Delete Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /Report/Delete/{id} [delete]
func DeleteReport(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}

	res := services.DeleteReport(id)

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}