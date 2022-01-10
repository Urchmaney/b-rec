package models

import (
  "time"
)

type Account struct {
  ID int64
  OwnerFullName string
  Password string
  Email string
  StartingDebt int
  CreatedAt time.Time
}

type User struct {
  ID int64
  Name string
  AccountId int64
  CreatedAt time.Time
}

type Bill struct {
  ID int64
  AccountId int64
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

