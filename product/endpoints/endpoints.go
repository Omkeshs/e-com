package endpoints

import (
	"errors"
	"net/http"

	"practice/ecom/product/service"
	"practice/ecom/product/spec"

	"github.com/gorilla/mux"
)

func NewRoute(r *mux.Router) error {
	if r == nil {
		return errors.New("mux router is nil")
	}

	r.HandleFunc(spec.ListProductURL, service.ListProduct).Methods(http.MethodGet)
	r.HandleFunc(spec.UpdateProductURL, service.UpdateProduct).Methods(http.MethodPut)

	return nil
}
