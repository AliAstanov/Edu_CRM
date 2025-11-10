package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type StudentGroupRepo struct {
	db *pgxpool.Pool
}

func NewStudentGroupRepo(db *pgxpool.Pool) repoi.StudentGroupI {
	return &StudentGroupRepo{db: db}
}

func (s *StudentGroupRepo) CreateStudentGroup(ctx context.Context, req *models.StudentGroup) error {
	query := `
		INSERT INTO 
			student_groups(
				id,
				student_id,
				group_id	
			)VALUES(
			$1, $2, $3
			)`
	_, err := s.db.Exec(ctx, query,
		req.ID,
		req.StudentID,
		req.GroupID,
	)
	if err != nil {
		log.Println("Error on CreateStudentGroup:", err)
		return err
	}

	return nil
}
func (s *StudentGroupRepo) GetListStudentGroup(ctx context.Context, req *models.GetListReq) (*models.GetListStudentGroup, error) {
	limit := req.Limit
	page := req.Page
	Offset := (page - 1) * limit

	query := `
		SELECT 
			id,
			student_id,
			group_id
		FROM
			studentGroups
		LIMIT $1 OFFSET $2
	`
	rows, err := s.db.Query(ctx, query, limit, Offset)
	if err != nil {
		log.Println("Error on GetListStudentGroup:", err)
		return nil, err
	}
	defer rows.Close()

	var list []models.StudentGroup
	for rows.Next() {
		var studentGroup models.StudentGroup
		if err := rows.Scan(
			&studentGroup.ID,
			&studentGroup.StudentID,
			&studentGroup.GroupID,
		); err != nil {
			log.Println("Error on scan studentGroup for get list StudentGrup:", err)
			return nil, err
		}
		list = append(list, studentGroup)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on GetListStudentGroup:", err)
		return nil, err
	}
	return &models.GetListStudentGroup{
		List:  list,
		Count: len(list),
	}, nil
}

func (s *StudentGroupRepo) GetStudentGroup(ctx context.Context, id string) (*models.StudentGroup, error) {
	var req models.StudentGroup
	query := `
		SELECT
			id,
			student_id,
			group_id
		FROM
		    studentGroups
		WHERE id = $1
	`
	row := s.db.QueryRow(ctx, query, id)
	if err := row.Scan(
		&req.ID,
		&req.StudentID,
		&req.GroupID,
	); err != nil {
		log.Println("error on Scan req for get studentGroup:", err)
		return nil, err
	}
	return &req, nil
}
func (s *StudentGroupRepo) UpdateStudentGroup(ctx context.Context, req *models.UpdateStudentGroup, id string) (*models.StudentGroup, error) {

	query := `
		UPDATE 
			student_groups
		SET
			student_id = $1,
			group_id = $2
		WHERE
			id = $3
		RETURNING id, student_id, group_id
	`
	olddata, err := s.GetStudentGroup(ctx, id)
	if err != nil {
		log.Println("Error on get old data for update studentGroups:", err)
		return nil, err
	}
	if req.StudentID == "" {
		req.StudentID = olddata.StudentID
	}

	if req.GroupID == "" {
		req.GroupID = olddata.GroupID
	}
	var studentGroup models.StudentGroup
	if err := s.db.QueryRow(ctx,query,id,
		req.StudentID,
		req.GroupID).Scan(
			&studentGroup.ID,
			&studentGroup.StudentID,
			&studentGroup.GroupID,
	); err != nil {
		log.Println("error on Scan req for update studentGroup:", err)
		return nil,err
	}

	return &studentGroup, nil
}
func (s *StudentGroupRepo) DeleteStudentGroup(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			studentGroups
		WHERE
			id = $1
	`
	_, err := s.db.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error on Delete studentGroup:",err)
		return err
	}
	
	return nil
}
