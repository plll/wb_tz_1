package store

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/plll/wb_tz_1/internal/datastruct"
)

type Orders interface {
	AddNewOrder(ctx context.Context, order datastruct.Order) (string, error)
	GetOrderById() string
	CollectOrderById(ctx context.Context,
		itemRepository *ItemsRepository,
		paymentRepository *PaymentsRepository,
		deliveryRepository *DeliveriesRepository,
		orderId string) (datastruct.Order, error)
}

type Deliveries interface {
	AddDelivery() string
	GetDeliveryByOrderId() string
}

type Items interface {
	AddItem() string
	GetAllItemsByOrderId() string
}

type Payments interface {
	AddPayment() string
	GetPaymentByOrderId() string
}

type Repositories struct {
	Deliveries Deliveries
	Orders     Orders
	Items      Items
	Payments   Payments
}

func NewRepositories(db *pgx.Conn) *Repositories {
	return &Repositories{
		Deliveries: NewDeliveriesRepository(db),
		Payments:   NewPaymentsRepository(db),
		Orders:     NewOrdersRepository(db),
		Items:      NewItemsRepository(db),
	}
}
