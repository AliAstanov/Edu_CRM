package service

import (
	"github.com/AliAstanov/Edu_CRM/service/group"
	"github.com/AliAstanov/Edu_CRM/service/groupSubjectTeacher"
	"github.com/AliAstanov/Edu_CRM/service/student"
	"github.com/AliAstanov/Edu_CRM/service/subject"
	teacher "github.com/AliAstanov/Edu_CRM/service/teacherService"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type Service struct {
	student             student.StudentServiceI
	subject             subject.SubjectServiceI
	teacher             teacher.TeacherServiceI
	group               group.GroupServiceI
	groupSubjectTeacher groupSubjectTeacher.GroupSubjectTeacherServiceI
}

type ServiceI interface {
	Student() student.StudentServiceI
	Subject() subject.SubjectServiceI
	Teacher() teacher.TeacherServiceI
	Group() group.GroupServiceI
	GroupSubjectTeacher() groupSubjectTeacher.GroupSubjectTeacherServiceI
}

func NewService(storage storage.StorageI) ServiceI {
	return &Service{
		student: student.NewStudentService(storage),
		subject: subject.NewSubjectService(storage),
		teacher: teacher.NewTeacherService(storage),
		group:   group.NewGroupService(storage),
		groupSubjectTeacher: groupSubjectTeacher.NewGroupSubjectTeacherService(storage),
	}
}

func (s *Service) Student() student.StudentServiceI {
	return s.student
}
func (s *Service) Subject() subject.SubjectServiceI {
	return s.subject
}
func (s *Service) Teacher() teacher.TeacherServiceI {
	return s.teacher
}

func (s *Service) Group() group.GroupServiceI {
	return s.group
}

func (s *Service) GroupSubjectTeacher() groupSubjectTeacher.GroupSubjectTeacherServiceI {
	return s.groupSubjectTeacher
}
