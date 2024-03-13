package controller

import (
	"fiber-go/model"
	"strconv"
	"time"

	db "fiber-go/config"

	"github.com/gofiber/fiber/v2"
)

func CreateCashierController(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Invalid data",
		})
	}

	if data["name"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	if data["password"] == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier password is required",
		})
	}

	cashier := model.Cashier{
		Name:      data["name"],
		Password:  data["password"],
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	db.DB.Create(&cashier)

	return c.Status(201).JSON(fiber.Map{
		"success": true,
		"message": "Cashier created successfully",
		"data":    cashier,
	})
}

func ListCashierController(c *fiber.Ctx) error {
	var cashiers []model.Cashier
	var count int64

	limit, _ := strconv.Atoi(c.Query("limit"))
	skip, _ := strconv.Atoi(c.Query("skip"))

	db.DB.Select("*").Limit(limit).Offset(skip).Find(&cashiers).Count(&count)

	return c.JSON(fiber.Map{
		"success": true,
		"data":    cashiers,
	})
}

func FindCashierController(c *fiber.Ctx) error {
	id := c.Params("id")

	var cashier model.Cashier

	db.DB.Select("id,name,created_at,updated_at").Where("id=?", id).First(&cashier)

	cashierData := make(map[string]interface{})
	cashierData["id"] = cashier.ID
	cashierData["name"] = cashier.Name
	cashierData["password"] = cashier.Password
	cashierData["created_at"] = cashier.CreatedAt
	cashierData["updated_at"] = cashier.UpdatedAt

	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
			"error":   map[string]interface{}{},
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"message": "success",
		"data":    cashier,
	})
}

func UpdateCashierController(c *fiber.Ctx) error {
	id := c.Params("id")

	var cashier model.Cashier

	db.DB.Where("id=?", id).First(&cashier)

	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	var updateCashier model.Cashier

	if err := c.BodyParser(&updateCashier); err != nil {
		return err
	}

	if updateCashier.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Cashier name is required",
		})
	}

	cashier.Name = updateCashier.Name

	db.DB.Save(&cashier)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Cashier updated successfully",
		"data":    cashier,
	})
}

func DeleteCashierController(c *fiber.Ctx) error {
	id := c.Params("id")

	var cashier model.Cashier

	db.DB.Where("id=?", id).First(&cashier)

	if cashier.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Cashier not found",
		})
	}

	db.DB.Where("id=?", id).Delete(&cashier)

	return c.JSON(fiber.Map{
		"success": true,
		"message": "Cashier deleted successfully",
	})
}
