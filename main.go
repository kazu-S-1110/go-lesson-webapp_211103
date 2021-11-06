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

	// fmt.Println(models.Db)

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
	// u, _ = models.GetUser(1)o
	// fmt.Println(u)

	// u, _ := models.GetUser(1)
	// u.DeleteUser()
	// user, _ := models.GetUser(1)
	// fmt.Println(user)
	// user.CreateTodo("first todo")

	t, _ := models.GetTodo(5)
	fmt.Println(t)

	// todos, _ := models.GetTodos()
	// fmt.Println(todos)
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todos, _ := u.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }
	// t.Content = "first todo updated!"
	// t.UpdateTodo()
	t.DeleteTodo()

}
