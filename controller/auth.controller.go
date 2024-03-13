package controller

import (
	db "fiber-go/config"
	"fiber-go/model"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

func Login(c *fiber.Ctx) error {
	id := c.Params("id")
	var data map[string]string
	var cashier model.Cashier

	err := c.BodyParser(&data)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password is required",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.First(&cashier, id)

	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
			"error":   map[string]interface{}{},
		})
	}

	if cashier.Password != data["password"] {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Password not match",
			"error":   map[string]interface{}{},
		})
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"Issuer":    strconv.Itoa(int(cashier.ID)),
		"ExpiresAt": time.Now().Add(time.Hour * 24).Unix(), // 1 day
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"message": "Token Expired or Invalid",
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["token"] = tokenString

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Login successful",
		"data":    cashierData,
	})
}

func Logout(c *fiber.Ctx) error {
	id := c.Params("id")
	var data map[string]string
	if err := c.BodyParser(&data); err != nil {
		return err
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"Message": "Password is required",
		})
	}

	var cashier model.Cashier
	db.DB.Where("id = ?", id).First(&cashier)

	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "Cashier Not found",
		})
	}

	if cashier.Password != data["password"] {
		return c.Status(401).JSON(fiber.Map{
			"success": false,
			"Message": "Password Not Match",
		})
	}
	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}

	c.Cookie(&cookie)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "success",
	})
}

func Password(c *fiber.Ctx) error {
	cashierId := c.Params("id")
	var cashier model.Cashier
	db.DB.Select("id,name,password").Where("id=?", cashierId).First(&cashier)

	if cashier.Name == "" || cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier Not Found",
			"error":   map[string]interface{}{},
		})
	}

	cashierData := make(map[string]interface{})
	cashierData["password"] = cashier.Password

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    cashierData,
	})
}
