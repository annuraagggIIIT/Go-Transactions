package routes

import (
	controllers "example/v2/Controllers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/cashiers/:cashierId/login", controllers.Login)
	app.Post("/cashiers/:cashierId/logout", controllers.Logout)
	app.Get("/cashiers/:cashierId/passcode", controllers.Passcode)
	//Cashier routes
	app.Get("/cashiers", controllers.CashiersList)
	app.Get("/cashiers/:cashierId", controllers.GetCashierDetails)
	app.Post("/cashiers", controllers.CreateCashier)
	app.Delete("/cashiers/:cashierId", controllers.DeleteCashier)
	app.Put("/cashiers/:cashierId", controllers.UpdateCashier)

	//Category routes
	app.Get("/categories", controllers.CategoryList)
	app.Get("/categories/:categoryId", controllers.GetCategoryDetails)
	app.Post("/categories", controllers.CreateCategory)
	app.Delete("/categories/:categoryId", controllers.DeleteCategory)
	app.Put("/categories/:categoryId", controllers.UpdateCategory)

	
	//Payment routes
	app.Get("/payments", controllers.PaymentList)
	app.Get("/payments/:paymentId", controllers.GetPaymentDetails)
	app.Post("/payments", controllers.CreatePayment)
	app.Delete("/payments/:paymentId", controllers.DeletePayment)
	app.Put("/payments/:paymentId", controllers.UpdatePayment)

	

	//reports
	app.Get("/revenues", controllers.GetRevenues)
	app.Get("/solds", controllers.GetSolds)

}
