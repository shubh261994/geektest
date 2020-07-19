package postgre

import (
	"fmt"

	"geektest/internal/logs"
)

type Order struct {
	ID 			int        `json:"id"`
	CreatedAt 	string     `json:"created_at"`
	CustomerId 	string     `json:"customer_id"`
	OrderName 	string     `json:"order_name"`
	Amount 		float32    `json:"amount"`
}

func GetOrdersList(offset int, limit int) []Order {
	query := fmt.Sprintf("WITH order_amounts AS (SELECT order_id, SUM(price_per_unit * quantity) as amount FROM orders JOIN order_items on orders.id=order_items.order_id group by order_id) SELECT orderss.id, orderss.created_at, orderss.customer_id, orderss.order_name, order_amounts.amount FROM orders orderss INNER JOIN order_amounts ON order_amounts.order_id = orderss.id order by orderss.id limit %d offset %d", limit, offset)
	data, _ := selectData(query)
	orders := make([]Order, 0)
	for data.Next() {
		order := Order{}
		err := data.Scan(&order.ID, &order.CreatedAt, &order.CustomerId, &order.OrderName, &order.Amount)
		if err != nil {
			logs.Error(err)
		}

		orders = append(orders, order)
	}

	return orders
}