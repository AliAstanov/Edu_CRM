package subject

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type SubjectService struct {
	storage storage.StorageI
}
type SubjectServiceI interface {
	Create(ctx context.Context, req *models.Subject) error
	GetByID(ctx context.Context, id string) (*models.Subject, error)
	GetAll(ctx context.Context, req *models.GetListReq) (*models.GetSubjects, error)
	Update(ctx context.Context, req *models.UpdateSubjectReq, id string) (*models.Subject, error)
	Delete(ctx context.Context, id string) error
}

func NewSubjectService(storage storage.StorageI) SubjectServiceI {
	return &SubjectService{storage: storage}
}

func(s *SubjectService)Create(ctx context.Context, req *models.Subject) error{
	return s.storage.GetSubject().CreateSubject(ctx, req)
}
func(s *SubjectService)GetByID(ctx context.Context, id string) (*models.Subject, error){
	return s.storage.GetSubject().GetSubject(ctx,id)
}
func(s *SubjectService)GetAll(ctx context.Context, req *models.GetListReq) (*models.GetSubjects, error){
	return s.storage.GetSubject().GetListSubject(ctx,req)
}
func(s *SubjectService)Update(ctx context.Context, req *models.UpdateSubjectReq, id string) (*models.Subject, error){
	return s.storage.GetSubject().UpdateSubject(ctx,req,id)
}
func(s *SubjectService)Delete(ctx context.Context, id string) error{
	return s.storage.GetSubject().DeleteSubject(ctx,id)
}
