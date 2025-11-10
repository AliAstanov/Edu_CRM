package postgres

import (
	"context"
	"log"

	"github.com/AliAstanov/Edu_CRM/models"
	repoi "github.com/AliAstanov/Edu_CRM/storage/repoI"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PaymentsRepo struct {
	db *pgxpool.Pool
}

func NewPaymentsRepo(db *pgxpool.Pool) repoi.PaymentsRepoI {
	return &PaymentsRepo{db: db}
}

func (p *PaymentsRepo) CreatePayments(ctx context.Context, req *models.Payment) error {
	query := `
		INSERT INTO 
			payments(
				id,
				student_id,
				amount,
				month,
				status,
				paid_at
			)VALUES(
				$1, $2, $3, $4, $5, $6
			)`

	if _, err := p.db.Exec(ctx, query,
		req.ID,
		req.StudentID,
		req.Amount,
		req.Month,
		req.Status,
		req.PaidAt,
	); err != nil {
		log.Println("Error on CreatePayments:", err)
		return err
	}

	return nil
}
func (p *PaymentsRepo) GetListPayments(ctx context.Context, req *models.GetListReq) (*models.GetPayments, error) {
	limit := req.Limit
	page := req.Page
	offset := (page - 1) * limit

	query := `
		SELECT 
			id,
			student_id,
			amount,
			month,
			status,
			paid_at
		FROM
			payments
		LIMIT $1
		OFFSET $2`

	rows, err := p.db.Query(ctx, query, limit, offset)
	if err != nil {
		log.Println("Error on get rows for GetListPayments:", err)
		return nil, err
	}
	defer rows.Close()

	var payments []models.Payment
	for rows.Next() {
		var data models.Payment
		if err := rows.Scan(
			&data.ID,
			&data.StudentID,
			&data.Amount,
			&data.Month,
			&data.Status,
			&data.PaidAt,
		); err != nil {
			log.Println("Error on scan for GetListPayments:", err)
			return nil, err
		}
		payments = append(payments, data)

	}

	return &models.GetPayments{
		Payments: payments,
		Count:    len(payments),
	}, nil
}
func (p *PaymentsRepo) GetPayments(ctx context.Context, id string) (*models.Payment, error) {
	query := `
		SELECT 
			id,
			student_id,
			amount,
			month,
			status,
			paid_at
		FROM
			payments
		WHERE 
			id = $1 
	`
	var data models.Payment
	if err := p.db.QueryRow(ctx, query, id).Scan(
		&data.ID,
		&data.StudentID,
		&data.Amount,
		&data.Month,
		&data.Status,
		&data.PaidAt,
	); err != nil {
		log.Println("Error on scan for GetPayment:", err)
		return nil, err
	}

	return &data, nil
}
func (p *PaymentsRepo) UpdatePayments(ctx context.Context, req *models.UpdatePaymentReq, id string) (*models.Payment, error) {
	query := `
		UPDATE
			payments
		SET
			amount = $1,
			month = $2,
			status = $3
		WHERE 
			id = $4
		RETURNING id, student_id, amount, month, status, paid_at
	`
	oldData, err := p.GetPayments(ctx,id)
	if err != nil {
		log.Println("Error on Get oldData for update payments:",err)
		return nil,err
	}

	if req.Amount == nil{
		req.Amount = &oldData.Amount
	}
	if req.Month == nil{
		req.Month = &oldData.Month
	}
	if req.Status == nil{
		req.Status = &oldData.Status
	}

	var data models.Payment
	if err := p.db.QueryRow(ctx, query,
		req.Amount, 
		req.Month, 
		req.Status,).Scan(
			&data.ID,
			&data.StudentID,
			&data.Amount,
			&data.Month,
			&data.Status,
			&data.PaidAt,
			); err != nil {
				log.Println("Error on scan for UpdatePayment:", err)
				return nil, err
			}

	
	return &data, nil
}
func (p *PaymentsRepo) DeletePayments(ctx context.Context, id string) error {
	query := `
		DELETE FROM
			payments
		WHERE
			id = $1
	`
	_, err := p.db.Exec(ctx, query, id)
	if err != nil {
		log.Println("Error on DeletePayment:", err)
		return err
	}
	log.Println("Delete Payment successfully")
	
	return nil
}
