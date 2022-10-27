package apiserver

func (server *APIserver) RegisterHome() {
	server.router.HandleFunc("/home", HomeHandeler())
}
