package main

import (
	"GoWebSite/model"
	"fmt"
)

func main() {
	user := model.User{Name: "aaa", Password: "aaa", Age: 20}
	model.UserAdd(user)
	test, err := model.GetUser(1)
	fmt.Println(test, err)
}
