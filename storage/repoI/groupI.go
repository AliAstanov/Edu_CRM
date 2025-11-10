package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type GroupI interface {
	CreateGroup(ctx context.Context, req *models.Group) error
	GetListGroup(ctx context.Context, req *models.GetListReq) (*models.GetGroups, error)
	GetGroup(ctx context.Context, id string) (*models.Group, error)
	UpdateGroup(ctx context.Context, req *models.UpdateGroupReq, id string) (*models.Group, error)
	DeleteGroup(ctx context.Context, id string) error
}
