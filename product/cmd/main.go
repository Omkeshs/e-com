package main

import (
	"context"
	"net/http"
	"os"
	"practice/ecom/product/endpoints"
	logconst "practice/ecom/product/svcparams"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

func main() {
	ctx := context.Background()
	// set global log level
	logger.SetLevel(logger.DebugLevel)

	r := mux.NewRouter()

	err := endpoints.NewRoute(r)
	if err != nil {
		logger.Debug(ctx, logconst.Layer, logconst.MainLayer, "NewRouterError", err)
		os.Exit(1)
	}

	logger.Info("E-Com : [Product] service running on http://localhost:8000/product")
	logger.Fatal(http.ListenAndServe(":"+logconst.ProductSVCPort, r))

}
