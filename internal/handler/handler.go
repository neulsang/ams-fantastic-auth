package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

// User struct
type User struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func Singup(c *fiber.Ctx) error {
	log.Println("Singup")
	return nil
}

func Signin(c *fiber.Ctx) error {
	log.Println("Signin")
	return nil
}

func Signout(c *fiber.Ctx) error {
	fmt.Println("Signout")
	return nil
}

func GetUsers(c *fiber.Ctx) error {
	fmt.Println("Users")
	return nil
}

// @Summary Get user
// @Description Get user's info
// @Accept json
// @Produce json
// @Param id path string true "id of the user"
// @Success 200 {object} handler.User
// @Router /users/{id} [get]
func GetUser(c *fiber.Ctx) error {
	log.Println("User")
	u := User{
		ID:   "test",
		Name: "tester",
		Age:  99,
	}
	return c.JSON(u)
}

func PutUser(c *fiber.Ctx) error {
	log.Println("PutUser")
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	log.Println("DeleteUser")
	return nil
}
