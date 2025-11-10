package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	halpers "github.com/AliAstanov/helper"
	"github.com/jackc/pgx/v5/pgxpool"
)

type groupSubjectTeacherRepo struct {
	db *pgxpool.Pool
}

func NewGroupSubjectTeacher(db *pgxpool.Pool) repoi.GroupSubjectTeacherI {
	return &groupSubjectTeacherRepo{db: db}
}

func (g *groupSubjectTeacherRepo) CreateGroupSubjectTeacher(ctx context.Context, req *models.GroupSubjectTeacher) error {
	query := `
		INSERT INTO
			group_subject_teacher(
				id,
				group_id,
				subject_id,
				teacher_id,
				start_date,
				end_date	
			)VALUES(
			$1,$2,$3,$4,$5,$6)
	`
	if _, err := g.db.Exec(ctx, query,
		req.ID,
		req.GroupID,
		req.SubjectID,
		req.TeacherID,
		req.StartDate,
		req.EndDate,
	); err != nil {
		log.Println("Error on CreateGroupSubjectTeacher:", err)
		return err
	}

	return nil
}
func (g *groupSubjectTeacherRepo) GetListGroupSubjectTeacher(ctx context.Context, req *models.GetListReq) (*models.GetGroupSubjectTeachers, error) {
	limit := req.Limit
	page := req.Page

	offset := halpers.Offset(limit, page)

	query := `
		SELECT
			id,
			group_id,
			subject_id,
			teacher_id,
			start_date,
			end_date
		FROM
			group_subject_teachers
		LIMIT  $1
		OFFSET $2	
	`
	rows, err := g.db.Query(ctx, query, req.Limit, offset)
	if err != nil {
		log.Println("error on GetListGroupSubjectTeacher:", err)
		return nil, err
	}
	defer rows.Close()

	var groupSubjectTeachers []models.GroupSubjectTeacher
	for rows.Next() {
		var groupSubjectTeacher models.GroupSubjectTeacher
		if err := rows.Scan(
			&groupSubjectTeacher.ID,
			&groupSubjectTeacher.GroupID,
			&groupSubjectTeacher.SubjectID,
			&groupSubjectTeacher.TeacherID,
			&groupSubjectTeacher.StartDate,
			&groupSubjectTeacher.EndDate,
		); err != nil {
			log.Println("Error in scan for GetListGroupSubjectTeacher:", err)
			return nil, err
		}
		groupSubjectTeachers = append(groupSubjectTeachers, groupSubjectTeacher)
	}

	return &models.GetGroupSubjectTeachers{
		List:  groupSubjectTeachers,
		Count: len(groupSubjectTeachers),
	}, nil
}
func (g *groupSubjectTeacherRepo) GetGroupSubjectTeacher(ctx context.Context, id string) (*models.GroupSubjectTeacher, error) {
	var req models.GroupSubjectTeacher
	query := `
		SELECT 
			id,
			group_id,
			subject_id,
			teacher_id,
			start_date,
			end_date
		FROM
			group_subject_teachers
		WHERE
			id = $1
	`
	if err := g.db.QueryRow(ctx, query, id).Scan(
		&req.ID,
		&req.GroupID,
		&req.SubjectID,
		&req.TeacherID,
		&req.StartDate,
		&req.EndDate,
	); err != nil {
		log.Println("Error on GetGroupSubjectTeacher:", err)
		return nil, err
	}
	return &req, nil

}
func (g *groupSubjectTeacherRepo) UpdateGroupSubjectTeacher(ctx context.Context, req *models.UpdateGroupSubjectTeacher, id string) (*models.GroupSubjectTeacher, error) {
	oldData, err := g.GetGroupSubjectTeacher(ctx, id)
	if err != nil {
		log.Println("error on GetOldGroupSubjectTeacher for update:", err)
		return nil, err
	}

	// To'ldirilmagan joylarni eski qiymat bilan to'ldirish
	if req.GroupID == nil {
		req.GroupID = &oldData.GroupID
	}
	if req.SubjectID == nil {
		req.SubjectID = &oldData.SubjectID
	}
	if req.TeacherID == nil {
		req.TeacherID = oldData.TeacherID
	}
	if req.StartDate == nil {
		req.StartDate = oldData.StartDate
	}
	if req.EndDate == nil {
		req.EndDate = oldData.EndDate
	}

	query := `
		UPDATE 
			group_subject_teacher
		SET
			group_id = $1,
			subject_id = $2,
			teacher_id = $3,
			start_date = $4,
			end_date = $5
		WHERE
			id = $6
		RETURNING id, group_id, subject_id, teacher_id, start_date, end_date
	`

	var updated models.GroupSubjectTeacher
	err = g.db.QueryRow(ctx, query,
		req.GroupID,
		req.SubjectID,
		req.TeacherID,
		req.StartDate,
		req.EndDate,
		id,
	).Scan(
		&updated.ID,
		&updated.GroupID,
		&updated.SubjectID,
		&updated.TeacherID,
		&updated.StartDate,
		&updated.EndDate,
	)
	if err != nil {
		log.Println("Error on Returning data for update GroupSubjectTeacher:", err)
		return nil, err
	}

	return &updated, nil
}
func (g *groupSubjectTeacherRepo) DeleteGroupSubjectTeacher(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			group_suvject_teachers
		WHERE 
			id = $1
	`
	_, err := g.db.Exec(ctx,query, id)
	if err != nil{
		log.Println("Error on delete groupSubjectTeacher:",err)
		return err
	}

	return nil
}
