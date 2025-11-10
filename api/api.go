package api

import (
	"os"

	"github.com/AliAstanov/Edu_CRM/api/handler"
	"github.com/AliAstanov/Edu_CRM/service"
	"github.com/gin-gonic/gin"
)

func Api(service service.ServiceI) {
	router := gin.Default()
	h := handler.NewHandler(service)

	// API versiya yo'nalishi
	v1 := router.Group("/v1")

	// Student route group
	student := v1.Group("/students")
	{
		student.POST("/", h.CreateStudent)
		student.GET("/", h.GetAllStudents)
		student.GET("/:id", h.GetStudentById)
		student.PUT("/:id", h.UpdateStudent)
		student.DELETE("/:id", h.DeleteStudent)
	}

	// Subject route group
	subject := v1.Group("/subjects")
	{
		subject.POST("/", h.CreateSubject)
		subject.GET("/", h.GetAllSubjects)
		subject.GET("/:id", h.GetByIdSubject)
		subject.PUT("/:id", h.UpdateSubject)
		subject.DELETE("/:id", h.DeleteSubject)
	}

	// Teacher route group
	teacher := v1.Group("/teachers")
	{
		teacher.POST("/", h.CreateTeacher)
		teacher.GET("/", h.GetAllTeachers)
		teacher.GET("/:id", h.GetByIdTeacher)
		teacher.PUT("/:id", h.UpdateTeacher)
		teacher.DELETE("/:id", h.DeleteTeacher)
	}

	// Group route group
	group := v1.Group("/groups")
	{
		group.POST("/", h.CreateGroup)
		group.GET("/", h.GetListGroup)
		group.GET("/:id", h.GetGroupById)
		group.PUT("/:id", h.UpdateGroup)
		group.DELETE("/:id", h.DeleteGroup)
	}

	//GroupSubjectTeacher route group
	grp_sbj_teacher := v1.Group("/grp_sbj_teachers")
	{
		grp_sbj_teacher.POST("/create_grp_sbj_teacher", h.CreateGroupSubjectTeacher)
		grp_sbj_teacher.GET("/get_all_grp_sbj_teacher", h.GetListGroupSubjectTeacher)
		grp_sbj_teacher.GET("/get_by_id_grp_sbj_teacher", h.GetByIdGroupSubjectTeacher)
		grp_sbj_teacher.PUT("/update_grp_sbj_teacher", h.UpdateGroupSubjectTeacher)
	}


	// Ping (test)
	router.GET("/ping", h.Ping)

	// Run server
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	} else {
		port = ":" + port
	}
	router.Run(port)
}
