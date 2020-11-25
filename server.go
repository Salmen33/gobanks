package main

import (
	conn "gobanks/config"
	"gobanks/model"

	"gobanks/config"
	"gobanks/model/user"
	"gobanks/route"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/handlebars"
)

func main() {
	config.DBConn()

	//key := conn.GenerateKey()
	//encrypt := conn.Encrypt(key, "This is secret")
	//decrypt := conn.Decrypt(key, encrypt)

	db := conn.DBConn()
	db.AutoMigrate(&user.User{})
	db.AutoMigrate(&user.Transfer{})
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Price{})

	engine := handlebars.New("./views", ".hbs")
	r := fiber.New(fiber.Config{
		Views: engine,
	})

	r.Static("/static/", "./static")
	route.RouteInit(r)
	r.Listen(":8080")
}
