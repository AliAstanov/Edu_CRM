package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type StudentRepoI interface {
	CreateStudent(ctx context.Context, req *models.Student)error
	GetListStudent(ctx context.Context, req *models.GetListReq)(*models.GetStudens, error)
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	UpdateStudent(ctx context.Context, req *models.UpdateStudentReq, id string)(*models.Student, error)
	DeleteStudent(ctx context.Context, id string)error
}
