package handler

import (
	"log"
	"strconv"
	"time"

	"github.com/AliAstanov/Edu_CRM/models"
	helpers "github.com/AliAstanov/helper"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateStudent(c *gin.Context) {
	var reqBody models.CreateStudent


	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Invalid request Create_student:", err)
		c.JSON(400, gin.H{"error": "invalid request body"})
		return
	}

	student:= &models.Student{
		StudentId: uuid.New(),
		CreatedAt: time.Now(),
	}
	if err := helpers.DataParser1(reqBody, student); err != nil {
		log.Println("Failed to parse request body:", err)
		c.JSON(400, gin.H{"error": "Failed to parse Reqest body"})
		return
	}

	err := h.service.Student().CreateStudent(c, student)
	if err != nil {
		log.Println("Failed to create_student in package handler:", err)
		c.JSON(500, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(201, gin.H{"message": "create student soccessfully"})
}

func (h *handler) GetAllStudents(c *gin.Context) {

	var reqBody models.GetListReq
	var err error

	limit := c.Query("limit")
	page := c.Query("page")

	if reqBody.Limit, err = strconv.Atoi(limit); err != nil {
		log.Println("Invalid limit query param:", err)
		c.JSON(400, gin.H{"error": "Invalid limit query param"})
		return
	}
	if reqBody.Page, err = strconv.Atoi(page); err != nil {
		log.Println("Invalid page query param:", err)
		c.JSON(400, gin.H{"error": "Invalid page query param"})
		return
	}

	students, err := h.service.Student().GetAll(c, &reqBody)
	if err != nil {
		log.Println("Failed to get all students in package handler:", err)
		c.JSON(500, gin.H{"error": "Failed to get all students"})
		return
	}
	c.JSON(200, gin.H{"data": students})
}

func (h *handler) GetStudentById(c *gin.Context) {
	studentId := c.Param("id")

	if studentId == "" {
		log.Println("Student id is empty")
		c.JSON(400, gin.H{"error": "Student id is empty"})
		return
	}

	student, err := h.service.Student().GetByID(c, studentId)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to get user by ID"})
		log.Println("Failed to get user by ID:", err)
		return
	}

	c.JSON(200, gin.H{"data": student})

}

func (h *handler) UpdateStudent(c *gin.Context) {
	studentId := c.Param("id")
	if studentId == "" {
		log.Println("Student id is empty")
		c.JSON(400, gin.H{"error": "Student id is empty"})
		return
	}

	var reqBody models.UpdateStudentReq
	if err := c.BindJSON(&reqBody); err != nil {
		c.JSON(400, gin.H{"error": "Invalid request body"})
		log.Println("Invalid request body:", err)
		return
	}

	respUpdatedData, err := h.service.Student().Update(c,&reqBody,studentId)
	if err !=nil{
		log.Println("Failed to update student:",err)
		c.JSON(500,gin.H{"error":"Failed to update student"})
		return
	}
	c.JSON(200,gin.H{"updated data:":respUpdatedData})
}

func(h *handler)DeleteStudent(c *gin.Context){
	id := c.Param("id")

	if err := h.service.Student().Delete(c,id); err != nil {
		log.Println("Failed to delete student:", err)
		c.JSON(500, gin.H{"error": "Failed to delete student"})
		return
	}
}
