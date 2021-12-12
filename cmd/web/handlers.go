package main

import(
  "net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
  w.Write([]byte ("Welcome to bill recorder moving onward connected to docker" ))
}