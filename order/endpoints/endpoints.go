package endpoints

import (
	"errors"
	"net/http"

	"practice/ecom/order/service"
	"practice/ecom/order/spec"

	"github.com/gorilla/mux"
)

func NewRoute(r *mux.Router) error {
	if r == nil {
		return errors.New("mux router is nil")
	}

	r.HandleFunc(spec.ListOrderURL, service.ListOrder).Methods(http.MethodGet)
	r.HandleFunc(spec.PlaceOrderURL, service.PlaceOrder).Methods(http.MethodPost)
	r.HandleFunc(spec.UpdateOrderURL, service.UpdateOrder).Methods(http.MethodPut)

	// TODO - delete
	// r.HandleFunc(spec.DeleteOrderURL, service.DeleteOrder).Methods(http.MethodDelete)

	return nil
}
