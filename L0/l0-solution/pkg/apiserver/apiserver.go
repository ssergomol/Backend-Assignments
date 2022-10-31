package apiserver

import (
	"backend-assignments/l0/pkg/database"
	"backend-assignments/l0/pkg/streaming"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

type APIserver struct {
	config *Config
	router *mux.Router
	logger *logrus.Logger
	db     *database.Store
	cache  map[string]streaming.JSONstructure
}

func NewServer(config *Config) *APIserver {
	server := &APIserver{
		config: config,
		router: mux.NewRouter(),
		logger: logrus.New(),
		cache:  make(map[string]streaming.JSONstructure),
	}

	server.configureDatabase()
	server.recoverCache()
	return server
}

func (server *APIserver) Start() error {
	if err := server.configureLogger(); err != nil {
		return err
	}
	server.configureRouter()

	server.logger.Info("starting streaming API server")
	return http.ListenAndServe(server.config.BindAddr, server.router)
}

func (server *APIserver) configureLogger() error {
	level, err := logrus.ParseLevel(server.config.LogLevel)
	if err != nil {
		return err
	}

	logrus.SetLevel(level)
	return nil
}

func (server *APIserver) configureRouter() {
	server.RegisterGetOrderByID()
	server.RegisterGetAllOrders()
}

func (server *APIserver) configureDatabase() error {
	configDB := database.NewConfig()
	db := database.NewDB(configDB)
	server.logger.Info("connecting to database")
	if err := db.Connect(); err != nil {
		return err
	}
	server.logger.Info("database connected")

	server.db = db
	return nil
}

func (server *APIserver) UpdateDatabase(data streaming.JSONstructure) {
	server.cache[data.OrderUID] = data
	server.logger.Info("update database")
	server.db.Order().Create(data.Order)
	server.db.Payment().Create(data.Payment)
	server.db.Delivery().Create(data.Delivery)
	for _, item := range data.Items {
		server.db.Item().Create(item)
	}
}

func (server *APIserver) recoverCache() {
	ordersList := server.db.Order().GetAllData()

	for _, order := range ordersList {
		paymnet := server.db.Payment().GetDataByUID(order.OrderUID)
		delivery := server.db.Delivery().GetDataByUID(order.OrderUID)
		items := server.db.Item().GetDataByUID(delivery.OrderUID)
		combinedData := streaming.JSONstructure{
			Order:    order,
			Delivery: delivery,
			Payment:  paymnet,
			Items:    items,
		}
		server.cache[order.OrderUID] = combinedData
	}
}
