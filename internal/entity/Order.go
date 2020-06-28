package entity

import (
    "time"
)

type Order struct {
    ID int `json:"id" pg:"id,pk"`
    Items []*OrderItem `json:"items"`
    CreatedAt time.Time `json:"created_at" pg:"created_at" pg:"default:now()"`
}
