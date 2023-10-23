package store

import (
	"github.com/jackc/pgx/v4"
)

type OrderItemsRepository struct {
	db *pgx.Conn
}

func NewOrderItemsRepository(db *pgx.Conn) OrdersItems {
	return &OrderItemsRepository{
		db: db,
	}
}

func (i *OrderItemsRepository) AddOrderItem() string {
	query := `(INSERT INTO item_order (chrt_id, order_uid)
                             VALUES ($1, $2))`
	return query
}
