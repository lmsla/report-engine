package controller

import (
	// "fmt"
	"net/http"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"
	"strconv"

	"github.com/gin-gonic/gin"
)

// @Summary Create Group
// @Tags Group
// @Accept  json
// @Produce  json
// @Param instance body entities.Group true "new group"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/Create [post]
// @Security ApiKeyAuth
func CreateGroup(c *gin.Context) {

	body := new(entities.Group)
	// name := body.Name
	// id := body.GroupID
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateGroup(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get All Group
// @Tags Group
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.GroupMember
// @Router /api/v1/Group/GetAll [get]
// @Security ApiKeyAuth
func GetAllGroup(c *gin.Context) {

	// Find DB
	r, err := services.GetAllGroup()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Create Member
// @Tags Group
// @Accept  json
// @Produce  json
// @Param instance body entities.GroupMember true "new member"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/CreateMember [post]
// @Security ApiKeyAuth
func CreateMember(c *gin.Context) {

	body := new(entities.GroupMember)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	// Create DB
	r := services.CreateMember(*body)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Add Group Member to Group
// @Tags Group
// @Accept  json
// @Produce  json
// @Param groupmember body entities.GroupMember true "new member"
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/AddGroupMember/{id} [post]
// @Security ApiKeyAuth
func AddMemberToGroup(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	body := new(entities.GroupMember)
	err = c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// check point (check group id)
	res := services.CheckGroupID(id)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}

	// add dashboard to menu in db
	res = services.AddMemberToGroup(id, *body)
	if !res.Success {
		c.JSON(http.StatusBadRequest, res.Msg)
		handler.WriteErrorLog(c, res.Msg)
		return
	}
	c.JSON(http.StatusOK, res)
}

// @Summary Update Group
// @Tags Group
// @Accept  json
// @Produce  json
// @Param user body entities.Group true "group"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/Update [put]
// @Security ApiKeyAuth
func UpdateGroup(c *gin.Context) {

	body := new(entities.Group)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
	}
	id := body.GroupID
	name := body.Name
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be int")
		handler.WriteErrorLog(c, "id should be int")
		return
	}

	// Check Point 1
	chk_id, _ := services.GetGroupByID(id)
	if len(chk_id) == 0 {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}

	// Check Point 2
	chk_name, _ := services.GetGroupByName(name)
	if len(chk_name) != 0 && chk_name[0].Name != chk_id[0].Name {
		c.JSON(http.StatusBadRequest, "Name is already existed")
		handler.WriteErrorLog(c, "Name is already existed")
		return
	}

	// Create DB
	r1 := services.UpdateGroup(*body)
	if !r1.Success {
		c.JSON(http.StatusBadRequest, r1)
		handler.WriteErrorLog(c, r1.Msg)
		return
	}
	c.JSON(http.StatusOK, r1)
}

// @Summary Update Member
// @Tags Group
// @Accept  json
// @Produce  json
// @Param user body entities.GroupMember true "group"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/MemberUpdate [put]
// @Security ApiKeyAuth
func UpdateMember(c *gin.Context) {

	body := new(entities.GroupMember)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
	}
	id := body.GroupID
	name := body.MemberName
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be int")
		handler.WriteErrorLog(c, "id should be int")
		return
	}

	// Check Point 1
	chk_id, _ := services.GetMemberByID(id)
	if len(chk_id) == 0 {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}

	// Check Point 2
	chk_name, _ := services.GetMemberByName(name)
	if len(chk_name) != 0 && chk_name[0].MemberName != chk_id[0].MemberName {
		c.JSON(http.StatusBadRequest, "Name is already existed")
		handler.WriteErrorLog(c, "Name is already existed")
		return
	}

	// Create DB
	r1 := services.UpdateMember(*body)
	if !r1.Success {
		c.JSON(http.StatusBadRequest, r1)
		handler.WriteErrorLog(c, r1.Msg)
		return
	}
	c.JSON(http.StatusOK, r1)
}

// @Summary Delete Group
// @Tags Group
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/DeleteGroup/{id} [delete]
// @Security ApiKeyAuth
func DeleteGroup(c *gin.Context) {

	group_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "group id should be int")
		handler.WriteErrorLog(c, "group id should be integer")
		return
	}

	// Update DB
	r := services.DeleteGroup(group_id)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Delete Member
// @Tags Group
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/DeleteMember/{id} [delete]
// @Security ApiKeyAuth
func DeleteMember(c *gin.Context) {

	member_id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "member id should be int")
		handler.WriteErrorLog(c, "member id should be integer")
		return
	}

	// Update DB
	r := services.DeleteMember(member_id)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Delete MemberbyName
// @Tags Group
// @Accept  json
// @Produce  json
// @Param name path string true "name"
// @Success 200 {object} models.Response
// @Router /api/v1/Group/DeleteMemberbyName/{name} [delete]
// @Security ApiKeyAuth
func DeleteMemberbyName(c *gin.Context) {

	member_name := c.Param("name")
	// if err != nil {
	// 	c.JSON(http.StatusBadRequest, "member id should be int")
	// 	handler.WriteErrorLog(c, "member id should be integer")
	// 	return
	// }

	// Update DB
	r := services.DeleteMemberbyName(member_name)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r.Msg)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Get member by Group ID
// @Tags Group
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} entities.GroupMember
// @Router /api/v1/Group/GetMember/{id} [get]
// @Security ApiKeyAuth
func GetMemberByGroupID(c *gin.Context) {

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadGateway, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}

	// inventory, err := screenshot.GetInstanceByID(id)
	inventory, err := services.GetMemberByGroupID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, "error to get inventory details")
		handler.WriteErrorLog(c, "error to get inventory details")
		return
	}
	c.JSON(http.StatusOK, inventory)
}
