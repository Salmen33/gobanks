package userhandler

import (
	conn "gobanks/config"
	session "gobanks/config"
	"gobanks/middleware/auth"
	us "gobanks/model/user"

	"github.com/gofiber/fiber/v2"
	//"github.com/google/uuid"
)

func New(c *fiber.Ctx) error {
	return c.Render("users/new", fiber.Map{}, "layouts/layout")
}

func Login(c *fiber.Ctx) error {
	return c.Render("users/login", fiber.Map{}, "layouts/layout")
}

func Home(c *fiber.Ctx) error {
	sessions := session.Sessions()

	store := sessions.Get(c)
	defer store.Save()

	email := store.Get("email")

	return c.Render("users/home", fiber.Map{
		"email": email,
	}, "layouts/layout")
}

func Auth(c *fiber.Ctx) error {
	var user us.User

	db := conn.DBConn()
	sessions := session.Sessions()

	email := c.FormValue("email")
	password := c.FormValue("password")

	if email == "" || password == "" {
		return c.Redirect("/login")
	}

	result := db.Where("email = ?", email).First(&user)

	if result == nil {
		return c.Redirect("/login")
	} else if password != user.Password && email != user.Email {
		return c.Redirect("/login")
	}

	compare := auth.CompareHash(user.Password, password)

	if compare != true {
		return c.Redirect("/api/auth/login")
	}

	store := sessions.Get(c)
	defer store.Save()

	store.ID()
	store.Set("id", user.ID)
	store.Set("email", user.Email)
	store.Set("authenticate", "true")

	return c.Redirect("/api/auth/login")
}

func Create(c *fiber.Ctx) error {
	var user us.User

	db := conn.DBConn()
	fullname := c.FormValue("fullname")
	password := c.FormValue("password")
	email := c.FormValue("email")
	hash := auth.HashPassword(password)
	result := db.Where("email = ?", email).Find(&user)

	if fullname == "" || password == "" || email == "" {
		return c.Redirect("/api/auth/new")
	}

	if result != nil {
		return c.Redirect("/api/auth/new")
	}

	us.NewUser(fullname, email, hash)

	return c.Redirect("/api/auth/login")
}

func AllUsers(c *fiber.Ctx) error {
	user := us.GetAllUser()

	return c.JSON(user)
}

func DeleteUser(c *fiber.Ctx) error {
	var user us.User

	nID := c.Params("id")
	db := conn.DBConn()

	db.Delete(&user, nID)

	return c.JSON("DELETED")
}

func Logout(c *fiber.Ctx) error {
	sessions := session.Sessions()

	store := sessions.Get(c)
	defer store.Save()

	store.Delete("email")
	store.Delete("id")
	store.Destroy()

	return c.Redirect("/api/auth/login")
}
