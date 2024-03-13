package model

import "time"

type Payment struct {
	ID            int       `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	PaymentTypeID int       `json:"paymentTypeId"`
	Logo          string    `json:"logo"`
	CreatedAt     time.Time `json:"createdAt"`
	UpdatedAt     time.Time `json:"updatedAt"`
}

type PaymentDto struct {
	Id            uint   `json:"paymentId"`
	Name          string `json:"name"`
	Type          string `json:"type"`
	PaymentTypeId int    `json:"payment_type_id"`
	Logo          string `json:"logo"`
}

type PaymentType struct {
	ID        int       `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
}
