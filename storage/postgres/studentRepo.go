package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	halpers "github.com/AliAstanov/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentRepo struct {
	db *pgxpool.Pool
}

func NewStudentRepo(db *pgxpool.Pool) repoi.StudentRepoI {
	return &StudentRepo{db: db}
}

func (s *StudentRepo) CreateStudent(ctx context.Context, req *models.Student) error {
	query := `
		INSERT INTO
			students(
				student_id,
				full_name,
				phone_number,
				group_id,
				created_at
			)VALUES(
			$1,$2,$3,$4,$5)
	`
	_, err := s.db.Exec(ctx, query,
		req.StudentId,
		req.FullName,
		req.PhoneNumber,
		req.GroupID,
		req.CreatedAt,
	)
	if err != nil {
		log.Println("error in CreateStudent:", err)
		return err
	}
	log.Printf("Create Student: %s soccessfully", req)

	return nil
}

func (s *StudentRepo) GetListStudent(ctx context.Context, req *models.GetListReq) (*models.GetStudens, error) {
	limit := req.Limit
	page := req.Page

	offset := halpers.Offset(limit, page)

	query := `
		SELECT 
			student_id,
			full_name,
			phone_number,
			group_id,
			created_at
		FROM
			students
		LIMIT $1
		OFFSET $2
	`
	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("error in GetStudents to get rows:", err)
		return nil, err
	}
	defer rows.Close()

	var students []models.Student
	for rows.Next() {
		var student models.Student
		if err := rows.Scan(
			&student.StudentId,
			&student.FullName,
			&student.PhoneNumber,
			&student.GroupID,
			&student.CreatedAt,
		); err != nil {
			log.Println("error in GetStudents to scan rows:", err)
			return nil, err
		}
		students = append(students, student)
	}
	if err := rows.Err(); err != nil {
		log.Println("error in GetStudents to get rows:", err)
		return nil, err
	}
	return &models.GetStudens{
		Students: students,
		Count:    len(students),
	}, nil
}

func (s *StudentRepo) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	query := `
		SELECT 
			student_id,
			full_name,
			phone_number,
			group_id,
			created_at
		FROM
			students
		WHERE
			student_id = $1
	`
	row := s.db.QueryRow(ctx, query, id)
	var student models.Student
	if err := row.Scan(
		&student.StudentId,
		&student.FullName,
		&student.PhoneNumber,
		&student.GroupID,
		&student.CreatedAt,
	); err != nil {
		log.Println("error in get row to GetStudent")
		return nil, err
	}
	return &student, nil
}

func (s *StudentRepo) UpdateStudent(ctx context.Context, req *models.UpdateStudentReq, id string) (*models.Student, error) {
	oldStudentDate, err := s.GetStudent(ctx, id)
	if err != nil {
		log.Println("Error on GetStudentDate for update student:", err)
		return nil, err
	}

	if req.FullName == "" {
		req.FullName = oldStudentDate.FullName
	}
	if req.GroupID == "" {
		req.GroupID = oldStudentDate.GroupID
	}
	if req.PhoneNumber == "" {
		req.PhoneNumber = oldStudentDate.PhoneNumber
	}

	query := `
		UPDATE 
			students
		SET
			full_name    = $1,
			phone_number = $2,
			group_id     = $3
		WHERE
			student_id = $4
		RETURNING student_id, full_name, phone_number, group_id, created_at 
	`
	row := s.db.QueryRow(ctx, query,
		req.FullName,
		req.PhoneNumber,
		req.GroupID,
		id)
	var updatedData models.Student
	if err := row.Scan(
		&updatedData.StudentId,
		&updatedData.FullName,
		&updatedData.PhoneNumber,
		&updatedData.GroupID,
		&updatedData.CreatedAt,
	); err != nil {
		log.Println("error in get row to UpdateStudent")
		return nil, err
	}

	return &updatedData, nil
}

func (s *StudentRepo) DeleteStudent(ctx context.Context, id string) error {

	query := `
		DELETE FROM 
			students
		WHERE 
			student_id = $1
	`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error on Delete Students:", err)
		return err
	}

	log.Println("Delete student soccessfully")

	return nil
}
