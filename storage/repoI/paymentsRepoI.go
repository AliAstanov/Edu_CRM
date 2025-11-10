package repoi

import (
	"context"

	"github.com/AliAstanov/Edu_CRM/models"
)

type PaymentsRepoI interface {
	CreatePayments(ctx context.Context, req *models.Payment) error
	GetListPayments(ctx context.Context, req *models.GetListReq) (*models.GetPayments, error)
	GetPayments(ctx context.Context, id string) (*models.Payment, error)
	UpdatePayments(ctx context.Context, req *models.UpdatePaymentReq, id string) (*models.Payment, error)
	DeletePayments(ctx context.Context, id string) error
}
