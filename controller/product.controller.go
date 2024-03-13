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

func CreateProductController(c *fiber.Ctx) error {
	var data model.ProductDiscount
	var p []model.Product

	err := c.BodyParser(&data)
	if err != nil {
		log.Fatalf("Product error in post request %v", err)
	}

	db.DB.Find(&p)

	discount := model.Discount{
		Quantity:  data.Discount.Quantity,
		Type:      data.Discount.Type,
		Result:    data.Discount.Result,
		ExpiredAt: data.Discount.ExpiredAt,
	}
	db.DB.Create(&discount)

	product := model.Product{
		Name:       data.Name,
		Image:      data.Image,
		CategoryID: data.CategoryId,
		DiscountID: discount.ID,
		Price:      data.Price,
		Stock:      data.Stock,
	}
	db.DB.Create(&product)

	db.DB.Table("products").Where("id = ?", product.ID).Update("sku", "SK00"+strconv.Itoa(product.ID))

	fmt.Println("--------------------------------------->")
	fmt.Println("------------Product Creation Done----------->", product.ID)
	fmt.Println("--------------------------------------->")

	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    product,
	}
	return (c.JSON(Response))

}

func FindProductController(c *fiber.Ctx) error {
	id := c.Params("id")
	var products []model.Product
	var category model.Category
	var discount model.Discount

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

	productsRes := make([]*model.ProductResult, 0)

	db.DB.Where("id = ? ", id).Find(&products)

	for i := 0; i < len(products); i++ {
		db.DB.Where("id = ?", products[i].CategoryID).Find(&category)

		db.DB.Where("id = ?", products[i].DiscountID).Find(&discount)

		productsRes = append(productsRes,
			&model.ProductResult{
				Id:       products[i].ID,
				Sku:      products[i].Sku,
				Name:     products[i].Name,
				Stock:    products[i].Stock,
				Price:    products[i].Price,
				Image:    products[i].Image,
				Category: category,
				Discount: discount,
			},
		)
	}

	Response := map[string]interface{}{
		"success": true,
		"message": "Success",
		"data":    productsRes,
	}
	return (c.JSON(Response))
}

func ListProductsController(c *fiber.Ctx) error {
	id := c.Query("id")
	productName := c.Query("name")
	limit := c.Query("limit")
	skip := c.Query("skip")
	var products []model.Product

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

	intLimit, _ := strconv.Atoi(limit)
	intSkip, _ := strconv.Atoi(skip)

	productsRes := make([]*model.ProductResult, 0)

	if productName == "" {
		var count int64
		db.DB.Where("category_id = ?", id).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)

		var category model.Category
		var discount model.Discount
		for i := 0; i < len(products); i++ {

			db.DB.Table("categories").Where("id = ?", products[i].CategoryID).Find(&category)

			db.DB.Where("id = ?", products[i].DiscountID).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
			count = int64(len(products))
			//productsRes =
			productsRes = append(productsRes,
				&model.ProductResult{
					Id:       products[i].ID,
					Sku:      products[i].Sku,
					Name:     products[i].Name,
					Stock:    products[i].Stock,
					Price:    products[i].Price,
					Image:    products[i].Image,
					Category: category,
					Discount: discount,
				},
			)
		}

		meta := map[string]interface{}{
			"total": count,
			"limit": limit,
			"skip":  skip,
		}

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Success",
			"data": map[string]interface{}{
				"products": productsRes,
				"meta":     meta,
			},
		})
	} else {

		var count int64
		if id != "" {
			db.DB.Where("category_id = ? AND name= ?", id, productName).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		} else {
			db.DB.Where(" name= ?", productName).Limit(intLimit).Offset(intSkip).Find(&products).Count(&count)
		}
		var category model.Category
		var discount model.Discount
		for i := 0; i < len(products); i++ {
			db.DB.Where("id = ?", products[i].CategoryID).Find(&category)
			db.DB.Where("id = ?", products[i].DiscountID).Limit(intLimit).Offset(intSkip).Find(&discount).Count(&count)
			count = int64(len(products))
			productsRes = append(productsRes,
				&model.ProductResult{
					Id:       products[i].ID,
					Sku:      products[i].Sku,
					Name:     products[i].Name,
					Stock:    products[i].Stock,
					Price:    products[i].Price,
					Image:    products[i].Image,
					Category: category,
					Discount: discount,
				},
			)
		}

		meta := map[string]interface{}{
			"total": count,
			"limit": limit,
			"skip":  skip,
		}

		return c.Status(200).JSON(fiber.Map{
			"success": true,
			"message": "Success",
			"data": map[string]interface{}{
				"products": productsRes,
				"meta":     meta,
			},
		})
	}
}

func UpdateProductController(c *fiber.Ctx) error {
	productId := c.Params("productId")
	var product model.Product

	db.DB.Find(&product, "id = ?", productId)

	if product.Name == "" {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
			"error":   map[string]interface{}{},
		})
	}

	var updateProductData model.Product
	c.BodyParser(&updateProductData)

	if updateProductData.Name == "" {
		return c.Status(400).JSON(fiber.Map{
			"success": false,
			"message": "Product name is required",
			"error":   map[string]interface{}{},
		})
	}

	product.Name = updateProductData.Name
	product.CategoryID = updateProductData.CategoryID
	product.Image = updateProductData.Image
	product.Price = updateProductData.Price
	product.Stock = updateProductData.Stock

	db.DB.Save(&product)
	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"message": "Success",
		"data":    product,
	})
}

func DeleteProductController(c *fiber.Ctx) error {
	id := c.Params("id")
	var product model.Product

	db.DB.First(&product, id)
	if product.ID == 0 {
		return c.Status(404).JSON(fiber.Map{
			"success": false,
			"message": "Product Not Found",
			"error":   map[string]interface{}{},
		})
	}

	db.DB.Delete(&product)

	return c.Status(200).JSON(fiber.Map{
		"success": true,
		"Message": "Success",
	})
}
