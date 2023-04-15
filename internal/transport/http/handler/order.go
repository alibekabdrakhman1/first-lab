package handler

import (
	"encoding/json"
	"github.com/alibekabdrakhman1/first-lab/internal/model"
	"github.com/alibekabdrakhman1/first-lab/internal/service"
	"net/http"
	"strings"
)

type IOrderHandler interface {
	GetAll(http.ResponseWriter, *http.Request)
	GetNotReturned(http.ResponseWriter, *http.Request)
	GetCountOrders(http.ResponseWriter, *http.Request)
	Create(http.ResponseWriter, *http.Request)
	Delete(http.ResponseWriter, *http.Request)
	ReturnBook(w http.ResponseWriter, r *http.Request)
}

type OrderHandler struct {
	Service *service.Manager
}

func (h *OrderHandler) GetAll(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.Order.GetAll()
	_ = json.NewEncoder(writer).Encode(res)
}

func (h *OrderHandler) GetNotReturned(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.Order.GetNotReturned()
	_ = json.NewEncoder(writer).Encode(res)
}

func (h *OrderHandler) GetCountOrders(writer http.ResponseWriter, request *http.Request) {
	res, _ := h.Service.Order.GetCountOrders()
	_ = json.NewEncoder(writer).Encode(res)
}

func (h *OrderHandler) Create(writer http.ResponseWriter, request *http.Request) {
	var createOrder model.Order
	_ = json.NewDecoder(request.Body).Decode(&createOrder)
	_ = h.Service.Order.Create(createOrder)
	_, _ = writer.Write([]byte("Order created"))
}

func (h *OrderHandler) Delete(writer http.ResponseWriter, request *http.Request) {
	orderId := strings.TrimPrefix(request.URL.Path, "/api/orders/")
	_ = h.Service.Order.Delete(orderId)
	_, _ = writer.Write([]byte("Book deleted"))
}
func (h *OrderHandler) ReturnBook(writer http.ResponseWriter, request *http.Request) {
	orderId := strings.TrimPrefix(request.URL.Path, "/api/orders/return/")
	_ = h.Service.Order.ReturnBook(orderId)
	_, _ = writer.Write([]byte("Book returned"))
}

func NewOrderHandler(s *service.Manager) *OrderHandler {
	return &OrderHandler{
		Service: s,
	}
}
