package routes

import (
	"errors"
	"project-book-reservation/database"
	"project-book-reservation/models"
	"time"

	"github.com/gofiber/fiber/v2"
)

type Book struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	BookID    int       `json:"book_id"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"created_at"`
}

func CreateResponseBook(book models.Book) Book {
	return Book{ID: book.ID, Name: book.Name, BookID: book.BookID, State: book.State, CreatedAt: book.CreatedAt}
}

func CreateBook(c *fiber.Ctx) error {
	var Book models.Book

	if err := c.BodyParser(&Book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.Db.Create(&Book)
	responseBook := CreateResponseBook(Book)
	return c.Status(200).JSON(responseBook)
}

func GetBooks(c *fiber.Ctx) error {
	Books := []models.Book{}
	database.Database.Db.Find(&Books)
	responseBooks := []Book{}
	for _, Book := range Books {
		responseBook := CreateResponseBook(Book)
		responseBooks = append(responseBooks, responseBook)
	}

	return c.Status(200).JSON(responseBooks)
}

func findBook(id int, Book *models.Book) error {
	database.Database.Db.Find(&Book, "id = ?", id)
	if Book.ID == 0 {
		return errors.New("Book does not exist")
	}
	return nil
}

func GetBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var Book models.Book

	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is correct")
	}

	if err := findBook(id, &Book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	responseBook := CreateResponseBook(Book)

	return c.Status(200).JSON(responseBook)
}

func UpdateBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var Book models.Book

	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is correct")
	}

	err = findBook(id, &Book)

	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	type UpdateBook struct {
		Name   string `json:"name"`
		State  string `json:"state"`
		BookID int    `json:"book_id"`
	}

	var updateData UpdateBook

	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(500).JSON(err.Error())
	}

	Book.Name = updateData.Name
	Book.State = updateData.State
	Book.BookID = updateData.BookID

	database.Database.Db.Save(&Book)

	responseBook := CreateResponseBook(Book)

	return c.Status(200).JSON(responseBook)
}

func DeleteBook(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")

	var Book models.Book

	if err != nil {
		return c.Status(400).JSON("Please make sure that :id is correct")
	}

	if err := findBook(id, &Book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	if err := database.Database.Db.Delete(&Book).Error; err != nil {
		return c.Status(400).JSON(err.Error())
	}

	return c.Status(200).SendString("Successfully Delete Book")
}
