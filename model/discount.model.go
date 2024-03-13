package model

import "time"

type Discount struct {
	ID              int       `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Quantity        int       `json:"quantity"`
	Type            string    `json:"type"`
	Result          int       `json:"result"`
	ExpiredAt       int       `json:"expiredAt"`
	ExpiredAtFormat string    `json:"expiredAtFormat"`
	StringFormat    string    `json:"stringFormat"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

type DiscountDto struct {
	Quantity  int    `json:"qty"`
	Types     string `json:"type"`
	Result    int    `json:"result"`
	ExpiredAt int    `json:"expiredAt"`
}
