package service

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"practice/ecom/order/spec"
	"practice/ecom/order/svcparams"
	"practice/ecom/order/svcutils"
	"strings"
	"time"

	"github.com/gorilla/mux"
	logger "github.com/sirupsen/logrus"
)

type Product struct {
	ID       int32   `json:"id"`
	Name     string  `json:"product_name"`
	Price    float32 `json:"price"`
	Category int     `json:"category"`
	Quantity int32   `json:"quantity"`
}

var orderMap = map[int32]spec.Order{}

// ListOrder - to fetch existing order list
func ListOrder(w http.ResponseWriter, r *http.Request) {
	// Check if empty product list
	if len(orderMap) == 0 {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("empty product list"))
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "empty product list")
		return
	}

	json.NewEncoder(w).Encode(orderMap)
}

// PlaceOrder ...
func PlaceOrder(w http.ResponseWriter, r *http.Request) {

	// 1. Decode request
	orderReq := spec.OrderRequest{}
	err := json.NewDecoder(r.Body).Decode(&orderReq)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode body"))
		return
	}

	if len(orderReq.Items) == 0 {
		logger.Debug("empty order request")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// 2. Get Product map from product service
	requestURL := fmt.Sprintf("%s%s%s", os.Getenv("LOCALHOSTURL"), os.Getenv("PRODUCTSVCPORT"), os.Getenv("PRODUCTSVCNAME"))
	res, err := http.Get(requestURL)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to get product list", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	prodMap := map[int32]Product{}
	err = json.NewDecoder(res.Body).Decode(&prodMap)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to get product list", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var premiumCartCount int
	var orderCount int32
	orderResp := spec.Order{}
	updateProdMap := map[int32]int32{}
	itemMap := map[int32]int32{}

	// 3. Place order
	for _, item := range orderReq.Items {

		// Check is product exist
		if product, ok := prodMap[item.ID]; !ok {
			logger.Debug("product not found")
			w.WriteHeader(http.StatusBadRequest)
			return
		} else {
			orderCount += item.Quantity
			itemMap[item.ID] += item.Quantity
			// check max limit and inventory
			if product.Quantity < item.Quantity || itemMap[item.ID] > 10 || orderCount > 10 || item.Quantity == 0 {
				logger.Debug("invalid quantity")
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("invalid item quantity / no stock available"))
				return
			}

			// check if premium product
			if product.Category == svcparams.PremiumType {
				premiumCartCount += 1
			}

			orderResp.Amount += product.Price * float32(item.Quantity)
			orderResp.Quantity += item.Quantity

			// create product id wise map for item quantity
			updateProdMap[item.ID] = product.Quantity - item.Quantity
		}

		// creating order id - existing size + 1
		orderResp.ID = int32(len(orderMap) + 1)

		// initially marked as placed
		orderResp.Status = svcparams.Placed
	}

	// 4. check if discount applicable
	if premiumCartCount >= svcparams.PremiumTypeDiscountCount {
		orderResp.Discount = orderResp.Amount / svcparams.PremiumTypeDiscount
		orderResp.FinalAmount = orderResp.Amount - orderResp.Discount
	} else {
		orderResp.FinalAmount = orderResp.Amount
	}

	orderMap[orderResp.ID] = orderResp

	// 5. Update Product list
	updateProdReq := []spec.UpdateProductRequest{}
	for productID, productQuantity := range updateProdMap {
		uProd := spec.UpdateProductRequest{
			ID:       productID,
			Quantity: productQuantity,
		}
		updateProdReq = append(updateProdReq, uProd)
	}

	body, err := json.Marshal(updateProdReq)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to marshal body in update product", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 6. Call to product service to update product quantity
	updateProdReqURL := fmt.Sprintf(os.Getenv("LOCALHOSTURL")+"%s%s", os.Getenv("PRODUCTSVCPORT"), os.Getenv("PRODUCTSVCNAME"))
	req, err := http.NewRequest(http.MethodPut, updateProdReqURL, bytes.NewBuffer(body))
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to update product list", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	req.Header.Set("Content-Type", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to update product list", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp.Body.Close()

	logger.Debug(svcparams.Layer, svcparams.RouterLayer, svcparams.Message, "Order placed successfully ")
	json.NewEncoder(w).Encode(orderResp)
}

// UpdateOrder - to update order status
func UpdateOrder(w http.ResponseWriter, r *http.Request) {

	orderReq := spec.UpdateOrderRequest{}
	vars := mux.Vars(r)
	var err error

	// 1. get orderID from updateURL
	orderReq.ID, err = svcutils.GetIntPathVariable(vars, svcparams.OrderID)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode body"))
		return
	}
	// 2. decode Update request
	err = json.NewDecoder(r.Body).Decode(&orderReq)
	if err != nil {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "Failed to decode body", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Failed to decode body"))
		return
	}

	// 3. Check order are exist in map
	order, isExist := orderMap[orderReq.ID]
	if !isExist {
		logger.Debug(svcparams.Layer, svcparams.RouterLayer, "order does not exist", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("order does not exist"))
		return
	}

	// 4. if order status was "Dispatched" then set DispatchedDate
	if strings.EqualFold(orderReq.Status, svcparams.Dispatched) {
		y, m, d := time.Now().Date()
		order.DispatchedDate = fmt.Sprintf("%d:%d:%d", d, m, y)
	}
	order.Status = orderReq.Status
	orderMap[orderReq.ID] = order
	json.NewEncoder(w).Encode(order)
}
