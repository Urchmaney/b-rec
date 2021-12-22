package models

import (
  "time"
)

type Account struct {
  ID uint64
  Owner string
  StartingDebt int
  CreatedAt time.Time
}

type User struct {
  ID int64
  Name string
  AccountId uint64
  CreatedAt time.Time
}

type Bill struct {
  ID uint
  Amount float64
  Month uint8
  CreatedAt time.Time
}

type Payment struct {
  ID uint
  Amount uint32
  UserId uint
  Month uint8
  CreatedAt time.Time
}

