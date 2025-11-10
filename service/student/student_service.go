package student

import (
	"context"
	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type StudentService struct {
	storage storage.StorageI
}

type StudentServiceI interface {
	CreateStudent(ctx context.Context, student *models.Student) error
	GetByID(ctx context.Context, id string) (*models.Student, error)
	GetAll(ctx context.Context, req *models.GetListReq) (*models.GetStudens, error)
	Update(ctx context.Context, req *models.UpdateStudentReq, id string) (*models.Student, error)
	Delete(ctx context.Context, id string) error
}


func NewStudentService(storage storage.StorageI) StudentServiceI {
	return &StudentService{storage: storage}
}

func (s *StudentService) CreateStudent(ctx context.Context, student *models.Student) error {
	// AI log, validatsiya, log yozish shu yerda bo'ladi
	return s.storage.GetStudent().CreateStudent(ctx, student)
}

func (s *StudentService) GetByID(ctx context.Context, id string) (*models.Student, error) {
	return s.storage.GetStudent().GetStudent(ctx, id)
}

func (s *StudentService) GetAll(ctx context.Context, req *models.GetListReq) (*models.GetStudens, error) {
	return s.storage.GetStudent().GetListStudent(ctx,req)
}

func (s *StudentService) Update(ctx context.Context, student *models.UpdateStudentReq, id string) (*models.Student, error) {
	return s.storage.GetStudent().UpdateStudent(ctx, student, id)
}

func (s *StudentService) Delete(ctx context.Context, id string) error {
	return s.storage.GetStudent().DeleteStudent(ctx, id)
}
