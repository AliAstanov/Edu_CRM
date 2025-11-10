package models

import (
	"time"

	"github.com/google/uuid"
)

type GroupSubjectTeacher struct {
	ID        uuid.UUID  `json:"id"`
	GroupID   uuid.UUID  `json:"group_id"`
	SubjectID uuid.UUID  `json:"subject_id"`
	TeacherID *uuid.UUID `json:"teacher_id"` // nullable
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type CreateGroupSubjectTeacher struct {
	GroupID   uuid.UUID  `json:"group_id" binding:"required"`
	SubjectID uuid.UUID  `json:"subject_id" binding:"required"`
	TeacherID *uuid.UUID `json:"teacher_id"`
	StartDate *time.Time `json:"start_date"`
	EndDate   *time.Time `json:"end_date"`
}

type UpdateGroupSubjectTeacher struct {
	GroupID   *uuid.UUID  `json:"group_id"`
	SubjectID *uuid.UUID  `json:"subject_id"`
	TeacherID *uuid.UUID  `json:"teacher_id"`
	StartDate *time.Time  `json:"start_date"`
	EndDate   *time.Time  `json:"end_date"`
}

type GetGroupSubjectTeachers struct {
	List  []GroupSubjectTeacher `json:"list"`
	Count int                   `json:"count"`
}
