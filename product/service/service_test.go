package service

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func Test_ListProduct(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	prodList := `{"1":{"id":1,"product_name":"IPhone12","price":50000,"category":1,"quantity":100}}`
	tests := []struct {
		name string
		args args
	}{
		{
			name: "empty_product_list",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8000/product", bytes.NewBuffer([]byte(``))),
			},
		},
		{
			name: "test_success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodGet, "http://localhost:8000/product", bytes.NewBuffer([]byte(prodList))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ListProduct(tt.args.w, tt.args.r)
		})
	}
}

func Test_UpdateProduct(t *testing.T) {
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	productNotFoundJSON := `[{"id":0,"quantity":100}]`
	successJSON := `[{"id":1,"quantity":100}]`

	tests := []struct {
		name string
		args args
	}{
		{
			name: "failed_empty_body",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "http://localhost:8000/product", nil),
			},
		},
		{
			name: "failed_product_not_found",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "http://localhost:8000/product", bytes.NewBuffer([]byte(productNotFoundJSON))),
			},
		},
		{
			name: "success",
			args: args{
				w: httptest.NewRecorder(),
				r: httptest.NewRequest(http.MethodPut, "http://localhost:8000/product", bytes.NewBuffer([]byte(successJSON))),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateProduct(tt.args.w, tt.args.r)
		})
	}
}
