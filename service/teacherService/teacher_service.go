package teacher

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type TeacherService struct {
	storage storage.StorageI
}

type TeacherServiceI interface {
	Create(ctx context.Context, req *models.Teacher) error
	GetByID(ctx context.Context, id string) (*models.Teacher, error)
	GetAll(ctx context.Context, req *models.GetListReq) (*models.GetTeachers, error)
	Update(ctx context.Context, req *models.UpdateTeacherReq, id string) (*models.Teacher, error)
	Delete(ctx context.Context, id string) error
}

func NewTeacherService(storage storage.StorageI) TeacherServiceI {
	return &TeacherService{storage: storage}
}

func (t *TeacherService) Create(ctx context.Context, req *models.Teacher) error {
	return t.storage.GetTeacher().CreateTeacher(ctx, req)
}

func (t *TeacherService) GetByID(ctx context.Context, id string) (*models.Teacher, error) {
	return t.storage.GetTeacher().GetTeacher(ctx, id)
}

func (t *TeacherService) GetAll(ctx context.Context, req *models.GetListReq) (*models.GetTeachers, error) {
	return t.storage.GetTeacher().GetListTeacher(ctx, req)
}

func (t *TeacherService) Update(ctx context.Context, req *models.UpdateTeacherReq, id string) (*models.Teacher, error) {
	return t.storage.GetTeacher().UpdateTeacher(ctx, req, id)
}

func (t *TeacherService) Delete(ctx context.Context, id string) error {
	return t.storage.GetTeacher().DeleteTeacher(ctx, id)
}
