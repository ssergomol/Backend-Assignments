package database

import (
	"backend-assignments/l0/pkg/models"
	"log"
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

func (repo *PaymentRepo) GetData() []models.Payment {
	rows, err := repo.store.db.Query("SELECT * from payment")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	payments := []models.Payment{}
	for rows.Next() {
		payment := models.Payment{}
		err := rows.Scan(&payment.Transaction, &payment.OrderUID, &payment.RequestID, &payment.Currency, &payment.Provider,
			&payment.Amount, &payment.PaymentDT, &payment.Bank, &payment.DeliveryCost, &payment.GoodsTotal,
			&payment.CustomFee)
		if err != nil {
			log.Fatal(err)
		}
		payments = append(payments, payment)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal()
	}
	return payments
}
