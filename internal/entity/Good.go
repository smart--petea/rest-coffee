package entity

import (
    "time"
)

type Good struct {
    ID int `json:"id" pg:"id,pk"`
    CreatedAt time.Time `json:"created_at" pg:"created_at" pg:"default:now()"`
    IsDeleted bool `json:"is_deleted" pg:"is_deleted", pg:"default:false"`
    Name string `json:"name" pg:"name"`
}
