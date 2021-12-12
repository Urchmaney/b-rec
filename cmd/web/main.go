package main

import (
  "net/http"
  "log"
)

func main() {
  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  
  log.Println("starting server at port :3000")
  err := http.ListenAndServe(":3000", mux)
  log.Println(err)
}
