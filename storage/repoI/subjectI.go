package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type SubjectRepoI interface {
	CreateSubject(ctx context.Context, req *models.Subject) error
	GetListSubject(ctx context.Context, req *models.GetListReq) (*models.GetSubjects, error)
	GetSubject(ctx context.Context, id string) (*models.Subject, error)
	UpdateSubject(ctx context.Context, req *models.UpdateSubjectReq, id string) (*models.Subject, error)
	DeleteSubject(ctx context.Context, id string) error
}
