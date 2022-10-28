package database

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Store struct {
	db           *sql.DB
	Config       *ConfigDB
	orderRepo    *OrderRepo
	itemRepo     *ItemRepo
	deliveryRepo *DeliveryRepo
	paymentRepo  *PaymentRepo
}

func NewDB(config *ConfigDB) *Store {
	return &Store{
		Config: config,
	}
}

func (s *Store) Connect() error {

	db, err := sql.Open("postgres", s.Config.databaseURL)
	if err != nil {
		return err
	}

	err = db.Ping()
	if err != nil {
		return err
	}

	s.db = db

	return nil
}

func (s *Store) Disconnect() error {
	return s.db.Close()
}

func (s *Store) Order() *OrderRepo {
	if s.orderRepo != nil {
		return s.orderRepo
	}

	s.orderRepo = &OrderRepo{
		store: s,
	}

	return s.orderRepo
}

func (s *Store) Item() *ItemRepo {
	if s.itemRepo != nil {
		return s.itemRepo
	}

	s.itemRepo = &ItemRepo{
		store: s,
	}

	return s.itemRepo
}

func (s *Store) Delivery() *DeliveryRepo {
	if s.itemRepo != nil {
		return s.deliveryRepo
	}

	s.deliveryRepo = &DeliveryRepo{
		store: s,
	}

	return s.deliveryRepo
}

func (s *Store) Payment() *PaymentRepo {
	if s.paymentRepo != nil {
		return s.paymentRepo
	}

	s.paymentRepo = &PaymentRepo{
		store: s,
	}

	return s.paymentRepo
}
