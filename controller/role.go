package controller

import (
	"net/http"
	"strconv"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/models"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// @Summary Create Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param  role body entities.Role true "new role"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/Create [post]
// @Security ApiKeyAuth
func CreateRole(c *gin.Context) {

	body := new(entities.Role)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// // Check Admin
	// token := c.Request.Header["Authorization"]
	// if len(token) == 0 {
	// 	c.JSON(http.StatusBadRequest, "auhorization token is required")
	// 	handler.WriteErrorLog(c, "authorized token is required")
	// 	return
	// } else {
	// 	r, err := services.GetUser(token[0])
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, err.Error())
	// 		handler.WriteErrorLog(c, err.Error())
	// 		return
	// 	} else if r["role_id"] == nil {
	// 		c.JSON(http.StatusBadRequest, "cannot get role_id")
	// 		handler.WriteErrorLog(c, "cannot get role_id")
	// 		return
	// 	}
	// 	role_id := int(r["role_id"].(float64))

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }

	// Check Point
	check := services.CheckRoleByName(body.RoleName)
	if !check.Success {
		c.JSON(http.StatusBadRequest, check.Msg)
		handler.WriteErrorLog(c, check.Msg)
		return
	}

	// Create DB
	res := services.CreateRole(*body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Get All Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Role
// @Router /api/v1/Role/GetAll [get]
// @Security ApiKeyAuth
func GetAllRole(c *gin.Context) {

	// // Check Admin
	// token := c.Request.Header["Authorization"]
	// if len(token) == 0 {
	// 	c.JSON(http.StatusBadRequest, "auhorization token is required")
	// 	handler.WriteErrorLog(c, "authorized token is required")
	// 	return
	// } else {
	// 	r, err := services.GetUser(token[0])
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, err.Error())
	// 		handler.WriteErrorLog(c, err.Error())
	// 		return
	// 	}
	// 	role_id := int(r["role_id"].(float64))

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }

	r, err := services.GetRoleAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Update Role By ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param role body entities.Role true "update role info"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/Update/{id} [put]
// @Security ApiKeyAuth
func UpdateRole(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(entities.Role)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// // Check Admin
	// token := c.Request.Header["Authorization"]
	// if len(token) == 0 {
	// 	c.JSON(http.StatusBadRequest, "auhorization token is required")
	// 	handler.WriteErrorLog(c, "authorized token is required")
	// 	return
	// } else {
	// 	r, err := services.GetUser(token[0])
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, err.Error())
	// 		handler.WriteErrorLog(c, err.Error())
	// 		return
	// 	} else if r["role_id"] == nil {
	// 		c.JSON(http.StatusBadRequest, "cannot get role_id")
	// 		handler.WriteErrorLog(c, "cannot get role_id")
	// 		return
	// 	}

	// 	role_id := int(r["role_id"].(float64))

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }

	// Check Point ID
	check1 := services.CheckRoleById(id)
	if !check1.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Check Point Name
	check2 := services.CheckRoleByNameInOther(body.RoleID, body.RoleName)
	if !check2.Success {
		c.JSON(http.StatusBadRequest, "Name is already existed")
		handler.WriteErrorLog(c, "Name is already existed")
		return
	}

	// Update DB
	res := services.UpdateRole(*body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Update Dashboard By Role Id
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param update body models.UpdateDashboard true "old & new dashboard id"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/UpdateDashboard/{id} [put]
// @Security ApiKeyAuth
func UpdateRoleDashboard(c *gin.Context) {

	roleID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(models.UpdateDashboard)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Check Point ID
	check1 := services.CheckRoleById(roleID)
	if !check1.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Update DB
	res := services.UpdateRoleDashboard(roleID, body.OldID, body.NewID)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Add Dashboard to Role ID
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param dashboard_id query int true "dashboard_id"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/AddDashboard/{id} [post]
// @Security ApiKeyAuth
func AddRoleDashboard(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	dashboardID, err := strconv.Atoi(c.Query("dashboard_id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Check Point role ID
	check1 := services.CheckRoleById(id)
	if !check1.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}

	// Check point dashboard ID
	check2 := services.CheckDashboardID(dashboardID)
	if !check2.Success {
		c.JSON(http.StatusBadRequest, check2.Msg)
		handler.WriteErrorLog(c, check2.Msg)
		return
	}

	// // Check Admin
	// token := c.Request.Header["Authorization"]
	// if len(token) == 0 {
	// 	c.JSON(http.StatusBadRequest, "auhorization token is required")
	// 	handler.WriteErrorLog(c, "authorized token is required")
	// 	return
	// } else {
	// 	r, err := services.GetUser(token[0])
	// 	if err != nil {
	// 		c.JSON(http.StatusBadRequest, err.Error())
	// 		handler.WriteErrorLog(c, err.Error())
	// 		return
	// 	} else if r["role_id"] == nil {
	// 		c.JSON(http.StatusBadRequest, "cannot get role_id")
	// 		handler.WriteErrorLog(c, "cannot get role_id")
	// 		return
	// 	}

	// 	role_id := int(r["role_id"].(float64))

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }

	// Update DB
	res := services.AddRoleDashboard(id, dashboardID)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Delete Role By Id
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteRoleById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	if id == 1 {
		c.JSON(http.StatusBadRequest, "role admin is not allow to delete")
		handler.WriteErrorLog(c, "role admin is not allow to delete")
		return
	}

	// Check Point
	res := services.CheckRoleById(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}
	// Delete DB
	res = services.DeleteRoleById(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Delete Dashboard from Role
// @Tags Role
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Param dashboard_id query int true "dashboard_id"
// @Success 200 {object} models.Response
// @Router /api/v1/Role/DeleteDashboard/{id} [delete]
// @Security ApiKeyAuth
func DeleteRoleDashboard(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	dashboardID, err := strconv.Atoi(c.Query("dashboard_id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Check Point
	res := services.CheckRoleById(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}

	// Delete DB
	res = services.DeleteRoleDashboard(id, dashboardID)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}
