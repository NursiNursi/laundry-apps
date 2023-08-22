package api

import (
	"github.com/NursiNursi/laundry-apps/model"
	"github.com/NursiNursi/laundry-apps/usecase"
	"github.com/gin-gonic/gin"
)

type UomController struct {
	uomUC  usecase.UomUseCase
	router *gin.Engine
}

func (u *UomController) createHandler(c *gin.Context) {
	// inisiasi struct kosong untuk di lakukan bind di body json (POSTMAN)
	var uom model.Uom
	// cek error ketika melakukan bind body JSON, keluarkan status code 400 (bad request - CLIENT)
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// cek error ketikan server tidak merespon atau ada kesalahan, keluarkan status code 500 (internal server error - SERVER)
	// uom.Id = common.GenerateID()
	if err := u.uomUC.RegisterNewUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return // ini harus ada supaya gak diteruskan ke bawah
	}
	// jika semua aman dan tidak ada error
	c.JSON(201, uom)
}

func (u *UomController) listHandler(c *gin.Context) {
	uoms, err := u.uomUC.FindAllUom()
	if err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	// status : code, description
	// data : uoms
	status := map[string]any{
		"code":        200,
		"description": "Get All Data Successfully",
	}
	c.JSON(200, gin.H{
		"status": status,
		"data":   uoms,
	})
}

func (u *UomController) getHandler(c *gin.Context) {
	id := c.Param("id")
	uom, err := u.uomUC.FindByIdUom(id)
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
		"data":   uom,
	})
}

func (u *UomController) updateHandler(c *gin.Context) {
	var uom model.Uom
	if err := c.ShouldBindJSON(&uom); err != nil {
		c.JSON(400, gin.H{"err": err.Error()})
		return
	}
	if err := u.uomUC.UpdateUom(uom); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.JSON(200, uom)
}

func (u *UomController) deleteHandler(c *gin.Context) {
	id := c.Param("id")
	if err := u.uomUC.DeleteUom(id); err != nil {
		c.JSON(500, gin.H{"err": err.Error()})
		return
	}
	c.String(204, "")
}

func NewUomController(usecase usecase.UomUseCase, r *gin.Engine) *UomController {
	controller := UomController{
		router: r,
		uomUC:  usecase,
	}
	// daftarkan semua url path disini
	// /uom -> GET, POST, PUT, DELETE
	// /api/v1/uoms
	rg := r.Group("/api/v1")
	rg.POST("/uoms", controller.createHandler)
	rg.GET("/uoms", controller.listHandler)
	rg.GET("/uoms/:id", controller.getHandler)
	rg.PUT("/uoms", controller.updateHandler)
	rg.DELETE("/uoms/:id", controller.deleteHandler)
	return &controller
}
