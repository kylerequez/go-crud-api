package main

import (
	"log"

	"github.com/kylerequez/go-crud-api/src/common"
	"github.com/kylerequez/go-crud-api/src/controllers"
)

func main() {
	log.Println(":::-::: Launching the application...")

	if err := common.LoadEnvVariables(); err != nil {
		log.Fatal(err)
	}

	controllers.InitApplication()
}
