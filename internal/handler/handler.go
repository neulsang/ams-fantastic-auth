package handler

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

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

func Users(c *fiber.Ctx) error {
	fmt.Println("Users")
	return nil
}

func User(ctx *fiber.Ctx) error {
	log.Println("User")
	return nil
}

func PutUser(ctx *fiber.Ctx) error {
	log.Println("PutUser")
	return nil
}

func DeleteUser(c *fiber.Ctx) error {
	log.Println("DeleteUser")
	return nil
}
