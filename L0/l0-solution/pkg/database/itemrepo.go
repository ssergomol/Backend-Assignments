package database

import (
	"backend-assignments/l0/pkg/models"
	"log"
)

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

func (repo *ItemRepo) GetData() []models.Item {
	rows, err := repo.store.db.Query("SELECT * from item")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	items := []models.Item{}
	for rows.Next() {
		item := models.Item{}
		err := rows.Scan(&item.ChrtID, &item.OrderUID, &item.TrackNumber, &item.Price, &item.Rid, &item.Name,
			&item.Sale, &item.Size, &item.TotalPrice, &item.NmID, &item.Brand, &item.Status)
		if err != nil {
			log.Fatal(err)
		}
		items = append(items, item)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal()
	}
	return items
}
