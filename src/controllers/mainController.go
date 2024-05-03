package controllers

import (
	"log"
	"net/http"
	"os"

	"github.com/kylerequez/go-crud-api/src/common"
	"github.com/kylerequez/go-crud-api/src/repositories"
	"github.com/kylerequez/go-crud-api/src/services"
)

func InitApplication() {
	mux := http.NewServeMux()

	var servHost, servPort = os.Getenv("SERVER_HOST"), os.Getenv("SERVER_PORT")

	if err := InitControllers(mux); err != nil {
		log.Fatal(err)
		return
	}

	if err := http.ListenAndServe(servHost+":"+servPort, mux); err != nil {
		log.Fatal(":::-::: ", err.Error())
		return
	}
}

func InitControllers(mux *http.ServeMux) (err error) {
	log.Println(":::-::: Loading Controller/s...")

	var dbName string = os.Getenv("DB_NAME")

	log.Println(":::-::: Loading User Controller...")
	ur := repositories.NewUserRepository(dbName, common.COLLECTIONS["USER_COLLECTION"])
	us := services.NewUserService(ur)
	uc := NewUserController(us)
	uc.InitUserController(mux)
	log.Println(":::-::: Successfully Loaded User Controller")

	log.Println(":::-::: Successfully Loaded All Controller/s")

	return nil
}
