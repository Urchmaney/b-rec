 package handlers

import (
  "net/http"
  "log"
  "encoding/json"
  "b-rec/pkg/authenticator"
  "b-rec/pkg/models/mysql"
  "b-rec/pkg/models"
)

type AccountHandler struct {
  AccountService mysql.AccountDAO
  AuthenticationService authenticator.AuthenticationService
}

type LoginDetail struct {
  Email string
  Password string
}

func(handler AccountHandler) SignUp(rw http.ResponseWriter, req *http.Request) {
  body := json.NewDecoder(req.Body)
  var account models.Account
  if code_err := body.Decode(&account); code_err != nil {
    log.Println(code_err)
    return
  }
  hashPwd, err := handler.AuthenticationService.HashPassword(account.Password)
  if err != nil {
    log.Println(err)
    return
  }
  account.Password = hashPwd
  id, a_err := handler.AccountService.CreateAccount(account)
  if a_err != nil {
    log.Println(err)
    return
  }
  account.ID = int64(id)
  rw.Header().Set("Content-Type", "application/json")
  json.NewEncoder(rw).Encode(account)
}

func(handler AccountHandler) Login(rw http.ResponseWriter, req *http.Request) {
  body := json.NewDecoder(req.Body)
  var details LoginDetail
  if code_err := body.Decode(&details); code_err != nil {
    log.Println(code_err)
    return
  }
  account, err := handler.AccountService.GetAccountByEmail(details.Email)
  if err != nil {
    log.Println(err)
    return
  }
  validPwd := handler.AuthenticationService.CheckPasswordHash(details.Password, account.Password)
  if !validPwd {
    log.Println("Wrong Details Provided")
    return
  }

  token, err := handler.AuthenticationService.GenerateJWT(account.ID)
  if err != nil {
    log.Println(err)
    return
  }

  response := struct { Id int64; Email string; Token string } { account.ID, account.Email, token }
  rw.Header().Set("Content-Type", "application/json")
	json.NewEncoder(rw).Encode(response)
}