package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type AiLogsI interface {
	CreateAILog(ctx context.Context, req *models.AILog) error
	GetListAILog(ctx context.Context, req *models.GetListReq) (*models.GetAILogs, error)
	GetAILog(ctx context.Context, id string) (*models.AILog, error)
	UpdateAILog(ctx context.Context, req *models.UpdateAILogReq, id string) (*models.AILog, error)
	DeleteAILog(ctx context.Context, id string) error
}
