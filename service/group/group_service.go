package group

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
	"github.com/AliAstanov/Edu_CRM/storage"
)

type GroupService struct {
	storage storage.StorageI
}

type GroupServiceI interface {
	CreateGroup(ctx context.Context, req *models.Group) error
	GetListGroup(ctx context.Context, req *models.GetListReq) (*models.GetGroups, error)
	GetGroup(ctx context.Context, id string) (*models.Group, error)
	UpdateGroup(ctx context.Context, req *models.UpdateGroupReq, id string) (*models.Group, error)
	DeleteGroup(ctx context.Context, id string) error
}

func NewGroupService(storage storage.StorageI) GroupServiceI {
	return &GroupService{storage: storage}
}

func (s *GroupService) CreateGroup(ctx context.Context, req *models.Group) error {
	// Validatsiya, log, AI log yozish shu yerda bo'lishi mumkin
	return s.storage.GetGroup().CreateGroup(ctx, req)
}

func (s *GroupService) GetListGroup(ctx context.Context, req *models.GetListReq) (*models.GetGroups, error) {
	return s.storage.GetGroup().GetListGroup(ctx, req)
}

func (s *GroupService) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	return s.storage.GetGroup().GetGroup(ctx, id)
}

func (s *GroupService) UpdateGroup(ctx context.Context, req *models.UpdateGroupReq, id string) (*models.Group, error) {
	return s.storage.GetGroup().UpdateGroup(ctx, req, id)
}

func (s *GroupService) DeleteGroup(ctx context.Context, id string) error {
	return s.storage.GetGroup().DeleteGroup(ctx, id)
}
