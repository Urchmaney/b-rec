package mysql

import (
  "fmt"
  "database/sql"
  "b-rec/pkg/models"
)

type BillService struct {
  DB *sql.DB
}

func (service BillService) AddBill(bill models.Bill) (int64, error) {
  res, err := service.DB.Exec("INSERT INTO Bills(Amount, Month, AccountId) values (?,?,?)", bill.Amount, bill.Month, bill.AccountId)
  if err != nil {
    fmt.Println(err)
    return 0, err
  }

  inserted_id, _ := res.LastInsertId()
  return inserted_id, nil
}