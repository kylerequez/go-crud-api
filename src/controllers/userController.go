package controllers

import (
	"net/http"

	"github.com/kylerequez/go-crud-api/src/services"
)

type UserController struct {
	us *services.UserService
}

func NewUserController(us *services.UserService) *UserController {
	return &UserController{
		us: us,
	}
}

func (uc *UserController) InitUserController(mux *http.ServeMux) {
	mux.HandleFunc("GET /users", uc.GetAllUsers)
	mux.HandleFunc("GET /users/{id}", uc.GetUserById)
	mux.HandleFunc("POST /users", uc.InsertUser)
	mux.HandleFunc("PATCH /users/{id}", uc.PatchUpdateUser)
	mux.HandleFunc("PUT /users/{id}", uc.PutUpdateUser)
	mux.HandleFunc("DELETE /users/{id}", uc.DeleteUserById)
}

func (uc *UserController) GetAllUsers(res http.ResponseWriter, req *http.Request) {
	uc.us.GetAllUsers(res, req)
}

func (uc *UserController) GetUserById(res http.ResponseWriter, req *http.Request) {
	uc.us.GetUserById(res, req)
}

func (uc *UserController) InsertUser(res http.ResponseWriter, req *http.Request) {
	uc.us.InsertUser(res, req)
}

func (uc *UserController) PatchUpdateUser(res http.ResponseWriter, req *http.Request) {
	uc.us.PatchUpdateUser(res, req)
}

func (uc *UserController) PutUpdateUser(res http.ResponseWriter, req *http.Request) {
	uc.us.PutUpdateUser(res, req)
}

func (uc *UserController) DeleteUserById(res http.ResponseWriter, req *http.Request) {
	uc.us.DeleteUserById(res, req)
}
