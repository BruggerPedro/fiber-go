package model

import "time"

type Product struct {
	ID               int       `json:"id" gorm:"type:int;primaryKey;autoIncrement"`
	Sku              string    `json:"sku"`
	Name             string    `json:"name"`
	Stock            int       `json:"stock"`
	Price            int       `json:"price"`
	Image            string    `json:"image"`
	TotalFinalPrice  int       `json:"totalFinalPrice"`
	TotalNormalPrice int       `json:"totalNormalPrice"`
	CreatedAt        time.Time `json:"createdAt"`
	UpdatedAt        time.Time `json:"updatedAt"`
	CategoryID       int       `json:"categoryId"`
	DiscountID       int       `json:"discountId"`
}

type ProductResult struct {
	Id       int      `json:"productId" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku      string   `json:"sku"`
	Name     string   `json:"name"`
	Stock    int      `json:"stock"`
	Price    int      `json:"price"`
	Image    string   `json:"image"`
	Category Category `json:"category"`
	Discount Discount `json:"discount"`
}

type ProductsDto struct {
	Products     Product
	CategoriesId string `json:"categories_Id"`
}
type ProductDiscount struct {
	Id         int      `json:"id" gorm:"type:INT(10) UNSIGNED NOT NULL AUTO_INCREMENT;primaryKey"`
	Sku        string   `json:"sku"`
	Name       string   `json:"name"`
	Stock      int      `json:"stock"`
	Price      int      `json:"price"`
	Image      string   `json:"image"`
	CategoryId int      `json:"categoryId"`
	Discount   Discount `json:"discount"`
}
