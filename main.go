package main

import (
	"fmt"
	"go-lesson-webapp_211103/app/models"
	"log"
)

func main() {

	// controllers.StartMainServer()

	//取得できるかテスト
	user, _ := models.GetUserByEmail("test@test.com")
	fmt.Println(user)

	session, err := user.CreateSession()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(session)

}
