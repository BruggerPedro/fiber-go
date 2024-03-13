package model

import "time"

type Category struct {
	ID        int       `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
