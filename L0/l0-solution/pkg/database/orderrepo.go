package database

import "backend-assignments/l0/pkg/models"

type OrderRepo struct {
	store *Store
}

func (repo *OrderRepo) Create(order *models.Order) {
	repo.store.db.QueryRow(
		`INSERT INTO orders (order_uid, track_number, entry, locale, internal_signature,
			 customer_id, delivery_service, shardkey, sm_id, date_created, oof_shard)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		order.OrderUID, order.TrackNumber, order.Entry, order.Locale, order.InternalSignature,
		order.CustomerID, order.DeliveryService, order.ShardKey, order.SmID, order.DateCreated,
		order.OofShard,
	)
}