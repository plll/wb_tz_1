package store

import "github.com/jackc/pgx/v4/pgxpool"

type OrderItemsRepository struct {
	db *pgxpool.Pool
}

func NewOrderItemsRepository(db *pgxpool.Pool) OrdersItems {
	return &OrderItemsRepository{
		db: db,
	}
}

func (i *OrderItemsRepository) AddOrderItem() string {
	query := `INSERT INTO item_order (chrt_id, order_uid)
                             VALUES ($1, $2)`
	return query
}
