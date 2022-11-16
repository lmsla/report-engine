package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	. "report-backend-golang/controller"
	"report-backend-golang/utils"

	_ "report-backend-golang/docs"
	"report-backend-golang/global"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func LoadRouter() *gin.Engine {

	gin.SetMode(global.EnvConfig.Server.Mode)
	router := gin.Default()

	router.Use(utils.CorsConfig())

	// swagger
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	router.GET("healthcheck", func(c *gin.Context) {
		c.JSON(http.StatusOK, "")
	})
	apiv1 := router.Group("api/v1")

	apiv1.GET("/Instance/GetAll", GetAllInstances)
	apiv1.POST("/Instance/Create", CreateInstance)
	apiv1.PUT("/Instance/Update", UpdateInstance)
	apiv1.DELETE("/Instance/Delete/:id", DeleteInstance)

	return router
}
