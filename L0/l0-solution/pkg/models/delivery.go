package models

type Delivery struct {
	DeliveryID int    `json:"delivery_id"`
	OrderUID   string `json:"order_uid"`
	Name       string `json:"name"`
	Phone      string `json:"phone"`
	Zip        string `json:"zip"`
	City       string `json:"city"`
	Address    string `json:"address"`
	Region     string `json:"region"`
	Email      string `json:"email"`
}
