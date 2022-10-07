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

// @Summary Get Dashboard Menu By User ID
// @Tags Menu
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.DashboardMenu
// @Router /api/v1/Menu/GetDashboardMenu [get]
// @Security ApiKeyAuth
func GetDashboardMenuByID(c *gin.Context) {

	// Check token
	token := c.Request.Header["Authorization"]
	if len(token) == 0 {
		c.JSON(http.StatusBadRequest, "auhorization token is required")
		handler.WriteErrorLog(c, "authorized token is required")
		return
	}

	// Get User ID
	user, err := services.GetUser(token[0])
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	userID := int(user["id"].(float64))

	// Find DB
	r, err := services.GetDashboardMenuByID(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get All Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.Menu
// @Router /api/v1/Menu/GetAll [get]
// @Security ApiKeyAuth
func GetAllMenu(c *gin.Context) {

	// Find DB
	r, err := services.GetAllMenu()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get Dashboard of Menu by Menu ID
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.Menu
// @Router /api/v1/Menu/GetDashboards/{id} [get]
// @Security ApiKeyAuth
func GetDashboardByMenuID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Find DB
	r, err := services.GetDashboardByMenuID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Create Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param menu body entities.Menu true "new menu"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/Create [post]
// @Security ApiKeyAuth
func CreateMenu(c *gin.Context) {

	body := new(entities.Menu)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// check menu name
	ch1 := services.CheckMenuName(body.MenuName)
	if !ch1.Success {
		c.JSON(http.StatusBadRequest, ch1.Msg)
		return
	}

	// create menu in db
	res := services.CreateMenu(*body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Add Dashboard to Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param dashboards body entities.Dashboard true "new dashboards"
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/AddDashboard/{id} [post]
// @Security ApiKeyAuth
func AddMenuDashboard(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(entities.Dashboard)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// check point (check menu id)
	res := services.CheckMenubyID(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}

	// add dashboard to menu in db
	res = services.AddMenuDashboard(id, *body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param menu body entities.Menu true "update menu info without dashboards"
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/Update/{id} [put]
// @Security ApiKeyAuth
func UpdateMenu(c *gin.Context) {

	menuID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(entities.Menu)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// check point (check menu id)
	check1 := services.CheckMenubyID(menuID)
	if !check1.Success {
		c.JSON(http.StatusBadRequest, check1.Msg)
		handler.WriteErrorLog(c, check1.Msg)
		return
	}

	// check point (check menu name)
	check2 := services.CheckMenuNameInOthers(menuID, body.MenuName)
	if !check2.Success {
		c.JSON(http.StatusBadRequest, check2.Msg)
		handler.WriteErrorLog(c, check2.Msg)
		return
	}

	// update menu info to db
	res := services.UpdateMenu(*body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Dashboard Under Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param dashboard body models.Dashboard true "update dashboard"
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/UpdateDashboard/{id} [put]
// @Security ApiKeyAuth
func UpdateDashboard(c *gin.Context) {

	menuID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(models.Dashboard)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// check point (check menu id)
	check1 := services.CheckMenubyID(menuID)
	if !check1.Success {
		c.JSON(http.StatusBadRequest, check1.Msg)
		handler.WriteErrorLog(c, check1.Msg)
		return
	}

	// check point (check alias name)
	check2 := services.CheckAliasInOthers(body.DashboardID, menuID, body.Alias)
	if !check2.Success {
		c.JSON(http.StatusBadRequest, check2.Msg)
		handler.WriteErrorLog(c, check2.Msg)
		return
	}

	// update menu info to db
	res := services.UpdateDashboard(*body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Delete Menu By Id
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteMenuByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// Check Point
	res := services.CheckMenubyID(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}
	// Delete DB
	res = services.DeleteMenuByID(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}

// @Summary Delete Dashboard from Menu
// @Tags Menu
// @Accept  json
// @Produce  json
// @Param id path int true "menu id"
// @Param dashboard_id query int true "dashboard_id"
// @Success 200 {object} models.Response
// @Router /api/v1/Menu/DeleteDashboard/{id} [delete]
// @Security ApiKeyAuth
func DeleteMenuDashboard(c *gin.Context) {
	menuID, err := strconv.Atoi(c.Param("id"))

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
	res := services.CheckMenubyID(menuID)
	if !res.Success {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}
	// Delete DB
	res = services.DeleteMenuDashboard(menuID, dashboardID)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res.Msg)
}
