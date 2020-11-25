package auth

import (
	session "gobanks/config"

	"github.com/gofiber/fiber/v2"
)

func IfAuth(c *fiber.Ctx) error {
	sessions := session.Sessions()
	store := sessions.Get(c)

	email := store.Get("email")
	authenticated := store.Get("authenticate")

	if email != nil && authenticated != true {
		//path := c.Path()
		//fmt.Println(path)
		return c.Redirect("/")
	}
	return c.Next()
}
