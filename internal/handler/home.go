package handler

import (
	"html/template"
	"net/http"
)

func Home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl, err := template.ParseFiles("internal/template/page/404.html")
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
	render(
		w,
		nil,
		"internal/template/page/home.html",
		"internal/template/static/header.html",
		"internal/template/static/footer.html",
		"internal/template/static/head.html",
		"internal/template/assets/bootstrap_bundle.html")
}
