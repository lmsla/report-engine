package controller

import (
	"net/http"
	"report-backend-golang/handler"

	"github.com/gin-gonic/gin"
)

// @Summary Write Frontend Log
// @Tags Frontend
// @Accept  json
// @Produce  json
// @Param type query string true "type"
// @Param usr query string true "usr"
// @Param msg query string true "msg"
// @Success 200 {string} string
// @Router /write-frontend-log [get]
func WriteFrontendLog(c *gin.Context) {
	logType := c.Query("type")
	usr := c.Query("usr")
	msg := c.Query("msg")
	err := handler.WriteFrontendLog(logType, usr, msg)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	c.JSON(http.StatusOK, "ok")
}