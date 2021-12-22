package mysql

import (
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
    return 0, err
  }
  inserted_id, _ := res.LastInsertId()
  return inserted_id, nil
}

func (user_doa UserDAO) GetAllUsers()([] models.User, error) {
  return [] models.User {}, nil
}

func (user_doa UserDAO) GetUser()(models.User, error) {
  return models.User{ ID: 1, Name: "Kingsley", AccountId: 2, }, nil
}