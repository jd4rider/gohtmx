package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	//"github.com/gofiber/helmet"
	"github.com/jd4rider/GoHtmx/templates"
	"github.com/joho/godotenv"
	"net/http"
)

type jsonbody struct {
	Body string `json:"body.data"`
}

func main() {
	app := fiber.New()

	//app.Use(helmet.New())

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	index := templates.Index()

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, index)
	})

	//bookpicker := bookpicker()

	app.Post("/bookselect", func(c *fiber.Ctx) error {
		payload := struct {
			Bibleselect string `json:"bibleselect"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		bookpicker := templates.Bookpicker(payload.Bibleselect)
		return Render(c, bookpicker)
		//return c.SendString("<div>You have selected the following Bible: " + payload.Bibleselect + "</div>")
	})

	app.Post("/biblecontent", func(c *fiber.Ctx) error {
		payload := struct {
			Bibid  string `json:"bibid"`
			Chapid string `json:"chapid"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		return c.SendString(templates.Biblecontent(payload.Bibid, payload.Chapid))
	})

	app.Static("/static", "./static")

	app.Use(NotFoundMiddleware)

	log.Fatal(app.Listen(":3000"))
}

func NotFoundMiddleware(c *fiber.Ctx) error {
	return Render(c, templates.NotFound(), templ.WithStatus(http.StatusNotFound))
}

func Render(c *fiber.Ctx, component templ.Component, options ...func(*templ.ComponentHandler)) error {
	componentHandler := templ.Handler(component)
	for _, o := range options {
		o(componentHandler)
	}
	return adaptor.HTTPHandler(componentHandler)(c)
}
