package store

import "github.com/jackc/pgx/v4"

type DeliveriesRepository struct {
	db *pgx.Conn
}

func NewDeliveriesRepository(db *pgx.Conn) Deliveries {
	return &DeliveriesRepository{
		db: db,
	}
}

func (d *DeliveriesRepository) AddDelivery() string {
	query := `(INSERT INTO delivery_info (name, phone, zip, city, 
	                                                  address, region, email)
                            VALUES ($1, $2, $3, $4, $5, $6, $7)
                            RETURNING id)`
	return query
}

func (d *DeliveriesRepository) GetDeliveryByOrderId() string {
	query := `(SELECT name, phone, zip, city, 
                                address, region, email
                          FROM delivery_info
                          WHERE delivery_info.id = (SELECT delivery FROM orders WHERE orders.order_uid = $1))`
	return query
}
