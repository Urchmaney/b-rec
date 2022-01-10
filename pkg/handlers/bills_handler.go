package handlers

import (
  "log"
  "net/http"
  "encoding/json"
  "b-rec/pkg/models/mysql"
  "b-rec/pkg/models"
)

type BillHandler struct {
  BillService mysql.BillService
}

func (handler BillHandler) CreateBill(rw http.ResponseWriter, req *http.Request) {
  body := json.NewDecoder(req.Body)
  var bill models.Bill
  if decode_err := body.Decode(&bill); decode_err != nil {
    log.Println(decode_err)
    return
  }
  id, err := handler.BillService.AddBill(bill)
  if err != nil {
    log.Println(err)
    return
  }
  bill.ID = id
  rw.Header().Set("Content-Type", "application/json")
  json.NewEncoder(rw).Encode(bill)
}