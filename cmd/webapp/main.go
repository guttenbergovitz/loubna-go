package main

import (
	"fmt"
	"github.com/guttenbergovitz/loubna-go/web/templates"
	"net/http"

	"github.com/a-h/templ"
)

func main() {
	component := templates.Hello("Loubna Photo")

	http.Handle("/", templ.Handler(component))

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
