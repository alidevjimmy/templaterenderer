package handler

import (
	"github.com/alidevjimmy/templaterenderer/internal/handler/message"
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl, err := template.ParseFiles(NotFoundFiles...)
		if err != nil {
			http.NotFound(w, r)
			return
		}
		if err := tmpl.Execute(w, nil); err != nil {
			http.NotFound(w, r)
			return
		}
		return
	}
	hm := message.HomeMessage{
		Errors:  []string{},
		Success: true,
	}
	render(
		w,
		hm,
		homeFiles...)
}
