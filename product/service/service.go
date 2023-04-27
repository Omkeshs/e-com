package service

import (
	"encoding/json"
	"net/http"
	"practice/ecom/product/spec"

	logconst "practice/ecom/product/svcparams"

	logger "github.com/sirupsen/logrus"
)

// ProdMap - Product Map ...
var ProdMap = map[int32]spec.Product{
	1: {
		ID:       1,
		Name:     "IPhone12",
		Price:    50000,
		Category: 1,
		Quantity: 100,
	},
	2: {
		ID:       2,
		Name:     "IPhone13",
		Price:    60000,
		Category: 1,
		Quantity: 200,
	},
	3: {
		ID:       3,
		Name:     "IPhone14",
		Price:    70000,
		Category: 1,
		Quantity: 200,
	},
	101: {
		ID:       101,
		Name:     "IPad",
		Price:    30000,
		Category: 2,
		Quantity: 50,
	},
	102: {
		ID:       102,
		Name:     "IPad 3rd gen",
		Price:    330000,
		Category: 2,
		Quantity: 50,
	},
	201: {
		ID:       201,
		Name:     "Mobile case",
		Price:    100,
		Category: 3,
		Quantity: 50,
	},
}

// ListProduct - product details
func ListProduct(w http.ResponseWriter, r *http.Request) {
	if len(ProdMap) == 0 {
		logger.Debug(logconst.Layer, logconst.RouterLayer, logconst.Message, "Empty Product List")
		w.Write([]byte("Empty Product list"))
		return
	}
	// return product details
	logger.Debug(logconst.Layer, logconst.RouterLayer, logconst.Message, "product list fetch successfully.")
	json.NewEncoder(w).Encode(ProdMap)
}

func UpdateProduct(w http.ResponseWriter, r *http.Request) {
	productReq := []spec.ProductRequest{}
	err := json.NewDecoder(r.Body).Decode(&productReq)
	if err != nil {
		logger.Debug(logconst.Layer, logconst.RouterLayer, "Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode body"))
		return
	}

	for _, req := range productReq {
		product, ifExist := ProdMap[req.ID]
		if ifExist {
			// Currently we support only product quantity to update
			product.Quantity = req.Quantity
			ProdMap[product.ID] = product
		} else {
			logger.Debug(logconst.Layer, logconst.RouterLayer, "Product not found ")
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Invalid product id - Product not found"))
			return
		}
	}

	json.NewEncoder(w).Encode(ProdMap)
	logger.Debug(logconst.Layer, logconst.RouterLayer, logconst.Message, "product list successfully updated.")

}
