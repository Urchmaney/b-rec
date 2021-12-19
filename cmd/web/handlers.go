package main

import(
  "net/http"
  "encoding/json"
  "b-rec/pkg/models/mysql"
)

type UserHandler struct {
  userService mysql.UserDAO
}

func(usr_handler UserHandler) home(w http.ResponseWriter, r *http.Request) {
  data, _ := usr_handler.userService.GetUser()
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(data)
}