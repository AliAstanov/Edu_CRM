package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AttendanceRepo struct {
	db *pgxpool.Pool
}

func NewAttendanceRepo(db *pgxpool.Pool) repoi.AttendanceI {
	return &AttendanceRepo{db: db}
}

func (a *AttendanceRepo) CreateAttendance(ctx context.Context, req *models.Attendance) error {
	query := `
		INSERT INTO
			attendances(
				id,
				student_id,
				group_id,
				date,
				status	
			)VALUES($1, $2, $3, $4, $5)
	`
	_, err := a.db.Exec(ctx, query,
		req.ID,
		req.StudentID,
		req.GroupID,
		req.Date,
		req.Status,
	)
	if err != nil {
		log.Println("Error on Create attendance:", err)
		return err
	}
	return nil
}
func (a *AttendanceRepo) GetListAttendance(ctx context.Context, req *models.GetListReq) (*models.GetListAttendance, error) {
	limit := req.Limit
	page := req.Page
	offset := (page - 1) * limit

	query := `
		SELECT
			id,
			student_id,
			group_id,
			date,
			status
		FROM
			attendances
		LIMIT $1
		OFFSET $2
	`
	rows, err := a.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error on GetListAttendance:", err)
		return nil, err
	}
	defer rows.Close()

	var listAttendance []models.Attendance
	for rows.Next() {
		var date models.Attendance
		if err := rows.Scan(
			&date.ID,
			&date.StudentID,
			&date.GroupID,
			&date.Date,
			&date.Status,
		); err != nil {
			log.Println("Error on scan date for getListAttendance:", err)
			return nil, err
		}
		listAttendance = append(listAttendance, date)
	}
	if err := rows.Err(); err != nil {
		log.Println("error on GetListAttendance:", err)
		return nil, err
	}

	return &models.GetListAttendance{
		List:  listAttendance,
		Count: len(listAttendance),
	}, nil
}
func (a *AttendanceRepo) GetAttendance(ctx context.Context, id string) (*models.Attendance, error) {
	var attendance models.Attendance
	query := `
		SELECT
			id,
			student_id,
			group_id,
			date,
			status
		FROM
			attendances
		WHERE id =$1
	`
	if err := a.db.QueryRow(ctx, query, id).Scan(
		&attendance.ID,
		&attendance.StudentID,
		&attendance.GroupID,
		&attendance.Date,
		&attendance.Status,
	); err != nil {
		log.Println("Error on GetAttendance:", err)
		return nil, err
	}

	return &attendance, nil
}
func (a *AttendanceRepo) UpdateAttendance(ctx context.Context, req *models.UpdateAttendance, id string) (*models.Attendance, error) {
	oldDate, err := a.GetAttendance(ctx, id)
	if err != nil {
		log.Println("Error on Get Old date for update attendance:", err)
		return nil, err
	}
	if req.StudentID == "" {
		req.StudentID = oldDate.StudentID
	}
	if req.GroupID == "" {
		req.GroupID = oldDate.GroupID
	}
	if req.Date.IsZero() {
		req.Date = oldDate.Date
	}

	if req.Status == "" {
		req.Status = oldDate.Status
	}

	query := `
		UPDATE 
			attendances
		SET
			student_id = $1,
			group_id = $2,	
			date = $3,
			status = $4
		WHERE id = $5
		RETURNING id, student_id, group_id, date, status
	`
	var newAttendance models.Attendance
	if err := a.db.QueryRow(ctx, query, id,
		req.StudentID,
		req.GroupID,
		req.Date,
		req.Status,
	).Scan(
		&newAttendance.ID,
		&newAttendance.StudentID,
		&newAttendance.GroupID,
		&newAttendance.Date,
		&newAttendance.Status,
	); err != nil {
		log.Println("Error on Update Attendance:", err)
		return nil, err
	}

	return &newAttendance, nil
}
func (a *AttendanceRepo) DeleteAttendance(ctx context.Context, id string) error {
	query := `
		DELET FROM
			attendances
		WHERE id = $1
	`
	_, err := a.db.Exec(ctx,query, id)
	if err != nil {
		log.Println("Error on Delete attendance:",err)
		return err
	}
	log.Println("Delete Attendance Successfully")
	
	return nil
}
