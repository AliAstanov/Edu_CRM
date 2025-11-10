package handler

import (
	"log"
	"strconv"

	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func (h *handler) CreateGroupSubjectTeacher(c *gin.Context) {
	var reqBody models.CreateGroupSubjectTeacher
	ctx := c.Request.Context()

	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Invalid bind data:", err)
		c.JSON(400, gin.H{"Error": "Invalid bind data"})
		return
	}

	groupSbjTeacher := &models.GroupSubjectTeacher{
		ID:        uuid.New(),
		GroupID:   reqBody.GroupID,
		SubjectID: reqBody.SubjectID,
		TeacherID: reqBody.TeacherID,
		StartDate: reqBody.StartDate,
		EndDate:   reqBody.EndDate,
	}

	err := h.service.GroupSubjectTeacher().Create(ctx, groupSbjTeacher)
	if err != nil {
		log.Println("Invalid Create groupSubjectTeacher:", err)
		c.JSON(500, gin.H{"error": "Invalid create groupSubjectTeacher"})
		return
	}

	c.JSON(201, gin.H{"message": "Create groupSubjectTeacher soccessfully"})

}

func (h *handler) GetListGroupSubjectTeacher(c *gin.Context) {
	limit := c.Query("limit")
	page := c.Query("page")
	var err error
	var req models.GetListReq
	ctx := c.Request.Context()

	req.Limit, err = strconv.Atoi(limit)
	if err != nil {
		log.Println("nvalid to parse limit")
		c.JSON(400, gin.H{"error": "Invalid to parse limit"})
	}
	req.Page, err = strconv.Atoi(page)
	if err != nil {
		log.Println("Invalid to parse  page")
		c.JSON(400, gin.H{"error": "Invalid to pars page"})
	}

	dataList, err := h.service.GroupSubjectTeacher().GetList(ctx, &req)
	if err != nil {
		log.Println("Invalid to getListGroupSubjectTeacher:", err)
		c.JSON(500, gin.H{"error": "Invalid to GetListGroupTeacher"})
		return
	}

	c.JSON(200, dataList)
}

func (h *handler) GetByIdGroupSubjectTeacher(c *gin.Context) {
	ctx := c.Request.Context()
	id := c.Param("id")
	if id == "" {
		log.Println("id is empty")
		c.JSON(400, gin.H{"error": "id is empty"})
		return
	}

	data, err := h.service.GroupSubjectTeacher().GetByID(ctx, id)
	if err != nil {
		log.Println("Invalid GetByIdGroupSubjectTeacher:", err)
		c.JSON(500, gin.H{"error": "Ivalid GetByIdGroupSubjectTeacher"})
		return
	}

	c.JSON(200, data)
}

func (h *handler) UpdateGroupSubjectTeacher(c *gin.Context) {
	ctx := c.Request.Context()
	var reqBody models.UpdateGroupSubjectTeacher
	id  := c.Param("id")

	if err := c.BindJSON(&reqBody); err != nil {
		log.Println("Failed bind data to update groupSubjectTeacher:", err)
		c.JSON(400,gin.H{"error":"Fai;ed bind request"})
		return
	}

	data, err := h.service.GroupSubjectTeacher().Update(ctx,&reqBody, id)
	if err != nil {
		log.Println("Failed to update GroupSubjectTeacher:",err)
		c.JSON(500,gin.H{"error":"Failed to update GroupSubjectTeacher"})
		return
	}
	c.JSON(200,gin.H{"update group_subject_teacher soccessfully:":data})
}
