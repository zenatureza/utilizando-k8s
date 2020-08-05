package main

import (
	"./lib"
	"html/template"
	"log"
	"net/http"
)

type Page struct {
	Greeting template.HTML
}

// thanks to: https://stackoverflow.com/a/42055211/7250195
func noescape(str string) template.HTML {
	return template.HTML(str)
}

func main() {
	templates := template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		p := Page{Greeting: noescape(lib.Greeting("Code.education Rocks!"))}

		if err := templates.ExecuteTemplate(w, "index.html", p); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
