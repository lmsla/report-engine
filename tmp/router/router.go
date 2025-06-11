package router

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"report-backend-golang/controller"
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

	//*** 不需要驗證 Token ***//
	apiv1 := router.Group("api/v1")
	{
		apiv1.GET("/get-sso-url", controller.GetSsoURL)
		apiv1.GET("/user/get-server-menu", controller.GetServerMenu)
		apiv1.GET("/get-server-module", controller.GetServerModule)
		apiv1.POST("/ScreenShotbyUrl",controller.ScreenShotByUrl)
		apiv1.POST("/EncodeUrl",controller.EncodeUrl)
		// apiv1.GET("/Instance/GetAll", controller.GetAllInstances)
	}

	//*** 驗證 Token 跟 SSO 取 UserInfo AccessHosts ***//
	apiv1_auth := router.Group("api/v1")
	apiv1_auth.Use(controller.GetUserInfo)
	{
		apiv1_auth.GET("/Instance/GetAll", controller.GetAllInstances)
		apiv1_auth.GET("/Instance/GetInstance/:id",controller.GetInstanceByInstanceID)
		apiv1_auth.POST("/Instance/Create", controller.CreateInstance)
		apiv1_auth.PUT("/Instance/Update", controller.UpdateInstance)
		apiv1_auth.DELETE("/Instance/Delete/:id", controller.DeleteInstance)
		apiv1_auth.GET("/Instance/GetDashboards/:id", controller.GetDBDashboardByInstanceID)
		apiv1_auth.GET("/Instance/GetVisualizations/:id", controller.GetDBVisualizationByInstanceID)
		apiv1_auth.GET("/Instance/CheckInstance/:id",controller.CheckInstanceByInstanceID)

		apiv1_auth.GET("/Dropdown", controller.GetDropdownSource)
		apiv1_auth.GET("/DropdownFields", controller.DropdownFields)

		apiv1_auth.GET("/Report/GetAll", controller.GetAllReports)
		apiv1_auth.POST("/Report/Create", controller.CreateReport)
		apiv1_auth.PUT("/Report/Update", controller.UpdateReport)
		apiv1_auth.DELETE("/Report/Delete/:id", controller.DeleteReport)
		apiv1_auth.GET("/Report/GetReportByScheduleID/:id", controller.GetReportByScheduleID)
		apiv1_auth.GET("/Report/GetReport/:id", controller.GetReportByReportID)

		apiv1_auth.GET("/Element/GetAll", controller.GetAllElements)
		apiv1_auth.POST("/Element/Create", controller.CreateElement)
		apiv1_auth.PUT("/Element/Update", controller.UpdateElement)
		apiv1_auth.DELETE("/Element/Delete/:id", controller.DeleteElement)
		apiv1_auth.GET("/Element/GetElementByReportID/:id", controller.GetElementByReportID)

		apiv1_auth.GET("/Table/GetAll", controller.GetTables)
		apiv1_auth.POST("/Table/Create", controller.CreateTable)
		apiv1_auth.PUT("/Table/Update", controller.UpdateTable)
		apiv1_auth.DELETE("/Table/Delete/:id", controller.DeleteTable)
		apiv1_auth.GET("/Table/GetTable/:id", controller.GetTableByID)
		apiv1_auth.GET("/Table/GetTableByReportID/:id", controller.GetTableByReportID)

		apiv1_auth.GET("/Schedule/GetAll", controller.GetAllSchedule)
		apiv1_auth.GET("/Schedule/GetSchedule/:id", controller.GetScheduleByScheduleID)
		apiv1_auth.POST("/Schedule/Create", controller.CreateSchedule)
		apiv1_auth.DELETE("/Schedule/Delete/:id", controller.DeleteSchedule)
		apiv1_auth.PUT("/Schedule/Update", controller.UpdateSchedule)

		apiv1_auth.POST("/Html/Create/:id", controller.CreateHtml)
		// 報表試寄 PDF
		apiv1_auth.POST("/PDF/Create/:id", controller.CreatePDF)

		apiv1_auth.POST("/Mail/Send/:id", controller.SendEmailBySchedule)

		apiv1_auth.GET("/History/GetAll", controller.GetAllHitory)
		apiv1_auth.GET("/History/GetOldHistory", controller.GetOldHitory)
		apiv1_auth.GET("/History/GetHistory/:id", controller.GetHistoryByHistoryID)
		apiv1_auth.POST("/History/HistoryReport/:id", controller.CreateHistoryReport)
	}
	return router
}
