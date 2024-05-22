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

	apiv1.GET("/testdefer", TestDefer)

	apiv1.GET("/Instance/GetAll", GetAllInstances)
	apiv1.POST("/Instance/Create", CreateInstance)
	apiv1.PUT("/Instance/Update", UpdateInstance)
	apiv1.DELETE("/Instance/Delete/:id", DeleteInstance)
	apiv1.GET("/Instance/GetDashboards/:id", GetDBDashboardByInstanceID)
	apiv1.GET("/Instance/GetVisualizations/:id", GetDBVisualizationByInstanceID)

	apiv1.GET("/Dropdown", GetDropdownSource)

	apiv1.GET("/Report/GetAll", GetAllReports)
	apiv1.POST("/Report/Create", CreateReport)
	apiv1.PUT("/Report/Update", UpdateReport)
	apiv1.DELETE("/Report/Delete/:id", DeleteReport)
	apiv1.GET("/Report/GetReportByScheduleID/:id", GetReportByScheduleID)
	apiv1.GET("/Report/GetReport/:id", GetReportByReportID)

	apiv1.GET("/Element/GetAll", GetAllElements)
	apiv1.POST("/Element/Create", CreateElement)
	apiv1.PUT("/Element/Update", UpdateElement)
	apiv1.DELETE("/Element/Delete/:id", DeleteElement)
	apiv1.GET("/Element/GetElementByReportID/:id", GetElementByReportID)

	apiv1.GET("/Schedule/GetAll", GetAllSchedule)
	apiv1.GET("/Schedule/GetSchedule/:id", GetScheduleByScheduleID)
	apiv1.POST("/Schedule/Create", CreateSchedule)
	apiv1.DELETE("/Schedule/Delete/:id", DeleteSchedule)
	apiv1.PUT("/Schedule/Update", UpdateSchedule)

	// apiv1.POST("Screenshot/Create/:id", GetScreenShot)

	apiv1.POST("/Html/Create/:id", CreateHtml)
	apiv1.POST("/PDF/Create/:id", CreatePDF)

	apiv1.POST("/Mail/Send/:id", SendEmailBySchedule)

	apiv1.GET("/History/GetAll", GetAllHitory)
	apiv1.GET("/History/GetOldHistory", GetOldHitory)
	apiv1.GET("/History/GetHistory/:id", GetHistoryByHistoryID)
	// apiv1.POST("/History/HistoryReport/:id", CreateHistoryReport)
	apiv1.POST("/History/HistoryReport", CreateHistoryReport)


	apiv1.GET("/get-sso-url", GetSsoURL)

	return router
}
