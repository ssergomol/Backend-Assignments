package database

import (
	"backend-assignments/l0/pkg/models"

	"github.com/sirupsen/logrus"
)

type PaymentRepo struct {
	store *Store
}

func (repo *PaymentRepo) Create(payment models.Payment) {
	repo.store.db.QueryRow(
		`INSERT INTO payment (transaction, order_uid, request_id, currency, provider, amount,
			payment_dt, bank, delivery_cost, goods_total, custom_fee) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		payment.Transaction, payment.OrderUID, payment.RequestID, payment.Currency, payment.Provider,
		payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal,
		payment.CustomFee,
	)
}

func (repo *PaymentRepo) GetDataByUID(order_uid string) models.Payment {
	payment := models.Payment{}
	err := repo.store.db.QueryRow("SELECT * from payment where order_uid = $1", order_uid).Scan(
		&payment.Transaction, &payment.OrderUID, &payment.RequestID, &payment.Currency, &payment.Provider,
		&payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal,
		&payment.CustomFee,
	)
	if err != nil {
		logrus.Fatal(err)
	}

	return payment
}
