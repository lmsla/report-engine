package controller

import (
	"net/http"
	"strconv"
	// "report-backend-golang/models"
	"github.com/gin-gonic/gin"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"
)

// @Summary Get Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /Report/GetAll [get]
// @Security ApiKeyAuth
func GetAllReports(c *gin.Context) {

	res := services.GetAllReports()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}

// @Summary Get Report by Report ID
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
		c.JSON(http.StatusBadRequest, "error to get report details")
		handler.WriteErrorLog(c, "error to get report details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// @Summary Create Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param Report body entities.Report true "report"
// @Success 200 {object} models.Response
// @Router /Report/Create [post]
// @Security ApiKeyAuth
func CreateReport(c *gin.Context) {

	// body := new(models.Instance)
	body := new(entities.Report)

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
// @Param Report body entities.Report true "report"
// @Success 200 {object} string
// @Router /Report/Update [put]
// @Security ApiKeyAuth
func UpdateReport(c *gin.Context) {

	body := new(entities.Report)

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

// func UpdateReport(c *gin.Context) {

// 	// 創建一個新的 Report 實例
// 	var body entities.Report

// 	// 明確綁定 JSON 請求體
// 	if err := c.ShouldBindJSON(&body); err != nil {
// 		// 如果綁定失敗，返回錯誤響應
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	// 調用服務層更新報告
// 	res := services.UpdateReport(body)

// 	// 如果更新失敗，返回錯誤響應
// 	if !res.Success {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": res.Msg})
// 		return
// 	}

// 	// 如果更新成功，返回成功響應
// 	c.JSON(http.StatusOK, gin.H{"message": "Update successful", "data": res.Body})
// }

// @Summary Delete Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} string
// @Router /Report/Delete/{id} [delete]
// @Security ApiKeyAuth
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
