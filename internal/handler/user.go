package handler

import (
	"github.com/alidevjimmy/templaterenderer/internal/entity"
	"net/http"
)

func Profile(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		http.Error(w, "invalid username", http.StatusBadRequest)
	}
	user := entity.User{
		Username: username,
	}
	render(
		w,
		user,
		"internal/template/page/profile.html",
		"internal/template/static/header.html",
		"internal/template/static/footer.html",
		"internal/template/static/head.html",
		"internal/template/assets/bootstrap_bundle.html")
}

func Login(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	if username == "" {
		http.RedirectHandler("/", http.StatusUnauthorized)
		return
	}
	user := entity.User{
		Username: username,
	}
	render(
		w,
		user,
		"internal/template/page/profile.html",
		"internal/template/static/header.html",
		"internal/template/static/footer.html",
		"internal/template/static/head.html",
		"internal/template/assets/bootstrap_bundle.html",
	)
}
