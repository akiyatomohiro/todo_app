package main

import (
	"fmt"
	"todo_app/app/controllers"
	"todo_app/app/models"
)

func main() {
	fmt.Println(models.Db)
	// fmt.Println(config.Config.Port)
	// fmt.Println(config.Config.SQLDriver)
	// fmt.Println(config.Config.DbName)
	// fmt.Println(config.Config.LogFile)
	// log.Println("test")

	// fmt.Println(models.Db)
	// u := &models.User{}
	// u.Name = "test3"
	// u.Email = "test3@example.com"
	// u.PassWord = "testtest"
	// fmt.Println(u)
	// u.CreateUser()

	// u, _ := models.GetUser(1)
	// fmt.Println(u)
	// u.Name = "test2"
	// u.Email = "test2@example.com"
	// u.UpdateUser()

	// u, _ = models.GetUser(1)
	// fmt.Println(u)
	// u.DeleteUser()
	// u, _ = models.GetUser(1)
	// fmt.Println(u)

	// user, _ := models.GetUser(2)
	// user.CreateTodo("Third todo")

	// t, _ := models.GetTodo(1)
	// fmt.Println(t)

	// todos, _ := models.GetTodos()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// todos, _ := user.GetTodosByUser()
	// for _, v := range todos {
	// 	fmt.Println(v)
	// }

	// t, _ := models.GetTodo(1)
	// t.Content = "Update todo"
	// t.UpdateTodo()

	// t, _ := models.GetTodo(2)
	// t.DeleteTodo()

	controllers.StartMainServer()

	// u, _ := models.GetUserByEmail("test@example.com")
	// fmt.Println(u)

	// s, e := u.CreateSession()
	// if e != nil {
	// 	log.Println(e)
	// }
	// fmt.Println(s)

	// valid, _ := s.CheckSession()
	// fmt.Println(valid)
}
