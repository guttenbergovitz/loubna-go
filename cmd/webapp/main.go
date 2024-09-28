package main

import (
	"context"
	"fmt"
	"github.com/guttenbergovitz/loubna-go/web/templates"
	"net/http"
)

func main() {
	fs := http.FileServer(http.Dir("./web/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	component := templates.Hello("Loubna Photo")
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templates.Layout(component).Render(context.Background(), w)
		if err != nil {
			http.Error(w, "Failed to render page", http.StatusInternalServerError)
		}
	})

	fmt.Println("Listening on :3000")
	err := http.ListenAndServe(":3000", nil)
	if err != nil {
		fmt.Printf("Failed to start server: %v", err)
	}
}
