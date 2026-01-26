package Nautilus

import (
	Apps "Nautilus/app/modules/apps"
	Errors "Nautilus/app/modules/errors"
	Tags "Nautilus/app/modules/tags"
	Thoughts "Nautilus/app/modules/thoughts"
	Users "Nautilus/app/modules/users"
	General "Nautilus/general"
	"fmt"
	"log"
	"os"
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
	app.Post("/add_thought", func(c *fiber.Ctx) error {
		err, creatorId := Users.GetUserIdByToken(c.Cookies("user"))
		if err != nil {
			return c.Status(200).JSON(map[string]interface{}{
				"status":  500,
				"message": "Erro com o seu usuário, entre em contato com o suporte",
			})
		}
		var thoughtStructed Thoughts.ThoughtSaved
		if err := c.BodyParser(&thoughtStructed); err != nil {
			fmt.Println(err)
			return fiber.NewError(fiber.StatusBadRequest, "Invalid JSON body")
		}
		// err = json.Unmarshal([]byte(thought), thoughtStructed)
		if err != nil {
			fmt.Errorf(err.Error())
			return c.Status(200).JSON(map[string]interface{}{
				"status":  500,
				"message": "Erro ao processar Objeto, entre em contato com o suporte",
			})
		}
		thoughtStructed.Creator_id = creatorId
		fmt.Println(General.JsonViewInterface(thoughtStructed), creatorId)
		Thoughts.AddTought(thoughtStructed)
		return c.Status(200).JSON(map[string]interface{}{
			"status":  200,
			"message": "Seu comentário foi adicionado !",
		})
	})
	app.Get("/", func(c *fiber.Ctx) error {
		userId := c.Cookies("user")
		if userId == "" {
			userId = os.Getenv("USER_ID")
		}
		fmt.Println(userId)
		err, users := Users.GetUsers(General.ToInt(userId))
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
		err, errors := Errors.GetErrorsByAppId(*app.Id)
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		bindInfos := General.CreateBindInfos("home")
		bindInfos["user_name"] = user.Name
		bindInfos["errors_selected"] = errors
		bindInfos["errors_selected_length"] = len(errors)
		bindInfos["user_profile_picture"] = user.Profile_picture
		bindInfos["app_name"] = app.Name
		return c.Render("home", bindInfos)
	})
	app.Get("/error_:error_id?", func(c *fiber.Ctx) error {
		error_id := c.Params("error_id")

		fmt.Println(error_id)
		if error_id == "" {
			return c.Redirect("/")
		}
		error_id_in_int := General.ToInt(error_id)
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
		err, savedError := Errors.GetErrors(error_id_in_int)
		if len(savedError) == 0 {
			return c.Status(500).JSON(map[string]interface{}{"Error": fmt.Sprintf("O erro %s não existe", error_id)})

		}
		errorId := c.Params("error_id")
		err, thoughts := Thoughts.GetThoughtByErrorId(General.ToInt(errorId))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		fmt.Println(General.JsonViewInterface(thoughts))
		bindInfos := General.CreateBindInfos("home")
		bindInfos["user_name"] = user.Name
		bindInfos["user_profile_picture"] = user.Profile_picture
		bindInfos["app_name"] = app.Name
		bindInfos["selected_error"] = savedError[0]
		bindInfos["thoughts"] = thoughts
		return c.Render("errors", bindInfos)
	})
	app.Get("/get_error_tags/:app_id", func(c *fiber.Ctx) error {
		app_id := c.Params("app_id")
		err, errors := Tags.GetErrorTagsByAppId(General.ToInt(app_id))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		return c.JSON(errors)
	})
	app.Get("/get_error_thoughts/:error_id", func(c *fiber.Ctx) error {
		errorId := c.Params("error_id")
		err, thoughts := Thoughts.GetThoughtByErrorId(General.ToInt(errorId))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		return c.JSON(thoughts)
	})
	app.Get("/get_errors/:app_id", func(c *fiber.Ctx) error {
		appId := c.Params("app_id")
		err, errors := Errors.GetErrorsByAppId(General.ToInt(appId))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		return c.Render(General.SubPartialsPath("errors", "errors_table"), errors)
	})
	app.Get("/get_files_from_error/:error_id", func(c *fiber.Ctx) error {
		errorId := c.Params("error_id")
		err, errors := Errors.GetErrors(General.ToInt(errorId))
		if err != nil {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": err.Error(),
			})
		}
		if len(errors) == 0 {
			return c.Status(500).JSON(map[string]interface{}{
				"Error": "Não existe um erro com esse id",
			})
		}
		return c.Status(200).JSON(errors[0].Files)
	})
	log.Fatal(app.Listen(":3120"))
}
