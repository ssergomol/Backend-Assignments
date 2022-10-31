package apiserver

import (
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func GetOrderHandler(server *APIserver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Query().Has("datatype") {
			dataType := r.URL.Query().Get("datatype")
			if dataType != "json" {
				http.Redirect(w, r, r.URL.Hostname(), http.StatusSeeOther)
			}
			logrus.Info("get all order query")
			vars := mux.Vars(r)
			order := server.cache[vars["order_uid"]]
			marshaled, err := json.Marshal(&order)
			if err != nil {
				logrus.Fatal(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.Write(marshaled)

		} else {
			vars := mux.Vars(r)
			logrus.Info(fmt.Sprintf("get order %s query", vars["order_uid"]))
			order := server.cache[vars["order_uid"]]
			templatePath := "pkg/templates/order.html"

			var err error
			tmplt := template.New(filepath.Base(templatePath))
			tmplt, err = tmplt.ParseFiles(templatePath)
			if err != nil {
				logrus.Fatal(err)
			}
			tmplt.Execute(w, order)
		}
	}
}

func GetAllOrdersHandler(server *APIserver) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logrus.Info("get all orders query")
		marshaled, err := json.Marshal(server.cache)
		if err != nil {
			logrus.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(marshaled)
	}
}
