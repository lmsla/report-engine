package controller

import (
	"fmt"
	"net/http"
	"net/url"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/log"
	"report-backend-golang/services"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Html
// @Tags Html
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /Html/Create/{id} [post]
// @Security ApiKeyAuth
func CreateHtml(c *gin.Context) {

	ReportID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "report ID should be int")
		handler.WriteErrorLog(c, "report ID should be integer")
		return
	}
	nowtime := time.Now().Unix()
	services.CreateHtml(nowtime, ReportID)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, err.Error())
	// 	handler.WriteErrorLog(c, err.Error())
	// 	return
	// }
	// c.JSON(http.StatusOK, r)
}

// @Summary PDF
// @Tags PDF
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /PDF/Create/{id} [post]
// @Security ApiKeyAuth
func CreatePDF(c *gin.Context) {

	ScheduleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "Schedule ID should be int")
		handler.WriteErrorLog(c, "report ID should be integer")
		return
	}
	c.JSON(http.StatusOK, "報表試寄中，請稍候")

	log.Logrecord("排程", fmt.Sprintf("報表試寄, Schedule ID : %d", ScheduleID))
	test := true
	go services.FuncAddToCron(test, ScheduleID)

}

// @Summary Screenshot by Url
// @Tags Screenshot
// @Accept  json
// @Produce  json
// @Param user query string true "user"
// @Param password query string true "password"
// @Param url query string true "url"
// @Param name query string true "name"
// @Success 200
// @Security ApiKeyAuth
// @Router /ScreenShotbyUrl [post]
func ScreenShotByUrl(c *gin.Context) {

	body := new(entities.ScreenshotUrl)
	c.Bind(&body)

	err := services.ScreenShotByUrl(body.Url, body.User, body.Password, body.Name)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to Screenshot By Url")
		handler.WriteErrorLog(c, "error to Screenshot By Url")
		return
	}
	c.JSON(http.StatusOK, "Screenshot success")
}

// @Summary Encode Url
// @Tags Screenshot
// @Accept  json
// @Produce  json
// @Param user query string true "user"
// @Param password query string true "password"
// @Param url query string true "url"
// @Param name query string true "name"
// @Success 200
// @Security ApiKeyAuth
// @Router /EncodeUrl [post]
func EncodeUrl(c *gin.Context) {

	body := new(entities.ScreenshotUrl)
	c.Bind(&body)

	// 自動編碼 URL
	encodedUrl := url.QueryEscape(body.Url)

	// // 調用服務邏輯，傳遞編碼後的 URL
	// err := services.ScreenShotByUrl(encodedUrl, body.User, body.Password, body.Name)
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "error to Screenshot By Url")
	// 	handler.WriteErrorLog(c, "error to Screenshot By Url")
	// 	return
	// }
	c.JSON(http.StatusOK, gin.H{
		"encoded_url": encodedUrl,
	})
}
