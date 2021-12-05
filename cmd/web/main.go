package main

import (
  "net/http"
  "log"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  
  log.Println("starting server at port :8080")
  err := http.ListenAndServe(":8080", mux)
  log.Println(err)
}
