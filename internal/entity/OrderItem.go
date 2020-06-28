package entity

import (
    "time"
)

type OrderItem struct {
    ID int `json:"id" pg:"id,pk"`
    Quantity float32 `json:"quantity", pg:"quantity"`
    CreatedAt time.Time `json:"created_at" pg:"created_at" pg:"default:now()"`
    IsDeleted bool `json:"is_deleted" pg:"is_deleted", pg:"default:false"`
    GoodId int `json:"good_id" pg:"good_id,fk:Good"`
    OrderId int `json:"order_id" pg:"order_id,fk:Order"`
}
