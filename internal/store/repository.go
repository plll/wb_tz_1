package store

import (
	"context"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/plll/wb_tz_1/internal/datastruct"
)

type Orders interface {
	GetOrderById() string
	AddNewOrder(ctx context.Context,
		itemRepository Items,
		paymentRepository Payments,
		deliveryRepository Deliveries,
		orderItemRepository OrdersItems,
		order datastruct.Order) error
	CollectOrderById(ctx context.Context,
		itemRepository Items,
		paymentRepository Payments,
		deliveryRepository Deliveries,
		orderId string) (datastruct.Order, error)
	GetNLastOrders(ctx context.Context,
		itemRepository Items,
		paymentRepository Payments,
		deliveryRepository Deliveries,
		n int) ([]datastruct.Order, error)
}

type Deliveries interface {
	AddDelivery() string
	GetDeliveryByOrderId() string
}

type Items interface {
	AddItem() string
	GetAllItemsByOrderId() string
}

type OrdersItems interface {
	AddOrderItem() string
}

type Payments interface {
	AddPayment() string
	GetPaymentByOrderId() string
}

type Repositories struct {
	Deliveries  Deliveries
	Orders      Orders
	Items       Items
	Payments    Payments
	OrdersItems OrdersItems
}

func NewRepositories(db *pgxpool.Pool) *Repositories {
	return &Repositories{
		Deliveries:  NewDeliveriesRepository(db),
		Payments:    NewPaymentsRepository(db),
		Orders:      NewOrdersRepository(db),
		Items:       NewItemsRepository(db),
		OrdersItems: NewOrderItemsRepository(db),
	}
}
