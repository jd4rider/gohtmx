package main

import (
	"github.com/a-h/templ"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/gofiber/fiber/v2/middleware/adaptor"
	//"github.com/gofiber/helmet"
	//"database/sql"
	"github.com/jd4rider/GoHtmx/templates"
	"github.com/joho/godotenv"
	//_ "github.com/tursodatabase/go-libsql"
	"net/http"
	"os"
)

type jsonbody struct {
	Body string `json:"body.data"`
}

var (
	language_code string
	language      string
)

func main() {

	//dbName := "file:./test.db"

	//db, err := sql.Open("libsql", dbName)
	//if err != nil {
	//	log.Fatal("failed to open db %s", err)
	//	os.Exit(1)
	//}
	//rows, err := db.Query("select language_code, language from translations")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer rows.Close()
	//for rows.Next() {
	//	err := rows.Scan(&language_code, &language)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	log.Info(language_code, language)
	//}
	//err = rows.Err()
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()

	app := fiber.New()

	//app.Use(helmet.New())

	if os.Getenv("API_KEY") == "" {

		err := godotenv.Load(".env")

		if err != nil {
			log.Fatalf("Error loading .env file: %s", err)
		}
	}

	index := templates.Index()

	app.Get("/", func(c *fiber.Ctx) error {
		return Render(c, index)
	})

	//bookpicker := bookpicker()

	app.Post("/bibleselect", func(c *fiber.Ctx) error {
		payload := struct {
			Langselect string `json:"langselect"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		biblepicker := templates.Biblepicker(payload.Langselect)
		return Render(c, biblepicker)
		//return c.SendString("<div>You have selected the following Bible: " + payload.Bibleselect + "</div>")
	})

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

	app.Post("/chapterselect", func(c *fiber.Ctx) error {
		payload := struct {
			Bibleselect string `json:"bibleselect"`
			Bookselect  string `json:"bookselect"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		chapterpicker := templates.Chapterpicker(payload.Bibleselect, payload.Bookselect)
		return Render(c, chapterpicker)
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

	app.Post("/biblecontentfinal", func(c *fiber.Ctx) error {
		payload := struct {
			Bibleselect   string `json:"bibleselect"`
			Bookselect    string `json:"bookselect"`
			Chapterselect string `json:"chapterselect"`
		}{}
		if err := c.BodyParser(&payload); err != nil {
			return err
		}

		//chapid := payload.Bookselect + "." + payload.Chapterselect

		return c.SendString(templates.Biblecontent(payload.Bibleselect, payload.Chapterselect))
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
