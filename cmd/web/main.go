package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Zhabskyi/go-web-app/pkg/config"
	"github.com/Zhabskyi/go-web-app/pkg/handlers"
	"github.com/Zhabskyi/go-web-app/pkg/render"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig

	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Can not create template cache")
	}

	app.TemplateCache = tc

	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)

	fmt.Printf("Starting application on port %s \n", portNumber)
	http.ListenAndServe(portNumber, nil)
}
