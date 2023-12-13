package main

import (
	db "example/v2/config"
	"example/v2/routes"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main(){
	fmt.Println("deve code app running...")
	//db connection
	db.Connect()

	app := fiber.New()
	app.Use(cors.New())
	//routing
	routes.Setup(app)
	app.Listen(":3030")

}
