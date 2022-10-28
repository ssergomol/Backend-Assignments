package database

import "backend-assignments/l0/pkg/models"

type ItemRepo struct {
	store *Store
}

func (repo *ItemRepo) Create(item models.Item) {
	repo.store.db.QueryRow(
		`INSERT INTO item (chrt_id, order_uid, track_number, price, rid, name, sale, size,
			 total_price, nm_id, brand, status) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12)`,
		item.ChrtID, item.OrderUID, item.TrackNumber, item.Price, item.Rid, item.Name,
		item.Sale, item.Size, item.TotalPrice, item.NmID, item.Brand, item.Status,
	)

}
