package mysql

import (
  "database/sql"
  "b-rec/pkg/models"
)

type UserDAO struct {
  DB *sql.DB
}

func (user_doa UserDAO) CreateUser(user models.User)(int, error) {
  return 1, nil
}

func (user_doa UserDAO) GetUser()(models.User, error) {
  return models.User{ ID: 1, Name: "Kingsley", AccountId: 2, }, nil
}