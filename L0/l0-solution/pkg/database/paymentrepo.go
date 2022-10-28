package database

import (
	"backend-assignments/l0/pkg/models"
)

type PaymentRepo struct {
	store *Store
}

func (repo *PaymentRepo) Create(payment *models.Payment, order *models.Order) {
	repo.store.db.QueryRow(
		`INSERT INTO payment (transaction, order_uid, request_id, currency, provider, amount,
			 payment_dt, bank, delivery_cost, goods_total, custom_fee) 
			 VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11)`,
		payment.Transaction, order.OrderUID, payment.RequestID, payment.Currency, payment.Provider,
		payment.Amount, payment.PaymentDT, payment.Bank, payment.DeliveryCost, payment.GoodsTotal,
		payment.GoodsTotal, payment.CustomFee,
	)
}
