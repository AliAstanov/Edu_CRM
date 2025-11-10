package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	helpers "github.com/AliAstanov/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

type SubjectRepo struct {
	db *pgxpool.Pool
}

func NewSubject(db *pgxpool.Pool) repoi.SubjectRepoI {
	return &SubjectRepo{db: db}
}

func (s *SubjectRepo) CreateSubject(ctx context.Context, req *models.Subject) error {
	query := `
		INSERT INTO
			subjects(
				id,
				name
			)VALUES(
			$1,$2)
	`
	_, err := s.db.Exec(ctx, query, req.ID, req.Name)
	if err != nil {
		log.Println("Error on CreateSubject:", err)
		return err
	}

	log.Println("create subject Soccessfully")

	return nil
}
func (s *SubjectRepo) GetListSubject(ctx context.Context, req *models.GetListReq) (*models.GetSubjects, error) {

	limit := req.Limit
	page := req.Page
	offset := helpers.Offset(limit, page)

	query := `
		SELECT 
			id,
			name
		FROM
			subject
		LIMIT $1 
		OFFSET $2
	`

	rows, err := s.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error on GetListSubject:", err)
		return nil, err
	}
	defer rows.Close()

	var subjects []models.Subject

	for rows.Next() {
		var subject models.Subject
		if err = rows.Scan(
			&subject.ID,
			&subject.Name,
		); err != nil {
			log.Println("Error on rows.Next for GetListSubject:", err)
			return nil, err
		}
		subjects = append(subjects, subject)

	}

	return &models.GetSubjects{
		Subjects: subjects,
		Count:    len(subjects),
	}, nil
}
func (s *SubjectRepo) GetSubject(ctx context.Context, id string) (*models.Subject, error) {
	query := `
		SELECT 
			id,
			name
		FROM
			subjects
	`
	var subject models.Subject
	if err := s.db.QueryRow(ctx, query, id).Scan(
		&subject.ID,
		&subject.Name,
	); err != nil {
		log.Println("Error on scan data for GetSubject:", err)
		return nil, err
	}
	return &subject, nil
}
func (s *SubjectRepo) UpdateSubject(ctx context.Context, req *models.UpdateSubjectReq, id string) (*models.Subject, error) {
	query := `
		UPDATE
			subjects
		SET
			id = $1
			name = $2
		WHERE
		id = $3
		RETURNING id, name
	`
	var subject models.Subject
	if err := s.db.QueryRow(ctx, query, id).Scan(&subject.ID, &subject.Name); err != nil {
		log.Println("Error on scan data for UpdateSubject:", err)
		return nil, err
	}

	return nil, nil
}
func (s *SubjectRepo) DeleteSubject(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			subjects
		WHERE
			id = $1
	`
	if _, err := s.db.Exec(ctx,query,id); err != nil {
		log.Println("Error on Delete subject:",err)
		return  err
	}
		log.Println("Delete subject soccessfully")
	return nil
}
