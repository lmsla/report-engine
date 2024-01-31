package controller

import (
	"fmt"
	"net/http"
	// "report-backend-golang/models"
	"report-backend-golang/handler"
	"report-backend-golang/log"
	"report-backend-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
	// "report-backend-golang/entities"
)

// @Summary Get history
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /History/GetAll [get]
func GetAllHitory(c *gin.Context) {

	res := services.GetAllHistory()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}




// @Summary Get History by History ID
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.History
// @Router /History/GetHistory/{id} [get]
// @Security ApiKeyAuth
func GetHistoryByHistoryID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetHistoryByHistoryID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}


// @Summary Get Old history
// @Tags History
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Response
// @Router /History/GetOldHistory [get]
func GetOldHitory(c *gin.Context) {

	res := services.GetOldHistory()

	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		return
	}

	c.JSON(http.StatusOK, res.Body)
}



// @Summary Create History Report by History ID
// @Tags History
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.History
// @Router /History/HistoryReport/{id} [post]
// @Security ApiKeyAuth
func CreateHistoryReport(c *gin.Context) {

	HistoryID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "HistoryID ID should be int")
		// handler.WriteErrorLog(c, "report ID should be integer")
		return
	}
	log.Logrecord("排程",fmt.Sprintf("重新寄送歷史報表, History ID : %d",HistoryID))

	services.CreateHistoryReport(HistoryID)
	
	log.Logrecord("排程",fmt.Sprintf("歷史報表寄送完成, History ID : %d",HistoryID))

}

