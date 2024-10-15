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

	//*** 不需要驗證 Token ***//
	apiv1 := router.Group("api/v1")
	{
		apiv1.GET("/get-sso-url", GetSsoURL)
	}

	//*** 驗證 Token 跟 SSO 取 UserInfo AccessHosts ***//
	apiv1_auth := router.Group("api/v1")
	apiv1_auth.Use(GetUserInfo)
	{
		apiv1_auth.GET("/Instance/GetAll", GetAllInstances)
		apiv1_auth.POST("/Instance/Create", CreateInstance)
		apiv1_auth.PUT("/Instance/Update", UpdateInstance)
		apiv1_auth.DELETE("/Instance/Delete/:id", DeleteInstance)
		apiv1_auth.GET("/Instance/GetDashboards/:id", GetDBDashboardByInstanceID)
		apiv1_auth.GET("/Instance/GetVisualizations/:id", GetDBVisualizationByInstanceID)

		apiv1_auth.GET("/Dropdown", GetDropdownSource)
		apiv1_auth.GET("/DropdownFields", DropdownFields)

		apiv1_auth.GET("/Report/GetAll", GetAllReports)
		apiv1_auth.POST("/Report/Create", CreateReport)
		apiv1_auth.PUT("/Report/Update", UpdateReport)
		apiv1_auth.DELETE("/Report/Delete/:id", DeleteReport)
		apiv1_auth.GET("/Report/GetReportByScheduleID/:id", GetReportByScheduleID)
		apiv1_auth.GET("/Report/GetReport/:id", GetReportByReportID)

		apiv1_auth.GET("/Element/GetAll", GetAllElements)
		apiv1_auth.POST("/Element/Create", CreateElement)
		apiv1_auth.PUT("/Element/Update", UpdateElement)
		apiv1_auth.DELETE("/Element/Delete/:id", DeleteElement)
		apiv1_auth.GET("/Element/GetElementByReportID/:id", GetElementByReportID)

		apiv1_auth.GET("/Table/GetAll", GetTables)
		apiv1_auth.POST("/Table/Create", CreateTable)
		apiv1_auth.PUT("/Table/Update", UpdateTable)
		apiv1_auth.DELETE("/Table/Delete/:id", DeleteTable)
		apiv1_auth.GET("/Table/GetTable/:id", GetTableByID)
		apiv1_auth.GET("/Table/GetTableByReportID/:id", GetTableByReportID)

		apiv1_auth.GET("/Schedule/GetAll", GetAllSchedule)
		apiv1_auth.GET("/Schedule/GetSchedule/:id", GetScheduleByScheduleID)
		apiv1_auth.POST("/Schedule/Create", CreateSchedule)
		apiv1_auth.DELETE("/Schedule/Delete/:id", DeleteSchedule)
		apiv1_auth.PUT("/Schedule/Update", UpdateSchedule)

		apiv1_auth.POST("/Html/Create/:id", CreateHtml)
		apiv1_auth.POST("/PDF/Create/:id", CreatePDF)

		apiv1_auth.POST("/Mail/Send/:id", SendEmailBySchedule)

		apiv1_auth.GET("/History/GetAll", GetAllHitory)
		apiv1_auth.GET("/History/GetOldHistory", GetOldHitory)
		apiv1_auth.GET("/History/GetHistory/:id", GetHistoryByHistoryID)
		apiv1_auth.POST("/History/HistoryReport/:id", CreateHistoryReport)


	}

	return router
}
