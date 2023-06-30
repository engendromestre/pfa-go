package database

import (
	"database/sql"

	"github.com/engendromestre/pfa-go/internal/order/entity"
	_ "github.com/go-sql-driver/mysql"
)

type OrderRespository struct {
	Db *sql.DB
}

func NewOrderRespository(db *sql.DB) *OrderRespository {
	return &OrderRespository{Db: db}
}

func (r *OrderRespository) Save(order *entity.Order) error {
	stmt, err := r.Db.Prepare("INSERT INTO orders (id, price, tax, final_price) VALUES (?, ?, ?, ?)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(order.ID, order.Price, order.Tax, order.FinalPrice)
	if err != nil {
		return err
	}
	return nil
}
