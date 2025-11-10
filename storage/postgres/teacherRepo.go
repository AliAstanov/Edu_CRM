package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type TeacherRepo struct {
	db *pgxpool.Pool
}

func NewTeacherRepo(db *pgxpool.Pool) repoi.TeacherRepoI {
	return &TeacherRepo{db: db}
}

func (t *TeacherRepo) CreateTeacher(ctx context.Context, req *models.Teacher) error {
	
	query := `
		INSERT INTO
			teachers(
				id,
				full_name,
				phone,
				created_at
			)VALUES(
			$1,$2,$3,$4)
	`
	_, err := t.db.Exec(ctx, query,
		req.ID,
		req.FullName,
		req.Phone,
		req.CreatedAt,
	)
	if err != nil {
		log.Println("Erro on Create teacher:", err)
		return err
	}
	log.Printf("Create Student: %s soccessfully", req)

	return nil
}
func (t *TeacherRepo) GetListTeacher(ctx context.Context, req *models.GetListReq) (*models.GetTeachers, error) {
	query := `
		SELECT 
			id,
			full_name,
			phone,
			created_at
		FROM
			teachers
		LIMIT $1
		OFFSET $2
	`
	limit := req.Limit
	page := req.Page
	offset := (page - 1) * limit

	rows, err := t.db.Query(ctx, query, offset)
	if err != nil {
		log.Println("Error on get rows for GetListTeacher:", err)
		return nil, err
	}
	defer rows.Close()

	var teachers []models.Teacher
	for rows.Next() {
		var teacher models.Teacher
		if err = rows.Scan(
			&teacher.ID,
			&teacher.FullName,
			&teacher.Phone,
			&teacher.CreatedAt,
		); err != nil {
			log.Println("Error on scan GetListTeacher:", err)
			return nil, err
		}
		teachers = append(teachers, teacher)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on GetListTeacher:", err)
		return nil, err
	}
	return &models.GetTeachers{
		Teachers: teachers,
		Count:    len(teachers),
	}, nil
}
func (t *TeacherRepo) GetTeacher(ctx context.Context, id string) (*models.Teacher, error) {
	query := `
		SELECT 
			id,
			full_name,
			phone,
			created_at
		FROM
			teachers
		WHERE 
			id = $1	
	`
	var teacher models.Teacher
	if err := t.db.QueryRow(ctx, query, id).Scan(
		&teacher.ID,
		&teacher.FullName,
		&teacher.Phone,
		&teacher.CreatedAt,
	); err != nil {
		log.Println("Error on get teacher for GetTeacher:", err)
		return nil, err
	}
	return &teacher, nil
}
func (t *TeacherRepo) UpdateTeacher(ctx context.Context, req *models.UpdateTeacherReq, id string) (*models.Teacher, error) {
	var oldTeacher models.Teacher

	query := `
		UPDATE
			teachers
		SET
			full_name = $1
			phone = $2
		WHERE
			id = $3 
		RETURNING id, full_name, phone, created_at
	`
	if req.FullName == "" {
		req.FullName = oldTeacher.FullName
	}
	if req.Phone == "" {
		req.Phone = oldTeacher.Phone
	}
	var teacher models.Teacher
	err := t.db.QueryRow(ctx, query,
		req.FullName,
		req.Phone, id).Scan(
		&teacher.CreatedAt,
		&teacher.FullName,
		&teacher.Phone,
		&teacher.ID,
	)
	if err != nil {
		log.Println("Error on update teacher for UpdateTeacher:", err)
		return nil, err
	}

	return &teacher, nil
}
func (t *TeacherRepo) DeleteTeacher(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			teachers
		WHERE
			id = $1
	`
	_, err := t.db.Exec(ctx,query,id)
	if err != nil {
		log.Println("Error on delete teacher for DeleteTeacher:", err)
		return err
	}
	return nil
}
