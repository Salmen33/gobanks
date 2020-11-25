package route

import (
	trans "gobanks/handler/transactionhandler"
	user "gobanks/handler/userhandler"
	middle "gobanks/middleware/auth"

	"github.com/gofiber/fiber/v2"

	//"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func RouteInit(app *fiber.App) {

	app.Use(CommonMiddleware)
	//app.Use(middle.IfAuth)
	app.Use(logger.New())
	//app.Use(csrf.New())

	app.Get("/all", user.AllUsers)
	app.Post("/auth", user.Auth)
	app.Post("/insert", user.Create)
	app.Post("/logout", user.Logout)
	app.Get("/", user.Home)

	secure := app.Group("/api")

	v1 := secure.Group("/auth", middle.IfAuth)
	v1.Get("/new", user.New)
	v1.Get("/login", user.Login)

	v2 := secure.Group("/user", middle.Authenticate)
	v2.Get("/deposite", trans.NewDeposite)
	v2.Post("/depo", trans.Deposite)
	v2.Get("/service", trans.NewService)
	v2.Get("/new", trans.NewTransfer)
	v2.Post("/transfer", trans.Transfer)
	v2.Get("/transfer/history", trans.History)
	v2.Post("/delete/:id", user.DeleteUser)
	//v1.Get("/:id", user.Profile)

}

func CommonMiddleware(c *fiber.Ctx) error {
	// Set some security headers:
	c.Set("X-XSS-Protection", "1; mode=block")
	c.Set("X-Content-Type-Options", "nosniff")
	c.Set("X-Download-Options", "noopen")
	c.Set("Strict-Transport-Security", "max-age=5184000")
	c.Set("X-Frame-Options", "SAMEORIGIN")
	c.Set("X-DNS-Prefetch-Control", "off")

	// Go to next middleware:
	return c.Next()
}
