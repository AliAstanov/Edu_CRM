package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type StudentGroupI interface {
	CreateStudentGroup(ctx context.Context, req *models.StudentGroup) error
	GetListStudentGroup(ctx context.Context, req *models.GetListReq) (*models.GetListStudentGroup, error)
	GetStudentGroup(ctx context.Context, id string) (*models.StudentGroup, error)
	UpdateStudentGroup(ctx context.Context, req *models.UpdateStudentGroup, id string) (*models.StudentGroup, error)
	DeleteStudentGroup(ctx context.Context, id string) error
}
