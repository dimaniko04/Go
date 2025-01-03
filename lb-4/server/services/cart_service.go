package services

import (
	"database/sql"
	"errors"
	"time"

	"github.com/dimaniko04/Go/lb-4/server/models"
)

type CartService interface {
	GetAll(userId string) ([]models.CartItem, error)
	Add(item models.CartItem) error
	Delete(id string) error
	Checkout(userId string) error
}

type cartService struct {
	db *sql.DB
}

func (s *cartService) GetAll(userId string) ([]models.CartItem, error) {
	rows, err := s.db.Query("SELECT i.*, p.image_path, p.name, p.price FROM cart_items i INNER JOIN products p ON i.product_id = p.id WHERE i.user_id = ?", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	cartItems := []models.CartItem{}

	for rows.Next() {
		i := models.CartItem{}
		err := rows.Scan(&i.Id,
			&i.UserId,
			&i.ProductId,
			&i.Quantity,
			&i.Product.ImagePath,
			&i.Product.Name,
			&i.Product.Price)
		if err != nil {
			return nil, err
		}
		cartItems = append(cartItems, i)
	}

	return cartItems, nil
}

func (s *cartService) Add(item models.CartItem) error {
	query := "INSERT INTO cart_items (user_id, product_id, quantity) VALUES (?, ?, ?)"
	_, err := s.db.Exec(query, item.UserId, item.ProductId, item.Quantity)

	if err != nil {
		return errors.New("failed to add item to cart")
	}

	return nil
}

func (s *cartService) Delete(id string) error {
	query := "DELETE FROM cart WHERE id = ?"
	_, err := s.db.Exec(query, id)
	if err != nil {
		return errors.New("failed to remove item from cart")
	}

	return nil
}

func (s *cartService) Checkout(userId string) error {
	query := "INSERT INTO orders (user_id, created_at) VALUES (?, ?)"
	result, err := s.db.Exec(query, userId, time.Now())
	if err != nil {
		return errors.New("failed to create order")
	}

	orderId, _ := result.LastInsertId()

	itemsQuery := `
        INSERT INTO order_items (order_id, product_id, quantity)
        SELECT ?, c.product_id, c.quantity
        FROM cart c
        JOIN products p ON c.product_id = p.id
        WHERE c.user_id = ?`
	_, err = s.db.Exec(itemsQuery, orderId, userId)
	if err != nil {
		return errors.New("failed to add items to order")
	}

	clearCartQuery := "DELETE FROM cart WHERE user_id = ?"
	_, err = s.db.Exec(clearCartQuery, userId)
	if err != nil {
		return errors.New("failed to clear user cart")
	}

	return nil
}
