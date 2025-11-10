package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type TeacherRepoI interface {
	CreateTeacher(ctx context.Context, req *models.Teacher) error
	GetListTeacher(ctx context.Context, req *models.GetListReq) (*models.GetTeachers, error)
	GetTeacher(ctx context.Context, id string) (*models.Teacher, error)
	UpdateTeacher(ctx context.Context, req *models.UpdateTeacherReq, id string) (*models.Teacher, error)
	DeleteTeacher(ctx context.Context, id string) error
}
