package controller

import (
	db "fiber-go/config"
	"fiber-go/middleware"
	"fiber-go/model"
	"fmt"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreatePaymentController(c *fiber.Ctx) error {

	var data map[string]string
	paymentError := c.BodyParser(&data)
	if paymentError != nil {
		log.Fatalf("payment error in post request %v", paymentError)
	}
	if data["name"] == "" || data["type"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Payment Name and Type is required",
			"error":   map[string]interface{}{},
		})
	}

	var paymentTypes model.PaymentType
	db.DB.Where("name", data["type"]).First(&paymentTypes)
	payment := model.Payment{
		Name:          data["name"],
		Type:          data["type"],
		PaymentTypeID: int(paymentTypes.ID),
		Logo:          data["logo"],
	}
	db.DB.Create(&payment)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    payment,
	})
}
func ListPaymentController(c *fiber.Ctx) error {
	//Token authenticate
	headerToken := c.Get("Authorization")
	err := middleware.Auth(headerToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))
	var count int64
	var payment []model.PaymentDto
	db.DB.Select("id ,name,type,payment_type_id,logo,created_at,updated_at").Limit(limit).Offset(skip).Find(&payment).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	categoriesData := map[string]interface{}{
		"payments": payment,
		"meta":     metaMap,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    categoriesData,
	})

}

func FindPaymentController(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")
	var payment model.Payment

	//Token authenticate
	headerToken := c.Get("Authorization")
	err := middleware.Auth(headerToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Where("id=?", paymentId).First(&payment)

	if payment.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Payment Not Found",
			"error":   map[string]interface{}{},
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    payment,
	})
}

func DeletePaymentController(c *fiber.Ctx) error {
	paymentId := c.Params("paymentId")
	var payment model.Payment

	//Token authenticate
	headerToken := c.Get("Authorization")
	err := middleware.Auth(headerToken)
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Unauthorized",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.First(&payment, paymentId)
	if payment.Name == "" {
		return c.Status(500).JSON(fiber.Map{
			"success": false,
			"Message": "No payment found against this payment id",
		})
	}

	result := db.DB.Delete(&payment)
	if result.RowsAffected == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "payment removing failed",
		})
	}

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
	})
}

func UpdatePaymentController(c *fiber.Ctx) error {
	id := c.Params("id")
	var totalPayment model.Payment
	var payment model.Payment
	var paymentTypeID int
	var updatePaymentData model.Payment

	fmt.Println("-----------------------------------")
	fmt.Println("---------------Params payment id--------------------", id)
	fmt.Println("-----------------------------------")

	db.DB.Find(&totalPayment)

	fmt.Println("-----------------------------------")
	fmt.Println("---------------All payments--------------------", totalPayment)
	fmt.Println("-----------------------------------")

	db.DB.Find(&payment, "id = ?", id)

	c.BodyParser(&updatePaymentData)
	if updatePaymentData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Payment name is required",
			"error":   map[string]interface{}{},
		})
	}

	if updatePaymentData.Type == "CASH" {
		paymentTypeID = 1
	}
	if updatePaymentData.Type == "E-WALLET" {
		paymentTypeID = 2
	}
	if updatePaymentData.Type == "EDC" {
		paymentTypeID = 3
	}

	fmt.Println(paymentTypeID)

	payment.Name = updatePaymentData.Name
	payment.Type = updatePaymentData.Type
	payment.PaymentTypeID = paymentTypeID
	payment.Logo = updatePaymentData.Logo

	db.DB.Save(&payment)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    payment,
	})

}
