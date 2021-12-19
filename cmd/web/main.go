package main

import (
  "net/http"
  "log"
  "database/sql"
  "flag"
  _ "github.com/go-sql-driver/mysql"
  "b-rec/pkg/models/mysql"
)

type application struct {
  user_dao mysql.UserDAO
}

func main() {
  dsn := flag.String("dsn", "root:root@tcp(db)/b_rec?parseTime=true", "MySQL data source name")
  flag.Parse()

  _, err := OpenDB(*dsn)
  if err != nil {
    log.Println(err)
    return
  }

  // app := application{ user_dao: mysql.UserDAO{ DB: db } }

  mux := http.NewServeMux()
  mux.HandleFunc("/", home)
  
  log.Println("starting server at port :3000")
  err = http.ListenAndServe(":3000", mux)
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