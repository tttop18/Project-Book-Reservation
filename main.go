package main

import (
	"log"

	"project-book-reservation/database"
	"project-book-reservation/routes"

	"github.com/gofiber/fiber/v2"
)

func welcome(c *fiber.Ctx) error {
	return c.SendString("Welcome world")
}

func setupRoutes(app *fiber.App) {
	app.Get("/api", welcome)
	// User
	app.Post("/api/users", routes.CreateUser)
	app.Get("/api/users", routes.GetUsers)
	app.Get("/api/users/:id", routes.GetUser)
	app.Put("/api/users/:id", routes.UpdateUser)
	app.Delete("/api/users/:id", routes.DeleteUser)
	// Book
	app.Post("/api/books", routes.CreateBook)
	app.Get("/api/books", routes.GetBooks)
	app.Get("/api/books/:id", routes.GetBook)
	app.Put("/api/books/:id", routes.UpdateBook)
	app.Delete("/api/books/:id", routes.DeleteBook)
	// Order
	app.Post("/api/orders", routes.CreateOrder)
	app.Get("/api/orders", routes.GetOrders)
	app.Get("/api/orders/:id", routes.GetOrder)
}

func main() {
	database.ConnectDb()

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":3000"))

}
