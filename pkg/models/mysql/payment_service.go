package mysql

import (
  "fmt"
  "b-rec/pkg/models"
  "database/sql"
)

type PaymentService struct {
  DB *sql.DB
}

func (service PaymentService) SavePayments(payments []models.Payment) (int, error) {
  var count int
  for _, val := range payments {
    res, err := service.DB.Exec("INSERT INTO Payments(Amount, UserId, Month, CreatedAt) values (?,?,?,?)", val.Amount, val.UserId, val.Month, time.Now())
    if err != nil {
      fmt.Println(err)
      continue
    }
    count++
  }
  return count
}