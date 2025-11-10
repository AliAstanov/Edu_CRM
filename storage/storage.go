package storage

import (
	"github.com/AliAstanov/Edu_CRM/storage/postgres"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Storage struct {
	StudentRepo repoi.StudentRepoI
	SubjectRepo repoi.SubjectRepoI
	TeacherRepo repoi.TeacherRepoI
	GroupRepo   repoi.GroupI
	GroupSubjectTeacherRepo repoi.GroupSubjectTeacherI
}

type StorageI interface {
	GetStudent() repoi.StudentRepoI
	GetSubject() repoi.SubjectRepoI
	GetTeacher() repoi.TeacherRepoI
	GetGroup() repoi.GroupI
	GetGroupSubjectTeacher() repoi.GroupSubjectTeacherI
}

func NewStorage(db *pgxpool.Pool) StorageI {
	return &Storage{
		StudentRepo: postgres.NewStudentRepo(db),
		SubjectRepo: postgres.NewSubject(db),
		TeacherRepo: postgres.NewTeacherRepo(db),
		GroupRepo: postgres.NewGroupRepo(db),
		GroupSubjectTeacherRepo: postgres.NewGroupSubjectTeacher(db),
	}
}

func (s *Storage) GetStudent() repoi.StudentRepoI {
	return s.StudentRepo
}

func (s *Storage) GetSubject() repoi.SubjectRepoI {
	return s.SubjectRepo
}

func (s *Storage) GetTeacher() repoi.TeacherRepoI {
	return s.TeacherRepo
}

func (s *Storage) GetGroup() repoi.GroupI {
	return s.GroupRepo
}

func (s *Storage) GetGroupSubjectTeacher() repoi.GroupSubjectTeacherI{
	return  s.GroupSubjectTeacherRepo
}