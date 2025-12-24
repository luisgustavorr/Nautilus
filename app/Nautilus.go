package Nautilus

import (
	General "Nautilus/general"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
	"github.com/joho/godotenv"
)

var _ = godotenv.Load(".env")

var BindInfos fiber.Map = fiber.Map{
	"Title": "Nautilus",
}

func Start() {
	General.Setup()
	engine := html.New("./web/views", ".html")
	app := fiber.New(fiber.Config{
		ReadTimeout:  1 * time.Minute,
		WriteTimeout: 1 * time.Minute,
		BodyLimit:    20 * 1024 * 1024,
		Views:        engine,
	})
	app.Static("/", "./web/public/")

	app.Get("/", func(c *fiber.Ctx) error {

		return c.Render("home", General.CreateBindInfos("home"))
	})

	log.Fatal(app.Listen(":3120"))
}
