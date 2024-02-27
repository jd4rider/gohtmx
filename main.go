package main

import (
	"embed"
	"fmt"
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/joho/godotenv"
)

//go:embed static
var static embed.FS

type jsonbody struct {
	Body string `json:"body.data"`
}

type bibleId struct {
	Id   string
	Name string
}

func main() {

	//bibleIds := bibleid()
	//fmt.Println(bibleData["name"].(string))
	//fmt.Println(bibleIds)
	//for i := 0; i < len(bibleIds); i++ {
	//	fmt.Println(bibleIds[i].Name)
	//}
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	component := hello("jd4rider")

	index := index()

	http.Handle("/static/", http.FileServer(http.FS(static)))
	http.Handle("/hello", templ.Handler(component))
	http.Handle("/", templ.Handler(index))

	fmt.Println("Listening on :3000")
	fmt.Println("Press Ctrl+C to quit")
	http.ListenAndServe(":3000", nil)
}
