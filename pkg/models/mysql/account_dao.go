package mysql

import (
	"fmt"
  "time"
  "database/sql"
  "b-rec/pkg/models"
)

type AccountDAO struct {
  DB *sql.DB    
}

func (acc_dao AccountDAO) CreateAccount(acc models.Account) (int64, error) {
  res, err := acc_dao.DB.Exec("INSERT into Accounts (OwnerFullName, Email, Password, StartingDebt, CreatedAt) VALUES (?,?,?,?,?)", acc.OwnerFullName, acc.Email, acc.Password, acc.StartingDebt, time.Now())
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  inserted_id, _ := res.LastInsertId()
  return inserted_id, nil
}

func (acc_dao AccountDAO) GetAccount(id int64) (*models.Account, error) {
  var account models.Account
  row := acc_dao.DB.QueryRow("SELECT FROM Accounts WHERE id = ?", id)
  if err := row.Scan(&account.ID, &account.OwnerFullName, &account.Password, &account.Email, &account.StartingDebt); err != nil {
    fmt.Println(err)
    return nil, err
  }
  return &account, nil
}

func (acc_dao AccountDAO) GetAccountByEmail(email string) (*models.Account, error) {
  var account models.Account
  row := acc_dao.DB.QueryRow("SELECT Id, Email, Password FROM Accounts WHERE email = ?", email)
  if err := row.Scan(&account.ID, &account.Email, &account.Password); err != nil {
    fmt.Println(err)
    return nil, err
  }
  return &account, nil
}