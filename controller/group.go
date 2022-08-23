package controller

import (
	// "fmt"
	"net/http"
	"strconv"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

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