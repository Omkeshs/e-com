package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/jarcoal/httpmock"
)

func TestPlaceOrder(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	productNotFoundJSON := `{"items":[{"product_id":100,"quantity":10}]}`
	productJSON := `{"1":{"id":1,"product_name":"IPhone12","price":50000,"category":1,"quantity":100}}}`
	maxLimitOrder := `{"items":[{"product_id":1,"quantity":100}]}`
	orderURL := "http://localhost:" + os.Getenv("ORDERSVCPORT") + "/order"
	productURL := "http://localhost:" + os.Getenv("PRODUCTSVCPORT") + "/product"

	tests := []struct {
		name        string
		args        args
		prepareTest func()
	}{
		{
			name: "failed_invalid_request",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, nil),
			},
			prepareTest: func() {},
		},
		{
			name: "failed_empty_item_list",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, bytes.NewBuffer([]byte(`{}`))),
			},
			prepareTest: func() {},
		},
		{
			name: "failed_get_product",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, bytes.NewBuffer([]byte(productNotFoundJSON))),
			},
			prepareTest: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", "http://localhost:8000/invalid", httpmock.NewStringResponder(200, `{}`))
			},
		},
		{
			name: "failed_get_product_decode",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, bytes.NewBuffer([]byte(productNotFoundJSON))),
			},
			prepareTest: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", productURL, httpmock.NewStringResponder(200, `{`))
			},
		},
		{
			name: "failed_product_not_found",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, bytes.NewBuffer([]byte(productNotFoundJSON))),
			},
			prepareTest: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", productURL, httpmock.NewStringResponder(200, productJSON))
			},
		},
		{
			name: "failed_max_product_limt",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, orderURL, bytes.NewBuffer([]byte(maxLimitOrder))),
			},
			prepareTest: func() {
				httpmock.Activate()
				httpmock.RegisterResponder("GET", productURL, httpmock.NewStringResponder(200, productJSON))
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.prepareTest()
			PlaceOrder(tt.args.w, tt.args.r)
		})
	}
}
