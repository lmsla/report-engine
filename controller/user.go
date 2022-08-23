package controller

import (
	"net/http"
	"strconv"
	"report-backend-golang/entities"
	"report-backend-golang/handler"
	"report-backend-golang/services"

	"github.com/gin-gonic/gin"
)

// @Summary Get All User From SSO
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.User
// @Router /api/v1/User/GetAllFromSSO [get]
// @Security ApiKeyAuth
func GetUserAllFromSSO(c *gin.Context) {

	token := c.Request.Header.Get("Authorization")

	r, err := services.GetUserAllFromSSO(token)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Create User
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body entities.User true "user"
// @Success 200 {object} entities.User
// @Router /api/v1/User/Create [post]
// @Security ApiKeyAuth
func CreateUser(c *gin.Context) {

	body := new(entities.User)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
	}

	name := body.UserName

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
	// 	role_id := r["role_id"].(float64)

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }
	// Check Point
	chk_name, _ := services.GetUserByName(name)
	if len(chk_name) != 0 {
		c.JSON(http.StatusBadRequest, "Name is already existed")
		handler.WriteErrorLog(c, "Name is already existed")
		return
	}

	// Create DB
	r1 := services.CreateUser(*body)
	if !r1.Success {
		c.JSON(http.StatusBadRequest, r1)
		handler.WriteErrorLog(c, r1.Msg)
		return
	}
	c.JSON(http.StatusOK, r1)
}

// @Summary Get All User
// @Tags User
// @Accept  json
// @Produce  json
// @Success 200 {object} entities.User
// @Router /api/v1/User/GetAll [get]
// @Security ApiKeyAuth
func GetAllUser(c *gin.Context) {

	r, err := services.GetUserAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
		return
	}
	c.JSON(http.StatusOK, r)
}

// @Summary Update User By ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param user body entities.User true "user"
// @Success 200 {object} models.Response
// @Router /api/v1/User/Update [put]
// @Security ApiKeyAuth
func UpdateUser(c *gin.Context) {

	body := new(entities.User)
	err := c.Bind(&body)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		handler.WriteErrorLog(c, err.Error())
	}

	id := body.UserID
	name := body.UserName

	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be int")
		handler.WriteErrorLog(c, "id should be int")
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
	// 	}
	// 	role_id := r["role_id"].(float64)

	// 	if role_id != 1 {
	// 		c.JSON(http.StatusBadRequest, "unauthorize action")
	// 		handler.WriteErrorLog(c, "unauthorize action")
	// 		return
	// 	}
	// }

	// Check Point 1
	chk_id, _ := services.GetUserByID(id)
	if len(chk_id) == 0 {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}

	// Check Point 2
	chk_name, _ := services.GetUserByName(name)
	if len(chk_name) != 0 && chk_name[0].UserName != chk_id[0].UserName {
		c.JSON(http.StatusBadRequest, "Name is already existed")
		handler.WriteErrorLog(c, "Name is already existed")
		return
	}

	// Create DB
	r1 := services.UpdateUser(*body)
	if !r1.Success {
		c.JSON(http.StatusBadRequest, r1)
		handler.WriteErrorLog(c, r1.Msg)
		return
	}
	c.JSON(http.StatusOK, r1)
}

// @Summary Delete User By ID
// @Tags User
// @Accept  json
// @Produce  json
// @Param id path int true "id"
// @Success 200 {object} models.Response
// @Router /api/v1/User/Delete/{id} [delete]
// @Security ApiKeyAuth
func DeleteUserByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, "id should be int")
		handler.WriteErrorLog(c, "id should be int")
		return
	}

	// Check Point
	chk_id, _ := services.GetUserByID(id)
	if len(chk_id) == 0 {
		c.JSON(http.StatusBadRequest, "ID is not existed")
		handler.WriteErrorLog(c, "ID is not existed")
		return
	}
	// Delete DB
	r := services.DeleteUserByID(id)
	if !r.Success {
		c.JSON(http.StatusBadRequest, r)
		handler.WriteErrorLog(c, r.Msg)
		return
	}
	c.JSON(http.StatusOK, r)
}
