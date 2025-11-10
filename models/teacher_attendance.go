package models

import (
	"time"

	"github.com/google/uuid"
)

type TeacherAttendance struct {
	ID         uuid.UUID `json:"id"`
	TeacherID  uuid.UUID `json:"teacher_id"`
	Date       time.Time `json:"date"`
	Status     string    `json:"status"` // 'present', 'absent', 'late'
	Note       string    `json:"note,omitempty"`
	IsExcused  bool      `json:"is_excused"`
}

type GetTeacherAttendances struct {
	Attendances []TeacherAttendance `json:"attendances"`
	Count       int                 `json:"count"`
}

type CreateTeacherAttendance struct {
	TeacherID uuid.UUID `json:"teacher_id" binding:"required"`
	Date      string    `json:"date" binding:"required"` // Masalan: "2025-07-15", keyinchalik time.Parse bilan aylantiriladi
	Status    string    `json:"status" binding:"required,oneof=present absent late"`
	Note      string    `json:"note"`
	IsExcused bool      `json:"is_excused"`
}

type UpdateTeacherAttendanceReq struct {
	Date      string `json:"date"` // ixtiyoriy yangilash
	Status    string `json:"status"`
	Note      string `json:"note"`
	IsExcused *bool  `json:"is_excused"` // nil bo‘lishi mumkin, faqat o‘zgartirilsa yuboriladi
}
