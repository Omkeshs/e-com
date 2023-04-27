package main

import (
	"context"
	"net/http"
	"os"
	"practice/ecom/order/endpoints"
	logconst "practice/ecom/order/svcparams"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	logger "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	err := godotenv.Load(".env")
	if err != nil {
		logger.Debug(ctx, logconst.Layer, logconst.MainLayer, "NewRouterError", err)
		os.Exit(1)
	}
	// set global log level
	logger.SetLevel(logger.DebugLevel)

	r := mux.NewRouter()

	err = endpoints.NewRoute(r)
	if err != nil {
		logger.Debug(ctx, logconst.Layer, logconst.MainLayer, "NewRouterError", err)
		os.Exit(1)
	}

	logger.Info("E-Com : [Order] service running on http://localhost:8080/order")
	logger.Fatal(http.ListenAndServe(":"+os.Getenv("ORDERSVCPORT"), r))
}
