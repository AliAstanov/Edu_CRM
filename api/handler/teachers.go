package handler

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/AliAstanov/Edu_CRM/models"
	helpers "github.com/AliAstanov/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateTeacher(c *gin.Context) {
	var reqBody models.CreateTeacher
	ctx := c.Request.Context()
	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed to bind JSON:", err)
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	teacher := &models.Teacher{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
	}

	if err := helpers.DataParser1(reqBody, teacher); err != nil {
		log.Println("Failed pars data:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.service.Teacher().Create(ctx, teacher)
	if err != nil {
		log.Println("Failed to create teacher:", err)
		c.JSON(400, gin.H{"error": "Failed to create teacher"})
		return
	}

	c.JSON(201, gin.H{"message": "Teacher created successfully"})
}

func (h *handler) GetAllTeachers(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	ctx := c.Request.Context()
	var req models.GetListReq
	var err error

	req.Limit, err = strconv.Atoi(limit)
	if err != nil {
		log.Println("Failed to parse limit")
		c.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}

	req.Page, err = strconv.Atoi(page)
	if err != nil {
		log.Println("Failed to parse limit")
		c.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}

	teachers, err := h.service.Teacher().GetAll(ctx, &req)
	if err != nil {
		log.Println("Failed Get all subjects:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request get all subjects"})
		return
	}

	c.JSON(200, teachers)
}

func (h *handler) GetByIdTeacher(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		log.Println("id is empty")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	teacher, err := h.service.Teacher().GetByID(ctx, id)
	if err != nil {
		log.Println("Failed to get teacher")
		c.JSON(500, gin.H{"error": "Invalid request get teacher"})
		return
	}

	c.JSON(200, teacher)

}

func (h *handler) UpdateTeacher(c *gin.Context) {
	var reqBody models.UpdateTeacherReq
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		log.Println("id is empty")
		c.JSON(400, "id is empty")
		return
	}

	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed to bind json")
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	teacher, err := h.service.Teacher().Update(ctx, &reqBody, id)
	if err != nil {
		log.Println("Failed to update teacher:", err)
		c.JSON(500, gin.H{"error": "Invalid request update teacher"})
		return
	}

	c.JSON(200, teacher)
}

func (h *handler) DeleteTeacher(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		log.Println("id is empty to delete teacher")
		c.JSON(400, gin.H{"error": "id is empty "})
		return
	}
	err := h.service.Teacher().Delete(ctx, id)
	if err != nil {
		log.Println("Failed to delete teacher")
		c.JSON(500, gin.H{"error": "Invalid delete teacher"})
		return
	}

	c.JSON(200, gin.H{"message": "delete teacher successfully"})
}
