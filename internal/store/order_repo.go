package store

import (
	"context"
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
	orderId string) (datastract.Order, error) {
	var order Order
	tx, err := o.db.Begin(ctx)
	if err != nil {
		return order, err
	}
	defer tx.Rollback(ctx)
	var p Payment
	err = tx.QueryRow(ctx, paymentRepository.GetPaymentByOrderId(), orderId).Scan(&p.Transaction,
		&p.RequestId, &p.Currency, &p.Amount,
		&p.PaymentDt, &p.Bank, &p.DeliveryCost,
		&p.GoodsTotal, &p.CustomFee, &p.Provider)
	if err != nil {
		return order, err
	}
	var d Delivery
	err = tx.QueryRow(ctx, deliveryRepository.GetDeliveryByOrderId(), orderId).Scan(&d.Name,
		&d.Phone, &d.Zip, &d.City,
		&d.Address, &d.Region, &d.Email)
	if err != nil {
		return order, err
	}
	var items []Item
	rows, err := tx.Query(ctx, itemRepository.GetAllItemsByOrderId(), orderId)
	for rows.Next() {
		var i Item
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
		return o, err
	}
	order.Payment = p
	order.Delivery = d
	order.Items = items
	return order, nil
}

func (s *OrdersRepository) AddNewOrder(ctx context.Context, order datastruct.Orders) (string, err) {
	query := s.db.Rebind(`
		SELECT 
			s.id as session_id, s.token as token, u.id as user_id, 
		    u.name as name, u.psw as psw, 
		    u.mail as mail, u.params as params, s.user_agent as user_agent,
		    s.session_time as session_time, s.ip as ip
		FROM 
		     sessions s 
		INNER JOIN users u ON
		    u.id = s.user_id;
	`)
	return "sessions, nil"
}
