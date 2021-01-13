package controllers

import (
	"context"
	"fmt"
	"strconv"

	"github.com/team06/app/ent/prefix"

	"github.com/gin-gonic/gin"
	"github.com/team06/app/ent"
)

// PrefixController defines the struct for the prefix controller
type PrefixController struct {
	client *ent.Client
	router gin.IRouter
}

// CreatePrefix handles POST requests for adding prefix entities
// @Summary Create prefix
// @Description Create prefix
// @ID create-prefix
// @Accept   json
// @Produce  json
// @Param prefix body ent.Prefix true "Prefix entity"
// @Success 200 {object} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs [post]
func (ctl *PrefixController) CreatePrefix(c *gin.Context) {
	obj := ent.Prefix{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "prefix binding failed",
		})
		return
	}

	prefixs, err := ctl.client.Prefix.
		Create().
		SetPrefixValue(obj.PrefixValue).
		Save(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": "saving failed",
		})
		return
	}

	c.JSON(200, prefixs)
}

// GetPrefix handles GET requests to retrieve a prefix entity
// @Summary Get a prefix entity by ID
// @Description get prefix by ID
// @ID get-prefix
// @Produce  json
// @Param id path int true "Prefix ID"
// @Success 200 {object} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs/{id} [get]
func (ctl *PrefixController) GetPrefix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}
	prefixs, err := ctl.client.Prefix.
		Query().
		Where(prefix.IDEQ(int(id))).
		Only(context.Background())

	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prefixs)
}

// ListPrefix handles request to get a list of prefix entities
// @Summary List prefix entities
// @Description list prefix entities
// @ID list-prefix
// @Produce json
// @Param limit  query int false "Limit"
// @Param offset query int false "Offset"
// @Success 200 {array} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs [get]
func (ctl *PrefixController) ListPrefix(c *gin.Context) {
	limitQuery := c.Query("limit")
	limit := 10
	if limitQuery != "" {
		limit64, err := strconv.ParseInt(limitQuery, 10, 64)
		if err == nil {
			limit = int(limit64)
		}
	}

	offsetQuery := c.Query("offset")
	offset := 0
	if offsetQuery != "" {
		offset64, err := strconv.ParseInt(offsetQuery, 10, 64)
		if err == nil {
			offset = int(offset64)
		}
	}

	prefixs, err := ctl.client.Prefix.
		Query().
		Limit(limit).
		Offset(offset).
		All(context.Background())

	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, prefixs)
}

// DeletePrefix handles DELETE requests to delete a prefix entity
// @Summary Delete a prefix entity by ID
// @Description get prefix by ID
// @ID delete-prefix
// @Produce  json
// @Param id path int true "Prefix ID"
// @Success 200 {object} gin.H
// @Failure 400 {object} gin.H
// @Failure 404 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs/{id} [delete]
func (ctl *PrefixController) DeletePrefix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	err = ctl.client.Prefix.
		DeleteOneID(int(id)).
		Exec(context.Background())
	if err != nil {
		c.JSON(404, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{"result": fmt.Sprintf("ok deleted %v", id)})
}

// UpdatePrefix handles PUT requests to update a prefix entity
// @Summary Update a prefix entity by ID
// @Description update prefix by ID
// @ID update-prefix
// @Accept   json
// @Produce  json
// @Param id path int true "Prefix ID"
// @Param prefix body ent.Prefix true "Prefix entity"
// @Success 200 {object} ent.Prefix
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /prefixs/{id} [put]
func (ctl *PrefixController) UpdatePrefix(c *gin.Context) {
	id, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	obj := ent.Prefix{}
	if err := c.ShouldBind(&obj); err != nil {
		c.JSON(400, gin.H{
			"error": "prefix binding failed",
		})
		return
	}
	obj.ID = int(id)
	prefixs, err := ctl.client.Prefix.
		UpdateOne(&obj).
		Save(context.Background())
	if err != nil {
		c.JSON(400, gin.H{
			"error": "update failed",
		})
		return
	}

	c.JSON(200, prefixs)
}

// NewPrefixController creates and registers handles for the prefix controller
func NewPrefixController(router gin.IRouter, client *ent.Client) *PrefixController {
	prefixcontroller := &PrefixController{
		client: client,
		router: router,
	}

	prefixcontroller.register()

	return prefixcontroller

}

// InitPrefixController registers routes to the main engine
func (ctl *PrefixController) register() {
	prefixs := ctl.router.Group("/prefixs")

	prefixs.GET("", ctl.ListPrefix)

	// CRUD
	prefixs.POST("", ctl.CreatePrefix)
	prefixs.GET(":id", ctl.GetPrefix)
	prefixs.PUT(":id", ctl.UpdatePrefix)
	prefixs.DELETE(":id", ctl.DeletePrefix)
}
