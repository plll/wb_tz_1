package store

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/plll/wb_tz_1/internal/datastruct"
)

type OrdersRepository struct {
	db *pgx.Conn
}

func NewOrdersRepository(db *pgx.Conn) Orders {
	return &OrdersRepository{
		db: db,
	}
}

func (o *OrdersRepository) GetOrderById() string {
	query := `(SELECT order_uid, entry, locale, internal_signature, customer_id,
		       delivery_service, shardkey, sm_id, date_created, oof_shard, track_number
	           FROM orders
	           WHERE order_uid = $1)`
	return query
}

func (o *OrdersRepository) CollectOrderById(ctx context.Context,
	itemRepository *ItemsRepository,
	paymentRepository *PaymentsRepository,
	deliveryRepository *DeliveriesRepository,
	orderId string) (datastruct.Order, error) {
	var order datastruct.Order
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return order, err
	}
	defer tx.Rollback(ctx)
	var p datastruct.Payment
	err = tx.QueryRow(ctx, paymentRepository.GetPaymentByOrderId(), orderId).Scan(&p.Transaction,
		&p.RequestId, &p.Currency, &p.Amount,
		&p.PaymentDt, &p.Bank, &p.DeliveryCost,
		&p.GoodsTotal, &p.CustomFee, &p.Provider)
	if err != nil {
		return order, err
	}
	var d datastruct.Delivery
	err = tx.QueryRow(ctx, deliveryRepository.GetDeliveryByOrderId(), orderId).Scan(&d.Name,
		&d.Phone, &d.Zip, &d.City,
		&d.Address, &d.Region, &d.Email)
	if err != nil {
		return order, err
	}
	var items []datastruct.Item
	rows, err := tx.Query(ctx, itemRepository.GetAllItemsByOrderId(), orderId)
	for rows.Next() {
		var i datastruct.Item
		err := rows.Scan(&i.ChrtId, &i.TrackNumber, &i.Price, &i.Rid, &i.Name,
			&i.Sale, &i.Size, &i.TotalPrice, &i.NmId, &i.Brand, &i.Status)
		if err != nil {
			return order, err
		}
		items = append(items, i)
	}
	err = tx.QueryRow(ctx, o.GetOrderById(), orderId).Scan(&order.OrderUid, &order.Entry,
		&order.Locale, &order.InternalSignature, &order.CustomerId,
		&order.DeliveryService, &order.Shardkey, &order.SmId,
		&order.DateCreated, &order.OofShard, &order.TrackNumber)
	if err != nil {
		return order, err
	}
	order.Payment = p
	order.Delivery = d
	order.Items = items
	return order, nil
}

func (o *OrdersRepository) AddNewOrderInfo() string {
	query := `(INSERT INTO orders (order_uid, entry, delivery, 
	                                        payment, locale, internal_signature,
											customer_id, delivery_service, shardkey,
											sm_id, date_created, oof_shard, track_number )
                        VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13)
						RETURNING order_uid)`
	return query
}

func (o *OrdersRepository) AddNewOrder(ctx context.Context,
	itemRepository *ItemsRepository,
	paymentRepository *PaymentsRepository,
	deliveryRepository *DeliveriesRepository,
	orderItemRepository OrderItemsRepository,
	order datastruct.Order) error {
	p := order.Payment
	d := order.Delivery
	i := order.Items
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)
	transactionId := 0
	err = tx.QueryRow(ctx, paymentRepository.AddPayment(),
		p.Transaction, p.RequestId,
		p.Currency, p.Provider, p.Amount,
		p.PaymentDt, p.Bank, p.DeliveryCost,
		p.GoodsTotal, p.CustomFee).Scan(&transactionId)
	if err != nil {
		return err
	}

	deliveryId := 0
	err = tx.QueryRow(ctx, deliveryRepository.AddDelivery(),
		d.Name, d.Phone, d.Zip,
		d.City, d.Address,
		d.Region, d.Email).Scan(&deliveryId)
	if err != nil {
		return err
	}
	itemIds := make([]int, 0)
	for _, item := range i {
		itemId := 0
		err = tx.QueryRow(ctx, itemRepository.AddItem(),
			item.ChrtId, item.TrackNumber, item.Price,
			item.Rid, item.Name,
			item.Sale, item.Size,
			item.TotalPrice, item.NmId,
			item.Brand, item.Status).Scan(&itemId)
		if err != nil {
			return err
		}
		itemIds = append(itemIds, itemId)
	}
	orderId := ""
	err = tx.QueryRow(ctx, o.AddNewOrderInfo(),
		order.OrderUid, order.Entry, deliveryId,
		transactionId, order.Locale,
		order.InternalSignature, order.CustomerId,
		order.DeliveryService, order.Shardkey,
		order.SmId, order.DateCreated, order.OofShard, order.TrackNumber).Scan(&orderId)
	if err != nil {
		return err
	}
	for _, item := range itemIds {
		_, err = tx.Exec(ctx, orderItemRepository.AddOrderItem(), item, orderId)
		if err != nil {
			return err
		}
	}
	err = tx.Commit(ctx)
	if err != nil {
		return err
	}
	fmt.Println(transactionId, deliveryId)
	return nil
}

func (o *OrdersRepository) GetNLastOrders(ctx context.Context,
	itemRepository *ItemsRepository,
	paymentRepository *PaymentsRepository,
	deliveryRepository *DeliveriesRepository,
	n int) ([]datastruct.Order, error) {
	var orders []datastruct.Order
	query := `(SELECT order_uid FROM orders
               LIMIT $1)`
	rows, err := o.db.Query(ctx, query, n)
	if err != nil {
		return orders, err
	}
	for rows.Next() {
		orderId := ""
		err := rows.Scan(&orderId)
		if err != nil {
			return orders, err
		}
		order, err := o.CollectOrderById(ctx,
			itemRepository,
			paymentRepository,
			deliveryRepository,
			orderId)
		orders = append(orders, order)
	}
	return orders, nil
}
