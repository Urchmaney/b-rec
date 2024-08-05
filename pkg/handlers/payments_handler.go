package handlers

import (
	"fmt"
  "net/http"
  "encoding/json"
  "b-rec/pkg/models/mysql"
  "b-rec/pkg/models"
)

type PaymentHandler struct {
  PaymentService *mysql.PaymentService
}

func (handler PaymentHandler) AddPayments(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var payments []models.Payment
  if decode_err := body.Decode(&payments); decode_err != nil {
    log.Println(decode_err)
    return
  }

  res, err := handler.PaymentService.SavePayments(payments)
  if err != nil {
    log.Println(err)
    return
  }

  rw.Header().Set("Content-Type", "application/json")
  json.NewEncoder(rw).Encode(res)
}