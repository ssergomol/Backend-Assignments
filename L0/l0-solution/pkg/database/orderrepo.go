package database

import (
	"backend-assignments/l0/pkg/models"
	"log"
)

type OrderRepo struct {
	store *Store
}

func (repo *OrderRepo) Create(order models.Order) {
	repo.store.db.QueryRow(
		`INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
			 customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated,
		order.OofShard,
	)
}

func (repo *OrderRepo) GetData() []models.Order {
	rows, err := repo.store.db.Query("SELECT * from orders")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	orders := []models.Order{}
	for rows.Next() {
		order := models.Order{}
		err := rows.Scan(&order)
		if err != nil {
			log.Fatal(err)
		}
		orders = append(orders, order)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal()
	}
	return orders
}
