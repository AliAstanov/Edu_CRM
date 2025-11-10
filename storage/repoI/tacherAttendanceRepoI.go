package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type TeacherAttendanceI interface {
	CreateTeacherAttendance(ctx context.Context, req *models.TeacherAttendance) error
	GetListTeacherAttendance(ctx context.Context, req *models.GetListReq) (*models.GetTeacherAttendances, error)
	GetTeacherAttendance(ctx context.Context, id string) (*models.TeacherAttendance, error)
	UpdateTeacherAttendance(ctx context.Context, req *models.UpdateTeacherAttendanceReq, id string) (*models.TeacherAttendance, error)
	DeleteTeacherAttendance(ctx context.Context, id string) error
}
