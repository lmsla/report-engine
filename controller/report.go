package controller

import (
	// "fmt"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	// "report-backend-golang/screenshot"
	"report-backend-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param instance body entities.Report true "new report"
// @Success 200 {object} models.Response
// @Router /api/v1/Report/Create [post]
// @Security ApiKeyAuth
func CreateReport(c *gin.Context) {

	body := new(entities.Report)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateReport(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get All Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Report
// @Router /api/v1/Report/GetAll [get]
// @Security ApiKeyAuth
func GetAllReport(c *gin.Context) {

	// Find DB
	r, err := services.GetAllReport()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}


// @Summary Update Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param user body entities.Report true "report"
// @Success 200 {object} models.Response
// @Router /api/v1/Report/ReportUpdate [put]
// @Security ApiKeyAuth
func UpdateReport(c *gin.Context) {

	body := new(entities.Report)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
	}
	// id := body.ReportID
	// name := body.Name
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be int")
		handler.WriteErrorLog(c, "id should be int")
		return
	}

	// // Check Point 1
	// chk_id, _ := services.GetMemberByID(id)
	// if len(chk_id) == 0 {
	// 	c.JSON(http.StatusBadRequest, "ID is not existed")
	// 	handler.WriteErrorLog(c, "ID is not existed")
	// 	return
	// }

	// // Check Point 2
	// chk_name, _ := services.GetMemberByName(name)
	// if len(chk_name) != 0 && chk_name[0].MemberName != chk_id[0].MemberName {
	// 	c.JSON(http.StatusBadRequest, "Name is already existed")
	// 	handler.WriteErrorLog(c, "Name is already existed")
	// 	return
	// }

	// Create DB
	r1 := services.UpdateReport(*body)
	if !r1.Success {
		c.JSON(http.StatusBadRequest, r1)
		handler.WriteErrorLog(c, r1.Msg)
		return
	}
	c.JSON(http.StatusOK, r1)
}


// @Summary Get report by Report ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Report
// @Router /api/v1/Report/GetReport/{id} [get]
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


// @Summary Get Instance of report by Instance ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Instance
// @Router /api/v1/Report/GetInstance/{id} [get]
// @Security ApiKeyAuth
func GetInstanceByInstanceID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetInstanceTypeAndIP(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}


// @Summary Get dashboard of report by report ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.RDashboard
// @Router /api/v1/Report/GetDashboard/{id} [get]
// @Security ApiKeyAuth
func GetDashboardOfReportbyReportID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetDashboardInReport(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}


// @Summary Get Instance of report by report ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.RInstance
// @Router /api/v1/Report/GetInstanceOfReport/{id} [get]
// @Security ApiKeyAuth
func GetInstanceOfReportbyReportID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetInstanceInReportbyReportID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}


// @Summary Get report of instance by instance ID
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.RDashboard
// @Router /api/v1/Report/GetDashboardOfInstance/{id} [get]
// @Security ApiKeyAuth
func GetDashboardOfInstanceByInstanceID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetDashboardOfInstanceByInstanceID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}

// @Summary Delete Report
// @Tags Report
// @Accept  json
// @Produce  json
// @Param id path int true "report id"
// @Success 200 {object} models.Response
// @Router /api/v1/Report/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteReport(c *gin.Context) {

	reportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "report id should be int")
		handler.WriteErrorLog(c, "report id should be integer")
		return
	}

	// Update DB
	r := services.DeleteReport(reportID)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}





// // @Summary Get ReportDashboard of server by Report ID
// // @Tags Report
// // @Accept  json
// // @Produce  json
// // @Param id path int true "report id"
// // @Success 200 {object} entities.Report
// // @Router /api/v1/Report/GetDashboards/{id} [get]
// // @Security ApiKeyAuth
// func GetReportDashboardByInstanceID(c *gin.Context) {

// 	ReportID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "instance ID should be int")
// 		handler.WriteErrorLog(c, "instance ID should be integer")
// 		return
// 	}

// 	inventory, err := services.GetInstanceByID(ReportID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "error to get inventory details")
// 		handler.WriteErrorLog(c, "error to get inventory details")
// 		return
// 	}

// 	switch inventory.Type {
// 	case "grafana":
// 		r, err := services.GetAllGrafanaDashboardTitle(inventory)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, err.Error())
// 			handler.WriteErrorLog(c, err.Error())
// 			return
// 		}
// 		c.JSON(http.StatusOK, r)
// 	case "kibana":
// 		r, err := services.GetALLKibanaDashboardTitle(inventory)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, err.Error())
// 			// handler.WriteErrorLog(c, err.Error())
// 			return
// 		}
// 		c.JSON(http.StatusOK, r)
// 	default:
// 		c.JSON(http.StatusBadRequest, "instance type unknown")
// 	}
// }














// // @Summary Get Dashboard of server by Report ID
// // @Tags Report
// // @Accept  json
// // @Produce  json
// // @Param id path int true "report id"
// // @Success 200 {object} entities.Report
// // @Router /api/v1/Report/GetDashboards/{id} [get]
// // @Security ApiKeyAuth
// func GetDBDashboardByReportID(c *gin.Context) {

// 	reportID, err := strconv.Atoi(c.Param("id"))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "report ID should be int")
// 		handler.WriteErrorLog(c, "report ID should be integer")
// 		return
// 	}

// 	inventory, err := services.GetInstanceByReportID(reportID)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, "error to get inventory details")
// 		handler.WriteErrorLog(c, "error to get inventory details")
// 		return
// 	}

// 	switch inventory.Type {
// 	case "grafana":
// 		r, err := services.GetAllGrafanaDashboardTitle(inventory)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, err.Error())
// 			handler.WriteErrorLog(c, err.Error())
// 			return
// 		}
// 		c.JSON(http.StatusOK, r)
// 	case "kibana":
// 		r, err := services.GetALLKibanaDashboardTitle(inventory)
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, err.Error())
// 			// handler.WriteErrorLog(c, err.Error())
// 			return
// 		}
// 		c.JSON(http.StatusOK, r)
// 	default:
// 		c.JSON(http.StatusBadRequest, "instance type unknown")
// 	}
// }