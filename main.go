package main

import (
	"fmt"
	"go-lesson-webapp_211103/app/models"
)

func main() {
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)

	// log.Println("test")

	fmt.Println(models.Db)

	// u := &models.User{}
	// u.Name = "sage"
	// u.Email = "test@test.com"
	// u.Password = "test"
	// fmt.Println(u)

	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)

	// u.Name = "Test2"
	// u.Email = "hogehoge@test.com"
	// u.UpdateUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// u, _ := models.GetUser(2)
	// u.DeleteUser()
}
