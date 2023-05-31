package endpoints

import (
	"errors"
	"net/http"

	"practice/ecom/order/service"
	"practice/ecom/order/spec"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func NewRoute(r *mux.Router) error {
	if r == nil {
		return errors.New("mux router is nil")
	}

	// set global log level
	logger.SetLevel(logger.DebugLevel)

	// Handlers ...
	r.HandleFunc(spec.ListOrderURL, service.ListOrder).Methods(http.MethodGet)
	r.HandleFunc(spec.PlaceOrderURL, service.PlaceOrder).Methods(http.MethodPost)
	r.HandleFunc(spec.UpdateOrderURL, service.UpdateOrder).Methods(http.MethodPut)
	return nil
}
