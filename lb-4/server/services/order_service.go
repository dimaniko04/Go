package services

import (
	"database/sql"
	"errors"

	"github.com/dimaniko04/Go/lb-4/server/models"
)

type OrderService interface {
	GetUserOrders(userId string) ([]models.Order, error)
}

type orderService struct {
	db *sql.DB
}

func (s *orderService) GetUserOrders(userId string) ([]models.Order, error) {
	ordersQuery := `
        SELECT o.*, oi.id, oi.quantity, p.name, p.image_path, p.price  
		FROM orders o 
		INNER JOIN order_items oi
			ON oi.order_id = o.id
		INNER JOIN products p 
			ON p.id = oi.product_id
		ORDER BY o.created_at DESC`

	rows, err := s.db.Query(ordersQuery, userId)
	if err != nil {
		return nil, errors.New("failed to retrieve user orders")
	}
	defer rows.Close()

	var orders []models.Order
	var orderItems []models.OrderItem

	for rows.Next() {
		var o models.Order
		var oi models.OrderItem

		if err := rows.Scan(
			&o.Id,
			&o.UserId,
			&o.CreatedAt,
			&oi.Id,
			&oi.Quantity,
			&oi.Product.Name,
			&oi.Product.ImagePath,
			&oi.Product.Price); err != nil {
			return nil, errors.New("failed to parse order data")
		}

		last := len(orders) - 1
		if last < 0 {
			orders = append(orders, o)
		} else if orders[last].Id != o.Id {
			orders[last].Items = orderItems
			orders = append(orders, o)
		} else {
			orderItems = append(orderItems, oi)
		}
	}
	orders[len(orders)-1].Items = orderItems

	return orders, nil
}
