
package groupSubjectTeacher

import (
	"context"
	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type GroupSubjectTeacherService struct {
	storage storage.StorageI
}

type GroupSubjectTeacherServiceI interface {
	Create(ctx context.Context, req *models.GroupSubjectTeacher) error
	GetList(ctx context.Context, req *models.GetListReq) (*models.GetGroupSubjectTeachers, error)
	GetByID(ctx context.Context, id string) (*models.GroupSubjectTeacher, error)
	Update(ctx context.Context, req *models.UpdateGroupSubjectTeacher, id string) (*models.GroupSubjectTeacher, error)
	Delete(ctx context.Context, id string) error
}

func NewGroupSubjectTeacherService(storage storage.StorageI) GroupSubjectTeacherServiceI {
	return &GroupSubjectTeacherService{storage: storage}
}

func (s *GroupSubjectTeacherService) Create(ctx context.Context, req *models.GroupSubjectTeacher) error {
	// Kerak bo'lsa validatsiya yoki logger shu yerda yoziladi
	return s.storage.GetGroupSubjectTeacher().CreateGroupSubjectTeacher(ctx, req)
}

func (s *GroupSubjectTeacherService) GetList(ctx context.Context, req *models.GetListReq) (*models.GetGroupSubjectTeachers, error) {
	return s.storage.GetGroupSubjectTeacher().GetListGroupSubjectTeacher(ctx, req)
}

func (s *GroupSubjectTeacherService) GetByID(ctx context.Context, id string) (*models.GroupSubjectTeacher, error) {
	return s.storage.GetGroupSubjectTeacher().GetGroupSubjectTeacher(ctx, id)
}

func (s *GroupSubjectTeacherService) Update(ctx context.Context, req *models.UpdateGroupSubjectTeacher, id string) (*models.GroupSubjectTeacher, error) {
	return s.storage.GetGroupSubjectTeacher().UpdateGroupSubjectTeacher(ctx, req, id)
}

func (s *GroupSubjectTeacherService) Delete(ctx context.Context, id string) error {
	return s.storage.GetGroupSubjectTeacher().DeleteGroupSubjectTeacher(ctx, id)
}
