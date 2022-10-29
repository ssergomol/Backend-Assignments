package database

import (
	"backend-assignments/l0/pkg/models"
	"log"
)

type DeliveryRepo struct {
	store *Store
}

func (repo *DeliveryRepo) Create(delivery models.Delivery) {
	repo.store.db.QueryRow(
		`INSERT INTO delivery (delivery_id, order_uid, name, phone, zip, city, address,
			 region, email) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		delivery.DeliveryID, delivery.OrderUID, delivery.Name, delivery.Phone, delivery.Zip,
		delivery.City, delivery.Address, delivery.Region, delivery.Email,
	)
}

func (repo *DeliveryRepo) GetData() []models.Delivery {
	rows, err := repo.store.db.Query("SELECT * from delivery")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	deliveries := []models.Delivery{}
	for rows.Next() {
		delivery := models.Delivery{}
		err := rows.Scan(&delivery.DeliveryID, &delivery.OrderUID, &delivery.Name, &delivery.Phone, &delivery.Zip,
			&delivery.City, &delivery.Address, &delivery.Region, &delivery.Email)
		if err != nil {
			log.Fatal(err)
		}
		deliveries = append(deliveries, delivery)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal()
	}
	return deliveries
}
