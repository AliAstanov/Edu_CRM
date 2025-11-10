package models

import (
	"time"

	"github.com/google/uuid"
	"encoding/json"
)

// To‘liq AI log yozuvi (bazadan o‘qish uchun)
type AILog struct {
	ID         uuid.UUID       `json:"id"`
	UserID     *uuid.UUID      `json:"user_id,omitempty"` // optional
	InputType  string          `json:"input_type"`
	InputData  json.RawMessage `json:"input_data"`  // JSONB tipini ifodalaydi
	Response   string          `json:"response"`
	ModelUsed  string          `json:"model_used,omitempty"`
	CreatedAt  time.Time       `json:"created_at"`
}

// Ko‘p loglar bilan javob
type GetAILogs struct {
	Logs  []AILog `json:"logs"`
	Count int     `json:"count"`
}

// Yangi AI log qo‘shish uchun
type CreateAILog struct {
	UserID    *uuid.UUID     `json:"user_id"` // optional
	InputType string         `json:"input_type" binding:"required"`
	InputData json.RawMessage `json:"input_data" binding:"required"` // frontend JSON yuboradi
	Response  string         `json:"response" binding:"required"`
	ModelUsed string         `json:"model_used"`
}

type UpdateAILogReq struct {
	InputType *string          `json:"input_type"`  // optional
	InputData *json.RawMessage `json:"input_data"`  // optional
	Response  *string          `json:"response"`    // optional
	ModelUsed *string          `json:"model_used"`  // optional
}
