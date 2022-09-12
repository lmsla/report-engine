package router

import (
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

	router.GET("/api/v1/Menu/GetDashboardMenu", GetDashboardMenuByID)

	apiv1 := router.Group("api/v1")
	// apiv1.Use(CheckAdmin)
	{
		apiv1.POST("Instance/Create", CreateInstance)
		apiv1.POST("Instance/Verify", VerifyInstance)
		apiv1.GET("Instance/GetAll", GetAllInstance)
		apiv1.GET("Instance/GetDashboards/:id", GetDBDashboardByInstanceID)
		apiv1.PUT("Instance/Update/:id", UpdateInstance)
		apiv1.DELETE("Instance/Delete/:id", DeleteInstance)
	}

	{
		apiv1.GET("User/GetAllFromSSO", GetUserAllFromSSO)
		apiv1.POST("User/Create", CreateUser)
		apiv1.GET("User/GetAll", GetAllUser)
		apiv1.PUT("User/Update", UpdateUser)
		apiv1.DELETE("User/Delete/:id", DeleteUserByID)
	}

	{
		apiv1.POST("Role/Create", CreateRole)
		apiv1.POST("Role/AddDashboard/:id", AddRoleDashboard)
		apiv1.GET("Role/GetAll", GetAllRole)
		apiv1.PUT("Role/Update/:id", UpdateRole)
		apiv1.PUT("Role/UpdateDashboard/:id", UpdateRoleDashboard)
		apiv1.DELETE("Role/Delete/:id", DeleteRoleById)
		apiv1.DELETE("Role/DeleteDashboard/:id", DeleteRoleDashboard)
	}
	{

		apiv1.POST("Menu/Create", CreateMenu)
		apiv1.POST("Menu/AddDashboard/:id", AddMenuDashboard)
		apiv1.GET("Menu/GetAll", GetAllMenu)
		apiv1.GET("Menu/GetDashboards/:id", GetDashboardByMenuID)
		apiv1.PUT("Menu/Update/:id", UpdateMenu)
		apiv1.PUT("Menu/UpdateDashboard/:id", UpdateDashboard)
		apiv1.DELETE("Menu/Delete/:id", DeleteMenuByID)
		apiv1.DELETE("Menu/DeleteDashboard/:id", DeleteMenuDashboard)
	}

	{
		apiv1.POST("Group/Create", CreateGroup)
		apiv1.GET("Group/GetAll", GetAllGroup)
		apiv1.POST("Group/CreateMember", CreateMember)
		apiv1.DELETE("Group/DeleteMember/:id", DeleteMember)
		apiv1.DELETE("Group/DeleteMemberbyName/:name", DeleteMemberbyName)
		apiv1.GET("Group/GetMember/:id",GetMemberByGroupID)
	}

	{
		apiv1.POST("Report/Create", CreateReport)
		apiv1.GET("Report/GetAll", GetAllReport)
		apiv1.GET("Report/GetInstance/:id",GetInstanceByInstanceID)
		apiv1.GET("Report/GetReport/:id",GetReportByReportID)
		apiv1.GET("Report/GetDashboard/:id",GetDashboardOfReportbyReportID)
		apiv1.GET("Report/GetInstanceOfReport/:id",GetInstanceOfReportbyReportID)
		apiv1.GET("Report/GetDashboardOfInstance/:id",GetDashboardOfInstanceByInstanceID)

	}


	{
		apiv1.POST("Schedule/Create", CreateSchedule)
		apiv1.GET("Schedule/GetAll", GetAllSchedule)
		apiv1.GET("Schedule/GetSchedule/:id",GetScheduleByScheduleID)
		apiv1.POST("Schedule/ReportCreate/:id",CreateReportbySchedule)
	}

	{
		apiv1.POST("Screenshot/Create/:id", GetScreenShot)
		apiv1.POST("Html/Create/:id", CreateHtml)

	}

	return router
}
