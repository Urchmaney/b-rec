package main

import(
  "net/http"
  "encoding/json"
  "b-rec/pkg/models/mysql"
)

func home(w http.ResponseWriter, r *http.Request) {
  ua := mysql.UserDAO{}
  data, _ := ua.GetUser()
  w.Header().Set("Content-Type", "application/json")
  json.NewEncoder(w).Encode(data)
}