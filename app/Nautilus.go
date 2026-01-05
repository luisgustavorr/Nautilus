package Nautilus

import (
	Apps "Nautilus/app/crud/apps"
	Errors "Nautilus/app/crud/errors"
	Users "Nautilus/app/crud/users"
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
		err, users := Users.GetUsers(General.ToInt(c.Cookies("user")))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{"Error": err.Error()})
		}
		if len(users) != 1 {
			return c.Status(500).JSON(map[string]interface{}{"Error": "tem algo de errado com o seu login, contate o suporte"})
		}
		user := users[0]
		err, apps := Apps.GetApp(user.Id_apps)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{"Error": err.Error()})
		}
		if len(apps) != 1 {
			return c.Status(500).JSON(map[string]interface{}{"Error": "tem algo de errado com o seu login, contate o suporte, apperr"})
		}
		app := apps[0]
		bindInfos := General.CreateBindInfos("home")
		bindInfos["user_name"] = user.Name
		bindInfos["user_profile_picture"] = user.Profile_picture
		bindInfos["app_name"] = app.Name
		return c.Render("home", bindInfos)
	})
	app.Get("/get_errors/:app_id", func(c *fiber.Ctx) error {
		app_id := c.Params("app_id")
		err, errors := Errors.GetErrorsByAppId(General.ToInt(app_id))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		return c.JSON(errors)
	})
	log.Fatal(app.Listen(":3120"))
}
