package main

import (
	"fmt"
	"todo-app-withgo/app/controllers"
	"todo-app-withgo/app/models"
)

func main() {
	fmt.Println(models.Db)

	controllers.StartMainServer()

}