package controller

import (
	db "fiber-go/config"
	"fiber-go/middleware"
	"fiber-go/model"
	"log"
	"strconv"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetRevenues(c *fiber.Ctx) error {
	order := []model.Order{}

	// Token authenticate
	headerToken := c.Get("Authorization")
	err := middleware.Auth(headerToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}


	db.DB.Find(&order)
	
	TotalRevenues := make([]*model.RevenueResponse, 0)

	Resp1 := model.RevenueResponse{}
	Resp2 := model.RevenueResponse{}
	Resp3 := model.RevenueResponse{}

	sum1 := 0
	sum2 := 0
	sum3 := 0
	for _, v := range order {
		if v.PaymentTypesId == 1 {
			payment := model.Payment{}
			paymentTypes := model.PaymentType{}

			db.DB.Where("id=?", 1).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 1).First(&payment)

			sum1 += v.TotalPaid
			Resp1.Name = paymentTypes.Name
			Resp1.Logo = payment.Logo
			Resp1.TotalAmount = sum1
			Resp1.PaymentTypeId = v.PaymentTypesId
		}

		if v.PaymentTypesId == 2 {

			payment := model.Payment{}
			paymentTypes := model.PaymentType{}

			db.DB.Where("id=?", 2).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 2).First(&payment)

			sum2 += v.TotalPaid
			Resp2.Name = paymentTypes.Name
			Resp2.Logo = payment.Logo
			Resp2.TotalAmount = sum2
			Resp2.PaymentTypeId = v.PaymentTypesId
		}
		if v.PaymentTypesId == 3 {

			payment := model.Payment{}
			paymentTypes := model.PaymentType{}

			db.DB.Where("id=?", 3).First(&paymentTypes)
			db.DB.Where("payment_type_id=?", 3).First(&payment)

			sum3 += v.TotalPaid
			Resp3.Name = paymentTypes.Name
			Resp3.Logo = payment.Logo
			Resp3.TotalAmount = sum2
			Resp3.PaymentTypeId = v.PaymentTypesId
		}
	}
	TotalRevenues = append(TotalRevenues, &Resp1)
	TotalRevenues = append(TotalRevenues, &Resp2)
	TotalRevenues = append(TotalRevenues, &Resp3)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": map[string]interface{}{
			"totalRevenue": sum1 + sum2 + sum3,
			"paymentTypes": TotalRevenues,
		},
	})
}

type Sold struct {
	ProductId   string `json:"productId"`
	Quantities  string `json:"quantities"`
	TotalAmount int    `json:"totalAmount"`
}

func GetSolds(c *fiber.Ctx) error {
	orders := []model.Order{}
	db.DB.Find(&orders)

	TotalSolds := make([]*model.SoldResponse, 0)
	TotalSoldsFinal := make([]*model.SoldResponse, 0)

	for _, v := range orders {
		quantities := strings.Split(v.Quantities, ",")
		quantities = quantities[1:]

		products := strings.Split(v.ProductId, ",")
		products = products[1:]

		for i := 0; i < len(products); i++ {
			prods := model.Product{}
			p, err := strconv.Atoi(products[i])
			q, errq := strconv.Atoi(quantities[i])

			if err != nil {
				log.Fatalf("-> %v", err)
			}
			if errq != nil {
				log.Fatalf("-> %v", errq)
			}

			db.DB.Where("id", p).Find(&prods)
			TotalSolds = append(TotalSolds, &model.SoldResponse{
				Name:        prods.Name,
				ProductId:   p,
				TotalQty:    q,
				TotalAmount: q * prods.Price,
			})
		}

	}

	duplicates := []int{}
	for _, v := range TotalSolds {
		if !contains(duplicates, v.ProductId) {
			duplicates = append(duplicates, v.ProductId)
		}
	}

	quantityArray := []int{}
	for _, v := range duplicates {
		qty := 0
		for _, x := range TotalSolds {
			if v == x.ProductId {
				qty = qty + x.TotalQty
			}
		}
		quantityArray = append(quantityArray, qty)

	}

	for i := 0; i < len(duplicates); i++ {
		prods := model.Product{}

		db.DB.Where("id", duplicates[i]).Find(&prods)

		TotalSoldsFinal = append(TotalSoldsFinal, &model.SoldResponse{
			Name:        prods.Name,
			TotalQty:    quantityArray[i],
			TotalAmount: quantityArray[i] * prods.Price,
			ProductId:   duplicates[i],
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data": map[string]interface{}{
			"orderProducts": TotalSoldsFinal,
		},
	})
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
