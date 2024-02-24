package main

import (
	"embed"
	"fmt"
	"net/http"

	"github.com/a-h/templ"
)

//go:embed static
var static embed.FS

func main() {
	component := hello("jd4rider")

	index := index()

	http.Handle("/static/", http.FileServer(http.FS(static)))
	http.Handle("/hello", templ.Handler(component))
	http.Handle("/", templ.Handler(index))

	fmt.Println("Listening on :3000")
	fmt.Println("Press Ctrl+C to quit")
	http.ListenAndServe(":3000", nil)
}
