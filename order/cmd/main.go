package main

import (
	"context"
	"net/http"
	"os"
	"practice/ecom/order/endpoints"
	logconst "practice/ecom/order/svcparams"
	"time"

	"github.com/go-kit/log"
	"github.com/gorilla/mux"
)

func main() {
	ctx := context.Background()

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
	logger = log.With(logger, "TS", time.Now().UTC().UnixMilli())
	logger = log.With(logger, "Service", "OrderService")
	logger = log.With(logger, "Caller", log.DefaultCaller)

	// init mux router
	r := mux.NewRouter()
	err := endpoints.NewRoute(r)
	if err != nil {
		logger.Log(ctx, logconst.Layer, logconst.MainLayer, "NewRouterError", err)
		os.Exit(1)
	}

	_ = logger.Log(logconst.Message, "E-Com : [Order] service running on http://localhost:8080/order")
	_ = http.ListenAndServe(":8080", r)
}
