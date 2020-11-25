package auth

import (
	session "gobanks/config"

	"github.com/gofiber/fiber/v2"
)

func Authenticate(c *fiber.Ctx) error {

	sessions := session.Sessions()
	store := sessions.Get(c)

	email := store.Get("email")
	authenticated := store.Get("authenticate")

	if authenticated != "true" || email == nil {
		return c.Redirect("/api/auth/login")
	}

	return c.Next()
}
