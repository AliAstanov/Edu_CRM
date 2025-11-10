package models

type StudentGroup struct {
	ID        string `json:"id"`
	StudentID string `json:"student_id"`
	GroupID   string `json:"group_id"`
}

type CreateStudentGroup struct {
	StudentID string `json:"student_id"`
	GroupID   string `json:"group_id"`
}

type UpdateStudentGroup struct {
	StudentID string `json:"student_id"`
	GroupID   string `json:"group_id"`
}

type GetListStudentGroup struct {
	List  []StudentGroup `json:"student_group"`
	Count int            `json:"count"`
}
