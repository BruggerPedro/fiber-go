package controller

import (
	db "fiber-go/config"
	"fiber-go/middleware"
	"fiber-go/model"
	"log"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func CreateCategoryController(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("category error in post request %v", err)
	}

	if data["name"] == "" {
		c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
			"error":   map[string]interface{}{},
		})
	}

	category := model.Category{
		Name: data["name"],
	}

	db.DB.Create(&category)

	c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    category,
	})

	return nil
}

func FindCategoryController(c *fiber.Ctx) error {
	id := c.Params("id")
	var category model.Category

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

	db.DB.Select("id ,name").Where("id=?", id).First(&category)
	categoryData := make(map[string]interface{})
	categoryData["id"] = category.ID
	categoryData["name"] = category.Name

	if category.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"Message": "No category found",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    categoryData,
	})
}

func ListCategoryController(c *fiber.Ctx) error {
	var count int64
	var category []model.Category

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

	db.DB.Select("id ,name").Limit(limit).Offset(skip).Find(&category).Count(&count)
	metaMap := map[string]interface{}{
		"total": count,
		"limit": limit,
		"skip":  skip,
	}
	categoriesData := map[string]interface{}{
		"categories": category,
		"meta":       metaMap,
	}

	return c.JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    categoriesData,
	})

}

func UpdateCategoryController(c *fiber.Ctx) error {
	id := c.Params("id")
	var category model.Category
	var updateCashierData model.Category

	db.DB.Find(&category, "id = ?", id)

	if category.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Category not exist against this id",
		})
	}

	c.BodyParser(&updateCashierData)
	if updateCashierData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Category name is required",
			"error":   map[string]interface{}{},
		})
	}

	category.Name = updateCashierData.Name
	db.DB.Save(&category)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
		"data":    category,
	})

}

func DeleteCategoryController(c *fiber.Ctx) error {
	id := c.Params("id")
	var category model.Category

	db.DB.Where("id=?", id).First(&category)

	if category.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "category not found",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Where("id = ?", id).Delete(&category)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
	})
}
