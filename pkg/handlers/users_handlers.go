package handlers

import (
  "time"
  "log"
  "net/http"
  "encoding/json"
  "b-rec/pkg/models/mysql"
  "b-rec/pkg/models"
)

type UserHandler struct {
  UserService mysql.UserDAO
}

func (handler UserHandler) CreateUser(rw http.ResponseWriter, req *http.Request) {
  decoder := json.NewDecoder(req.Body)
  var user models.User
  if code_err := decoder.Decode(&user); code_err != nil {
    log.Println(code_err)
    return
  }
  user.CreatedAt = time.Now()
  
  if _, err := handler.UserService.CreateUser(user); err != nil {
    log.Println(err)
    return
  }
  rw.Header().Set("Content-Type", "application/json")
  json.NewEncoder(rw).Encode(user)
}

func (handler UserHandler) GetAllUsers(rw http.ResponseWriter, req *http.Request) {
  users, err := handler.UserService.GetAllUsers()
  if err != nil {
    log.Println(err)
    return
  }
  rw.Header().Set("Content-Type", "application/json")
  json.NewEncoder(rw).Encode(users)
}