package models

import "time"
//Davomat
type Attendance struct {
	ID        string    `json:"id"`
	StudentID string    `json:"student_id"`
	GroupID   string    `json:"group_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"` // 'present' yoki 'absent'
}

type CreateAttendance struct {
	StudentID string    `json:"student_id" binding:"required"`
	GroupID   string    `json:"group_id" binding:"required"`
	Date      time.Time `json:"date" binding:"required"`
	Status    string    `json:"status" binding:"required"` // gin validatsiya uchun
}

type UpdateAttendance struct {
	StudentID string    `json:"student_id"`
	GroupID   string    `json:"group_id"`
	Date      time.Time `json:"date"`
	Status    string    `json:"status"` // 'present' yoki 'absent'
}

type GetListAttendance struct {
	List  []Attendance `json:"attendances"`
	Count int          `json:"count"`
}
