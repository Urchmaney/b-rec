package main

import (
  "net/http"
  "log"
  "database/sql"
  "flag"
  _ "github.com/go-sql-driver/mysql"
  "b-rec/pkg/models/mysql"
  "github.com/go-chi/chi/v5"
  "b-rec/pkg/authenticator"
  "b-rec/pkg/handlers"
)

type application struct {
  user_dao mysql.UserDAO
}

func main() {
  dsn := flag.String("dsn", "root:root@tcp(db)/b_rec?parseTime=true", "MySQL data source name")
  flag.Parse()

  db, err := OpenDB(*dsn)
  if err != nil {
    log.Println(err)
    return
  }

  // app := application{  }
 
  user_dao := mysql.UserDAO{ DB: db }
  user_handler := handlers.UserHandler{ UserService: user_dao }

  account_dao := mysql.AccountDAO{ DB: db }
  authenticator := authenticator.AuthenticationService{}
  accounts_handler := handlers.AccountHandler { AccountService: account_dao, AuthenticationService: authenticator }

  bill_service := mysql.BillService{ DB: db }
  bill_handler := handlers.BillHandler{ BillService: bill_service }

  payment_service := mysql.PaymentService{ DB: db }
  payment_handler := handlers.PaymentHandler{ PaymentService: payment_service }

  r := chi.NewRouter()
  r.Route("/users", func(r chi.Router) {
    r.Get("/", user_handler.GetAllUsers)
    r.Post("/", user_handler.CreateUser)
  })

  r.Route("/accounts", func(r chi.Router) {
    r.Post("/signup", accounts_handler.SignUp)
    r.Post("/login", accounts_handler.Login)
  })

  r.Route("/bills", func(r chi.Router) {
    r.Post("/", bill_handler.CreateBill)
  })

  r.Route("/payments", func(r chi.Router) {
    r.Post("/", payment_handler.AddPayments)
  })
  // mux := http.NewServeMux()
  // mux.HandleFunc("/users", )
  
  log.Println("starting server at port :3000")
  err = http.ListenAndServe(":3000", r)
  log.Println(err)
}

func OpenDB(dsn string)(*sql.DB, error) {
  db, err := sql.Open("mysql", dsn)
  if err != nil {
    return nil, err
  }
  if err = db.Ping(); err != nil {
    return nil, err
  }
  return db, nil
}