package mysql

import (
  "fmt"
  "time"
  "database/sql"
  "b-rec/pkg/models"
)

type BillService struct {
  DB *sql.DB
}

func (service BillService) AddBill(bill models.Bill) (int64, error) {
  res, err := service.DB.Exec("INSERT INTO Bills(Amount, Month, AccountId, CreatedAt) values (?,?,?,?)", bill.Amount, bill.Month, bill.AccountId, time.Now())
  if err != nil {
    fmt.Println(err)
    return 0, err
  }

  inserted_id, _ := res.LastInsertId()
  return inserted_id, nil
}