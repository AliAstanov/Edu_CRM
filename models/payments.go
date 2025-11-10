package models

import (
	"time"

	"github.com/google/uuid"
)

// To‘liq payment yozuvi
type Payment struct {
	ID        uuid.UUID `json:"id"`
	StudentID uuid.UUID `json:"student_id"`
	Amount    int       `json:"amount"`
	Month     string    `json:"month"`   // format: '2025-07'
	Status    string    `json:"status"`  // 'paid', 'unpaid', 'pending'
	PaidAt    time.Time `json:"paid_at"`
}

// GET so‘rovlar uchun
type GetPayments struct {
	Payments []Payment `json:"payments"`
	Count    int       `json:"count"`
}

// Yangi payment yaratish uchun
type CreatePayment struct {
	StudentID uuid.UUID `json:"student_id" binding:"required"`
	Amount    int       `json:"amount" binding:"required"`
	Month     string    `json:"month" binding:"required"` // '2025-07' format
	Status    string    `json:"status" binding:"omitempty,oneof=paid unpaid pending"`
}

// Paymentni yangilash uchun
type UpdatePaymentReq struct {
	Amount *int    `json:"amount"` // optional
	Month  *string `json:"month"`  // optional
	Status *string `json:"status"` // optional: 'paid', 'unpaid', 'pending'
}


