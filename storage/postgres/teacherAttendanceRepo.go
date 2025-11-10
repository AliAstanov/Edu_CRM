package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type teacherAttendancerepo struct {
	db *pgxpool.Pool
}

func NewTeacherAttendanceRepo(db *pgxpool.Pool) repoi.TeacherAttendanceI {
	return &teacherAttendancerepo{
		db: db,
	}
}

func (t *teacherAttendancerepo) CreateTeacherAttendance(ctx context.Context, req *models.TeacherAttendance) error {
	query := `
		INSERT INTO 
			teacher_attendances(
				id,
				teacher_id,
				date,
				status,
				note,
				is_excused
			)VALUES(
				$1, $2, $3, $4, $5, $6
			)`
	if _, err := t.db.Exec(ctx, query,
		req.ID,
		req.TeacherID,
		req.Date,
		req.Status,
		req.IsExcused,
	); err != nil {
		log.Println("Error on CreateTeacherAttendance:", err)
		return err
	}

	return nil
}
func (t *teacherAttendancerepo) GetListTeacherAttendance(ctx context.Context, req *models.GetListReq) (*models.GetTeacherAttendances, error) {
	limit := req.Limit
	page := req.Page
	offset := (page - 1) * limit

	query := `
		SELECT
			id,
			teacher_id,
			date,
			status,
			note,
			is_excused
		FROM
			teacher_attendances
		LIMIT $1
		OFFSET $2

	`
	rows, err := t.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error on GetListTeacherAttendance:", err)
		return nil, err
	}
	defer rows.Close()

	var teacherAttendances []models.TeacherAttendance
	for rows.Next() {
		var date models.TeacherAttendance
		if err := rows.Scan(
			&date.ID,
			&date.TeacherID,
			&date.Date,
			&date.Status,
			&date.Note,
			&date.IsExcused,
		); err != nil {
			log.Println("Error on GetListTeacherAttendance:", err)
			return nil, err
		}
		teacherAttendances = append(teacherAttendances, date)
	}

	return &models.GetTeacherAttendances{
		Attendances: teacherAttendances,
		Count:       len(teacherAttendances),
	}, nil
}
func (t *teacherAttendancerepo) GetTeacherAttendance(ctx context.Context, id string) (*models.TeacherAttendance, error) {
	var teacherAttendance models.TeacherAttendance

	query := `
		SELECT 
			id,
			teacher_id,
			date,
			status,
			note,
			is_excused
		FROM
			teacher_attendances
		WHERE id = $1`

	if err := t.db.QueryRow(ctx, query, id).Scan(
		&teacherAttendance.ID,
		&teacherAttendance.TeacherID,
		&teacherAttendance.Date,
		&teacherAttendance.Status,
		&teacherAttendance.Note,
		&teacherAttendance.IsExcused,
	); err != nil {
		log.Println("Error on GetTeacherAttendance:", err)
		return nil, err
	}

	return &teacherAttendance, nil
}
func (t *teacherAttendancerepo) UpdateTeacherAttendance(ctx context.Context, req *models.UpdateTeacherAttendanceReq, id string) (*models.TeacherAttendance, error) {
	oldDate, err := t.GetTeacherAttendance(ctx, id)
	if err != nil {
		log.Println("Errorn on get OldDate for updateTeacherAttendance:", err)
		return nil, err
	}
	var newDate models.TeacherAttendance

	if req.Date == "" {
		req.Date = oldDate.Date.Format("2006-01-02")
	}
	if req.Status == "" {
		req.Status = oldDate.Status
	}
	if req.Note == "" {
		req.Note = oldDate.Note
	}
	query := `
		UPDATE 
			teacher_attendances
		SET
			date = $1,
			status = $2,
			note = $3,
			is_excused = $4
		WHERE id = $5
		RETURNING id, teacher_id, date, status, note, is_excused
	`
	if err := t.db.QueryRow(ctx, query, 
			req.Date, 
			req.Status, 
			req.Note, 
			req.IsExcused,
		).Scan(
			&newDate.ID,
			&newDate.TeacherID,
			&newDate.Date,
			&newDate.Status,
			&newDate.Note,
			&newDate.IsExcused,
		); err != nil {
			log.Println("Error on Update teacher attendance:",err)
			return nil, err
		}

	return &newDate, nil
}
func (t *teacherAttendancerepo) DeleteTeacherAttendance(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			teacher_atendances
		WHERE id = $1
	`
	_, err := t.db.Exec(ctx,query,id)
	if err != nil {
		log.Println("Error on Delete TeacherAttendance:",err)
		return err
	}
	return nil
}
