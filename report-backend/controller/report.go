package controller

import (
	"net/http"
	"strconv"

	"report-backend-golang/models"
	"report-backend-golang/services"
	"report-backend-golang/handler"

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


// @Summary Get report by Report ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Report
// @Router /Report/GetReport/{id} [get]
// @Security ApiKeyAuth
func GetReportByReportID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetReportByReportID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
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



// @Summary Get Report by Schedule ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "schedule id"
// @Success 200 {object} entities.Schedule
// @Router /Report/GetReportByScheduleID/{id} [get]
// @Security ApiKeyAuth
func GetReportByScheduleID(c *gin.Context) {

	ScheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "schedule ID should be int")
		handler.WriteErrorLog(c, "schedule ID should be integer")
		return
	}


	inventory, err := services.GetReportByScheduleID(ScheduleID)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
	// --------------------
	// res := services.GetReportByScheduleID(ScheduleID)

	// if !res.Success {
	// 	c.JSON(http.StatusBadRequest, res.Msg)
	// 	return
	// }

	// c.JSON(http.StatusOK, res.Body)

	// --------------------
	// inventory, err := services.GetElementsByReportID(ReportID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "error to get inventory details")
	// 	handler.WriteErrorLog(c, "error to get inventory details")
	// 	return
	// }
	// c.JSON(http.StatusOK, inventory)

}