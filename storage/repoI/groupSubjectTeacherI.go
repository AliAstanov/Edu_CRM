package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type GroupSubjectTeacherI interface {
	CreateGroupSubjectTeacher(ctx context.Context, req *models.GroupSubjectTeacher) error
	GetListGroupSubjectTeacher(ctx context.Context, req *models.GetListReq) (*models.GetGroupSubjectTeachers, error)
	GetGroupSubjectTeacher(ctx context.Context, id string) (*models.GroupSubjectTeacher, error)
	UpdateGroupSubjectTeacher(ctx context.Context, req *models.UpdateGroupSubjectTeacher, id string) (*models.GroupSubjectTeacher, error)
	DeleteGroupSubjectTeacher(ctx context.Context, id string) error
}
