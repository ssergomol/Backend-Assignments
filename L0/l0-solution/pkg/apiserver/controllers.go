package apiserver

import (
	"io"
	"net/http"
)

func HomeHandeler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Home")
	}
}
