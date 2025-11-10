package handler

import (
	"log"
	"strconv"
	"time"

	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateGroup(c *gin.Context) {
	var reqBody models.CreateGroup
	ctx := c.Request.Context()
	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request body in CreateGroup:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	group := models.Group{
		ID:        uuid.New(),
		Name:      reqBody.Name,
		CreatedAt: time.Now(),
	}

	if err := h.service.Group().CreateGroup(ctx, &group); err != nil {
		log.Println("Failed to create group:", err)
		c.JSON(500, gin.H{"Error": "Failed to create group"})
		return
	}

	c.JSON(201, gin.H{"message": "create group soccessfully"})
}

func (h *handler) GetListGroup(c *gin.Context) {
	ctx := c.Request.Context()
	var err error
	var reqBody models.GetListReq
	limit := c.Query("limit")
	page := c.Query("page")

	reqBody.Limit, err = strconv.Atoi(limit)
	if err != nil {
		log.Println("Failed to parse limit:", err)
		c.JSON(400, gin.H{"error": "failed to parse limit"})
		return
	}
	reqBody.Page, err = strconv.Atoi(page)
	if err != nil {
		log.Println("Failed to parse page:", err)
		c.JSON(400, gin.H{"error": "failed to parse page"})
		return
	}

	groups, err := h.service.Group().GetListGroup(ctx, &reqBody)
	if err != nil {
		log.Println("Failed to get list group:", err)
		c.JSON(500, gin.H{"error": "Failed to get list group"})
		return
	}

	c.JSON(200, groups)

}

func (h *handler) GetGroupById(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")

	if id == "" {
		log.Println("Failed to get id")
		c.JSON(400, gin.H{"error": "Failed to get id"})
		return
	}

	group, err := h.service.Group().GetGroup(ctx, id)
	if err != nil {
		log.Println("Failed to get group by id:", err)
		c.JSON(500, gin.H{"error": "Failed to get group by id"})
		return
	}

	c.JSON(200, group)
}

func (h *handler) UpdateGroup(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	var reqBody models.UpdateGroupReq

	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed to bind json:", err)
		c.JSON(400, gin.H{"error": "Failed to bind json"})
		return
	}

	group, err := h.service.Group().UpdateGroup(ctx, &reqBody, id)
	if err != nil {
		log.Println("Failed to update group:", err)
		c.JSON(500, gin.H{"error": "Failed to update group"})
		return
	}

	c.JSON(200, group)
}

func (h *handler) DeleteGroup(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		log.Println("Failed to get id")
		c.JSON(400, gin.H{"error": "Failed to get id"})
		return
	}

	err := h.service.Group().DeleteGroup(ctx, id)
	if err != nil {
		log.Println("Failed to delete group:", err)
		c.JSON(500, gin.H{"error": "Failed to delete group"})
		return
	}
	c.JSON(200, gin.H{"message": "Group deleted successfully"})
}


