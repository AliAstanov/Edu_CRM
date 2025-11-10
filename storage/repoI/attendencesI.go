package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type AttendanceI interface {
	CreateAttendance(ctx context.Context, req *models.Attendance) error
	GetListAttendance(ctx context.Context, req *models.GetListReq) (*models.GetListAttendance, error)
	GetAttendance(ctx context.Context, id string) (*models.Attendance, error)
	UpdateAttendance(ctx context.Context, req *models.UpdateAttendance, id string) (*models.Attendance, error)
	DeleteAttendance(ctx context.Context, id string) error
}
