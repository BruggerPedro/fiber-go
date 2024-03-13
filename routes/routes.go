package routes

import (
	"fiber-go/controller"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/id/login", controller.Login)
	app.Get("/cashiers/id/logout", controller.Logout)
	app.Post("/cashiers/id/password", controller.Password)

	// cashier routes
	app.Post("/cashiers", controller.CreateCashierController)
	app.Get("/cashiers", controller.ListCashierController)
	app.Get("/cashiers/:id", controller.FindCashierController)
	app.Patch("/cashiers/:id", controller.UpdateCashierController)
	app.Delete("/cashiers/:id", controller.DeleteCashierController)

	//Category routes
	app.Get("/categories", controller.ListCategoryController)
	app.Get("/categories/:categoryId", controller.FindCategoryController)
	app.Post("/categories", controller.CreateCategoryController)
	app.Delete("/categories/:categoryId", controller.DeleteCategoryController)
	app.Put("/categories/:categoryId", controller.UpdateCategoryController)

	//Products routes
	app.Get("/products", controller.ListProductsController)
	app.Get("/products/:productId", controller.FindProductController)
	app.Post("/products", controller.CreateProductController)
	app.Delete("/products/:productId", controller.DeleteProductController)
	app.Put("/products/:productId", controller.UpdateProductController)

	//Payment routes
	app.Get("/payments", controller.ListPaymentController)
	app.Get("/payments/:paymentId", controller.FindPaymentController)
	app.Post("/payments", controller.CreatePaymentController)
	app.Delete("/payments/:paymentId", controller.DeletePaymentController)
	app.Put("/payments/:paymentId", controller.UpdatePaymentController)

	//Order routes
	app.Get("/orders", controller.ListOrderController)
	app.Get("/orders/:orderId", controller.FindOrderController)
	app.Post("/orders", controller.CreateOrderController)
	app.Post("/orders/subtotal", controller.SubTotalOrderController)
	app.Get("/orders/:orderId/download", controller.DownloadOrderController)
	app.Get("/orders/:orderId/check-download", controller.CheckOrderController)

	//reports
	app.Get("/revenues", controller.GetRevenues)
	app.Get("/solds", controller.GetSolds)
}
