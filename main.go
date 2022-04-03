package main

import (
	"drink_hack_project/app/controllers"
	"drink_hack_project/app/models"
	"drink_hack_project/config"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(config.Config.Port)
	fmt.Println(config.Config.SQLDriver)
	fmt.Println(config.Config.DbName)
	fmt.Println(config.Config.LogFile)

	log.Println("test")

	fmt.Println(models.Db)

	//u := &models.User{}
	//u.Name = "test"
	//u.Email = "test@example.com"
	//u.Password = "test_test"
	//fmt.Println(u)
	//
	//u.CreateUser()
	//u, _ := models.GetUser(1)
	//fmt.Println(u)
	//
	//u.Name = "Test2"
	//u.Email = "test2@example.com"
	//u.UpdateUser()
	//u, _ = models.GetUser(1)
	//fmt.Println(u)
	//
	//u.DeleteUser()
	//u, _ = models.GetUser(1)
	//fmt.Println(u)

	fmt.Println(models.Db)

	port := os.Getenv("PORT")
	fmt.Println(port)

	controllers.StartMainServer()

	//user, _ := models.GetUserByEmail("hoge@hoge")
	//fmt.Println(user)
	//
	//session, err := user.CreateSession()
	//if err != nil {
	//	log.Println(err)
	//}
	//
	//fmt.Println(session)
	//valid, _ := session.CheckSession()
	//fmt.Println(valid)
}