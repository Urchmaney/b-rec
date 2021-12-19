package main

import (
  "net/http"
  "log"
  "database/sql"
  "flag"
  _ "github.com/go-sql-driver/mysql"
  "b-rec/pkg/models/mysql"
  "github.com/go-chi/chi/v5"
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
  user_handler := UserHandler{ userService: user_dao }

  r := chi.NewRouter()
  r.Route("/users", func(r chi.Router) {
    r.Get("/", user_handler.home)
  })

  r.Route("/accounts", func(r chi.Router) {
    r.Get("/", user_handler.home)
    r.Get("/{account_slug:[1-z-]+}", func (w http.ResponseWriter, r *http.Request) {
      slug := chi.URLParam(r, "account_slug")
      w.Write([]byte(slug))
    })
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