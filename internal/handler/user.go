package handler

import (
	"github.com/alidevjimmy/templaterenderer/internal/entity"
	"github.com/alidevjimmy/templaterenderer/internal/handler/message"
	"github.com/alidevjimmy/templaterenderer/internal/service"
	"log"
	"net/http"
	"strconv"
)

var (
	profileFiles = []string{
		"internal/template/base/base.html",
		"internal/template/page/profile.html",
		"internal/template/static/header.html",
		"internal/template/assets/bootstrap_bundle.html",
		"internal/template/assets/links.html",
	}

	homeFiles = []string{
		"internal/template/base/base.html",
		"internal/template/page/home.html",
		"internal/template/static/header.html",
		"internal/template/assets/bootstrap_bundle.html",
		"internal/template/assets/links.html",
	}
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	username := r.PathValue("username")
	if username == "" {
		render(w, nil, NotFoundFiles...)
		return
	}
	user, err := h.userService.ProfileByUsername(username)
	if err != nil {
		log.Println(err)
		render(w, nil, NotFoundFiles...)
		return
	}
	pm := message.ProfileMessage{
		User: user,
	}
	render(w, pm, profileFiles...)
}

func (h *UserHandler) Join(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	firstName := r.FormValue("firstName")
	lastName := r.FormValue("lastName")
	age := r.FormValue("age")
	log.Println(username, firstName, lastName, age)
	a, _ := strconv.Atoi(age)

	if username == "" ||
		firstName == "" ||
		lastName == "" ||
		a == 0 {

		hm := message.HomeMessage{
			Success: false,
			Errors:  []string{"please fill all the inputs"},
		}
		render(w, hm, homeFiles...)
		return
	}

	user := &entity.User{
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
		Age:       uint(a),
	}

	if err := h.userService.Join(user); err != nil {
		hm := message.HomeMessage{
			Success: false,
			Errors:  []string{err.Error()},
		}
		render(w, hm, homeFiles...)
		return
	}
	pm := message.ProfileMessage{
		User: user,
	}
	render(w, pm, profileFiles...)
}
