package controller

import (
	"net/http"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/usecase"
	"github.com/NursiNursi/laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	router *gin.Engine
	userUC usecase.UserUseCase
}

func (u *UserController) createHandler(c *gin.Context) {
	var user model.UserCredential
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	user.Id = common.GenerateID()
	if err := u.userUC.RegisterNewUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	userResponse := map[string]any{
		"id":       user.Id,
		"username": user.Username,
		"isActive": user.IsActive,
	}

	c.JSON(http.StatusOK, userResponse)
}

func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.userUC.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(http.StatusOK, gin.H{
		"status": status,
		"data":   users,
	})
}

func NewUserController(r *gin.Engine, usecase usecase.UserUseCase) *UserController {
	controller := UserController{
		router: r,
		userUC: usecase,
	}

	rg := r.Group("/api/v1")
	rg.POST("/users", controller.createHandler)
	rg.GET("/users", controller.listHandler)
	return &controller
}
