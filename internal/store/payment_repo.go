package store

import "github.com/jackc/pgx/v4/pgxpool"

type PaymentsRepository struct {
	db *pgxpool.Pool
}

func NewPaymentsRepository(db *pgxpool.Pool) Payments {
	return &PaymentsRepository{
		db: db,
	}
}

func (p *PaymentsRepository) AddPayment() string {
	query := `INSERT INTO payment_info (transaction, request_id, currency, provider, 
	                                                amount, payment_dt, bank, delivery_cost, 
	                                                goods_total, custom_fee) 
                           VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
                           RETURNING id`
	return query
}

func (p *PaymentsRepository) GetPaymentByOrderId() string {
	query := `SELECT transaction, request_id, currency, 
                               amount, payment_dt, bank, delivery_cost, 
                               goods_total, custom_fee, provider
                        FROM payment_info
                        WHERE payment_info.id = (SELECT payment FROM orders WHERE orders.order_uid = $1)`
	return query
}
