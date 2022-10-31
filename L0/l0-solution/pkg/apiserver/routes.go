package apiserver

func (server *APIserver) RegisterGetOrderByID() {
	server.router.HandleFunc("/orders/{order_uid}", GetOrderHandler(server)).Methods("GET")
}

func (server *APIserver) RegisterGetAllOrders() {
	server.router.HandleFunc("/orders", GetAllOrdersHandler(server)).Methods("GET")
}
