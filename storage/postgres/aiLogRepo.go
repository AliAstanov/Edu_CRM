package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type AiLogRepo struct {
	db *pgxpool.Pool
}

func NewAiLogRepo(db *pgxpool.Pool) repoi.AiLogsI {
	return &AiLogRepo{db: db}
}

func (a *AiLogRepo) CreateAILog(ctx context.Context, req *models.AILog) error {
	query := `
		INSERT INTO
			ai_logs(
				id,
				user_id,
				input_type,
				input_data,
				response,
				model_used,
				created_at
			)VALUES(
				$1, $2, $3, $4, $5, $6, $7
			)`
	if _, err := a.db.Exec(ctx, query,
		req.ID,
		req.UserID,
		req.InputType,
		req.InputData,
		req.Response,
		req.ModelUsed,
		req.CreatedAt,
	); err != nil {
		log.Println("Error on Create ai_log:", err)
		return err
	}

	return nil
}
func (a *AiLogRepo) GetListAILog(ctx context.Context, req *models.GetListReq) (*models.GetAILogs, error) {
	limit := req.Limit
	offset := (req.Page - 1) * limit

	query := `
		SELECT 
			id,
			user_id,
			input_type,
			input_data,
			response,
			model_used,
			created_at
		FROM
			ai_logs
		LIMIT $1
		OFFSET $2
	`
	rows, err := a.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Printf("GetListAiLog: failed to execute query: %v\n", err)
		return nil, err
	}
	defer rows.Close()

	var aiLogs []models.AILog
	for rows.Next() {
		var data models.AILog
		if err := rows.Scan(
			&data.ID,
			&data.UserID,
			&data.InputType,
			&data.InputData,
			&data.Response,
			&data.ModelUsed,
			&data.CreatedAt,
		); err != nil {
			log.Println("Error on GetListAiLog:", err)
			return nil, err
		}
		aiLogs = append(aiLogs, data)

	}

	return &models.GetAILogs{
		Logs:  aiLogs,
		Count: len(aiLogs),
	}, nil
}
func (a *AiLogRepo) GetAILog(ctx context.Context, id string) (*models.AILog, error) {
	query := `
		SELECT
			id,
			user_id,
			input_type,
			input_data,
			response,
			model_used,
			created_at
		FROM
			ai_logs
		WHERE 
			id = $1			
	`
	var req models.AILog
	if err := a.db.QueryRow(ctx, query, id).Scan(
		&req.ID,
		&req.UserID,
		&req.InputType,
		&req.InputData,
		&req.Response,
		&req.ModelUsed,
		&req.CreatedAt,
	); err != nil {
		log.Println("Error on GetAILog:", err)
		return nil, err
	}

	return &req, nil
}
func (a *AiLogRepo) UpdateAILog(ctx context.Context, req *models.UpdateAILogReq, id string) (*models.AILog, error) {
	query := `
		UPDATE
			ai_log
		SET
			input_type = $1,
			input_data = $2,
			response = $3,
			model_used = $4
		WHERE 
			id = $5
		RETURNING id, user_id, input_type, input_data, response, model_used, created_at
	`
	oldData, err := a.GetAILog(ctx, id)
	if err != nil {
		log.Println("UpdateAILog: failed to execute oldData:", err)
		return nil, err
	}

	if req.InputData == nil {
		req.InputData = &oldData.InputData
	}
	if req.InputType == nil {
		req.InputType = &oldData.InputType
	}
	if req.Response == nil {
		req.Response = &oldData.Response
	}
	if req.ModelUsed == nil {
		req.ModelUsed = &oldData.ModelUsed
	}

	var data models.AILog
	if err := a.db.QueryRow(ctx, query, id).Scan(
		&data.ID,
		&data.UserID,
		&data.InputType,
		&data.InputData,
		&data.Response,
		&data.ModelUsed,
		&data.CreatedAt,
	); err != nil {
		log.Println("Error to scan data on UpdateAiLog:", err)
		return nil, err
	}

	return &data, nil
}
func (a *AiLogRepo) DeleteAILog(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			ai_logs
		WHERE
		 id = $1
	`
	if _, err := a.db.Exec(ctx, query, id); err != nil {
		log.Println("Error on Delete ai_log:", err)
		return err
	}

	return nil
}
