package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type GroupRepo struct {
	db *pgxpool.Pool
}

func NewGroupRepo(db *pgxpool.Pool) repoi.GroupI {
	return &GroupRepo{db: db}
}
func (g *GroupRepo) CreateGroup(ctx context.Context, req *models.Group) error {
	query := `
		INSERT INTO
			groups(
			id,
			name,
			created_at
			)VALUES($1,$2,$3)
	`
	_, err := g.db.Exec(ctx, query,
		req.ID,
		req.Name,
		req.CreatedAt)
	if err != nil {
		log.Println("Error create group", err)
		return err
	}
	return nil
}
func (g *GroupRepo) GetListGroup(ctx context.Context, req *models.GetListReq) (*models.GetGroups, error) {
	query := `
		SELECT 
			id,
			name,
			created_at
		FROM
			groups
		LIMIT $1
		OFFSET $2
	`
	limit := req.Limit
	page := req.Page
	offset := (page - 1) * limit

	rows, err := g.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error on get rows for GetListGroup:", err)
		return nil, err
	}

	defer rows.Close()

	var groups []models.Group

	for rows.Next() {
		var group models.Group
		if err := rows.Scan(
			&group.ID,
			&group.Name,
			&group.CreatedAt,
		); err != nil {
			log.Println("Error on scan data for GetListGroup:", err)
			return nil, err
		}
		groups = append(groups, group)
	}
	if err := rows.Err(); err != nil {
		log.Println("Error on GetListGroup:", err)
		return nil, err
	}
	return &models.GetGroups{
		Groups: groups,
		Count:  len(groups),
	}, nil
}
func (g *GroupRepo) GetGroup(ctx context.Context, id string) (*models.Group, error) {
	query:= `
		SELECT
			id,
			name,
			created_at
		FROM
			groups
		WHERE
		id = $1
	`
	row := g.db.QueryRow(ctx,query,id)
	var group models.Group
	if err := row.Scan(
		&group.ID,
		&group.Name,
		&group.CreatedAt,
	); err != nil {
		log.Println("Error on scan data for GetGroup:", err)
		return  nil, err
	}

	return &group, nil
}
func (g *GroupRepo) UpdateGroup(ctx context.Context, req *models.UpdateGroupReq, id string) (*models.Group, error) {
	query := `
		UPDATE 
			groups
		SET
			name = $1
		WHERE id = $2
		RETURNING id, name, created_at
	`
	var group models.Group
	if err := g.db.QueryRow(ctx,query,id).Scan(
		&group.ID,
		&group.Name,
		&group.CreatedAt,	
	); err != nil{
		log.Println("Error on get QueryRov for UpdateGroup:",err)
		return nil,err
	}

	return &group, nil
}
func (g *GroupRepo) DeleteGroup(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			groups
		WHERE id = $1
	`
	_, err := g.db.Exec(ctx,query,id)
	if err != nil {
		log.Println("Error on delete group:", err)
		return err
	}
	
	return nil
}
