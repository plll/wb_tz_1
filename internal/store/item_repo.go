package store

import "github.com/jackc/pgx/v4/pgxpool"

type ItemsRepository struct {
	db *pgxpool.Pool
}

func NewItemsRepository(db *pgxpool.Pool) Items {
	return &ItemsRepository{
		db: db,
	}
}

func (i *ItemsRepository) AddItem() string {
	query := `INSERT INTO item (chrt_id, track_number, price, rid, 
	                                        name, sale, size, total_price, 
	                                        nm_id, brand, status) 
                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)
                        RETURNING chrt_id`
	return query
}

func (i *ItemsRepository) GetAllItemsByOrderId() string {
	query := `SELECT chrt_id, track_number, price, rid, name, 
                             sale, size, total_price, nm_id, brand,status
                       FROM item
                       WHERE item.chrt_id in (SELECT chrt_id FROM item_order WHERE order_uid = $1)`
	return query
}
