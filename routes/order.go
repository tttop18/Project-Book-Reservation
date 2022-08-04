package routes

import (
	"errors"
	"project-book-reservation/database"
	"project-book-reservation/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Order struct {
	ID        uint      `json:"id"`
	User      User      `json:"user"`
	Book      Book      `json:"book"`
	CreatedAt time.Time `json:"order_date"`
}

func CreateResponseOrder(order models.Order, user User, book Book) Order {
	return Order{ID: order.ID, User: user, Book: book, CreatedAt: order.CreatedAt}
}

func CreateOrder(c *fiber.Ctx) error {
	var order models.Order

	if err := c.BodyParser(&order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User

	if err := findUser(order.UserRefer, &user); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var book models.Book

	if err := findBook(order.BookRefer, &book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&order)

	responseUser := CreateResponseUser(user)
	responseBook := CreateResponseBook(book)
	responseOrder := CreateResponseOrder(order, responseUser, responseBook)

	return c.Status(200).JSON(responseOrder)
}

func GetOrders(c *fiber.Ctx) error {
	orders := []models.Order{}
	database.Database.Db.Find(&orders)
	responseOrders := []Order{}

	for _, order := range orders {
		var user models.User
		var book models.Book
		database.Database.Db.Find(&user, "id = ?", order.UserRefer)
		database.Database.Db.Find(&book, "id = ?", order.BookRefer)
		responseOrder := CreateResponseOrder(order, CreateResponseUser(user), CreateResponseBook(book))
		responseOrders = append(responseOrders, responseOrder)
	}

	return c.Status(200).JSON(responseOrders)
}

func FindOrder(id int, order *models.Order) error {
	database.Database.Db.Find(&order, "id = ?", id)
	if order.ID == 0 {
		return errors.New("order does not exist")
	}
	return nil
}

func GetOrder(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var order models.Order

	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is correct")
	}

	if err := FindOrder(id, &order); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	var user models.User
	var book models.Book

	database.Database.Db.First(&user, order.UserRefer)
	database.Database.Db.First(&book, order.BookRefer)
	responseUser := CreateResponseUser(user)
	responseBook := CreateResponseBook(book)

	responseOrder := CreateResponseOrder(order, responseUser, responseBook)

	return c.Status(200).JSON(responseOrder)

}
