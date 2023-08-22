package api

import (
	"net/http"
	"strconv"

	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/model/dto"
	"github.com/NursiNursi/laundry-apps/usecase"
	"github.com/NursiNursi/laundry-apps/utils/common"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router  *gin.Engine
	usecase usecase.CustomerUseCase
}

func (cc *CustomerController) createHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	customer.Id = common.GenerateID()
	if err := cc.usecase.RegisterNewCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, customer)
}
func (cc *CustomerController) listHandler(c *gin.Context) {
	page, _ := strconv.Atoi(c.Query("page"))
	limit, _ := strconv.Atoi(c.Query("limit"))
	paginationParam := dto.PaginationParam{
		Page:  page,
		Limit: limit,
	}
	customers, paging, err := cc.usecase.FindAllCustomer(paginationParam)
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
		"data":   customers,
		"paging": paging,
	})
}
func (cc *CustomerController) getHandler(c *gin.Context) {
	id := c.Param("id")
	customer, err := cc.usecase.FindByIdCustomer(id)
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	status := map[string]any{
		"code":        200,
		"description": "Get By Id Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   customer,
	})
}
func (cc *CustomerController) updateHandler(c *gin.Context) {
	var customer model.Customer
	if err := c.ShouldBindJSON(&customer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}

	if err := cc.usecase.UpdateCustomer(customer); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	c.JSON(http.StatusOK, customer)
}
func (cc *CustomerController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := cc.usecase.DeleteCustomer(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewCustomerController(r *gin.Engine, usecase usecase.CustomerUseCase) *CustomerController {
	controller := CustomerController{
		router:  r,
		usecase: usecase,
	}
	rg := r.Group("/api/v1")
	rg.POST("/customers", controller.createHandler)
	rg.GET("/customers", controller.listHandler)
	rg.GET("/customers/:id", controller.getHandler)
	rg.PUT("/customers", controller.updateHandler)
	rg.DELETE("/customers/:id", controller.deleteHandler)
	return &controller
}
