package mysql

import (
  "fmt"
  "time"
  "database/sql"
  "b-rec/pkg/models"
)

type UserDAO struct {
  DB *sql.DB
}

func (user_doa UserDAO) CreateUser(user models.User)(int64, error) {
  res, err := user_doa.DB.Exec("INSERT INTO Users(Name, AccountId, CreatedAt) VALUES (?,?,?)", user.Name, user.AccountId, time.Now())
  if err != nil {
    fmt.Println(err)
    return 0, err
  }
  inserted_id, _ := res.LastInsertId()
  return inserted_id, nil
}

func (user_doa UserDAO) GetAllUsers()([] models.User, error) {
  rows, err := user_doa.DB.Query("SELECT Id, Name, AccountId FROM Users")
  if err != nil {
    fmt.Println(err)
    return nil, err
  }
  var users [] models.User
  for rows.Next() {
    var (
      Id int64
      Name string
      AccountId int
    )
    err := rows.Scan(&Id, &Name, &AccountId)
    if err != nil {
      fmt.Println(err)
      return nil, err
    }
    users = append(users, models.User{ ID: Id, Name: Name, AccountId: int64(AccountId) })
  }
  return users, nil
}

func (user_doa UserDAO) GetUser()(models.User, error) {
  return models.User{ ID: 1, Name: "Kingsley", AccountId: 2, }, nil
}