package controller

import (
	"net/http"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

func CheckAdmin(c *gin.Context) {

	token := c.Request.Header["Authorization"]
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, "auhorization token is required")
		handler.WriteErrorLog(c, "authorized token is required")
		c.Abort()
		return
	} else {
		r, err := services.GetUser(token[0])
		if err != nil {
			c.JSON(http.StatusBadRequest, err.Error())
			handler.WriteErrorLog(c, err.Error())
			c.Abort()
			return
		} else if r["role_id"] == nil {
			c.JSON(http.StatusBadRequest, "cannot get role_id")
			handler.WriteErrorLog(c, "cannot get role_id")
			c.Abort()
			return
		}
		// role_id := int(r["role_id"].(float64))

		// if role_id != 1 {
		// 	c.JSON(http.StatusBadRequest, "unauthorize action")
		// 	handler.WriteErrorLog(c, "unauthorize action")
		// 	c.Abort()
		// 	return
		// }
	}
	c.Next()
}
