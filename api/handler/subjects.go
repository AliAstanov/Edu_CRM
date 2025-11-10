package handler

import (
	"log"
	"net/http"
	"strconv"

helpers "github.com/AliAstanov/helper"
	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateSubject(c *gin.Context) {
	ctx := c.Request.Context()
	var reqBody models.CreateSubject

	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed Bind Json data:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}
	subject := &models.Subject{
		ID:   uuid.New(),
		Name: reqBody.Name,
	}

	if err := helpers.DataParser1(reqBody,subject); err != nil {
		log.Println("Failed pars data:", err)
		c.JSON(400, gin.H{"error": "Invalid request body"})
		return
	}

	err := h.service.Subject().Create(ctx, subject)
	if err != nil {
		log.Println("Failed Create Subject:", err)
		c.JSON(400, gin.H{"error": "Failed to Create Subject"})
		return
	}

	log.Println("create subject soccessfully")
	c.JSON(201, subject)
}

func (h *handler) GetAllSubjects(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	var err error
	var req models.GetListReq
	ctx := c.Request.Context()

	if req.Limit, err = strconv.Atoi(limit); err != nil {
		log.Println("Failed to parse limit:", err)
		c.JSON(400, gin.H{"error": "Invalid limit"})
		return
	}
	if req.Page, err = strconv.Atoi(page); err != nil {
		log.Println("Failed to parse page")
		c.JSON(400, gin.H{"error": "Invalid page"})
		return
	}

	subjects, err := h.service.Subject().GetAll(ctx, &req)
	if err != nil {
		log.Println("Failed Get all subjects:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request get all subjects"})
		return
	}
	c.JSON(200, subjects)
}

func (h *handler) GetByIdSubject(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if id == "" {
		log.Println("Failed to get id")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}

	subject, err := h.service.Subject().GetByID(ctx, id)
	if err != nil {
		log.Println("Failed Get subject by id:", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request get subject by id"})
		return
	}
	c.JSON(200, subject)
}

func (h *handler) UpdateSubject(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()
	var reqBody models.UpdateSubjectReq
	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed to bind json:", err)
		c.JSON(400, gin.H{"error": "Invalid request"})
		return
	}

	respUpdatedData, err := h.service.Subject().Update(ctx, &reqBody, id)
	if err != nil {
		log.Println("Failed to update subject")
		c.JSON(500, gin.H{"error": "Invalid update subject"})
		return
	}

	c.JSON(200, respUpdatedData)
}

func (h *handler) DeleteSubject(c *gin.Context) {
	id := c.Param("id")
	ctx := c.Request.Context()

	if id == "" {
		log.Println("empty data to id:")
		c.JSON(400, gin.H{"error": "Invalid id"})
		return
	}
	if err := h.service.Subject().Delete(ctx, id); err != nil {
		log.Println("Failed to delete subject:", err)
		c.JSON(500, gin.H{"error": "Invalid delete subject"})
		return
	}

	c.JSON(200, gin.H{"message": "Subject deleted successfully"})
}
