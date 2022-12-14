package database

import (
	"backend-assignments/l0/pkg/models"

	"github.com/sirupsen/logrus"
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

func (repo *DeliveryRepo) GetDataByUID(order_uid string) models.Delivery {
	delivery := models.Delivery{}
	err := repo.store.db.QueryRow("SELECT * from delivery where order_uid = $1", order_uid).Scan(
		&delivery.DeliveryID, &delivery.OrderUID, &delivery.Name, &delivery.Phone, &delivery.Zip,
		&delivery.City, &delivery.Address, &delivery.Region, &delivery.Email,
	)
	if err != nil {
		logrus.Fatal(err)
	}
	return delivery
}
