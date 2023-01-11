package main

import (
	"GoWebSite/model"
	"fmt"
)

func main() {
	fmt.Println("わけわからん")
	fmt.Println("今回はどうやろか")
	// AddUser
	user := model.User{Name: "Namae", Password: "pass", Age: 21}
	ok := model.UserAdd(user)
	if ok != nil {
		fmt.Println(ok)
	}
	// GetUser(id)
	id := 1
	getWithId, err := model.GetUser(id)
	if err != nil {
		fmt.Println("GetUserに失敗:", err)
	} else {
		fmt.Println(getWithId)
	}
	// GetListUsers
	users, err := model.GetListUsers()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(users)
	}
	//UpdateUser
	oldId := 4
	oldUser, err := model.GetUser(oldId)
	if err != nil {
		fmt.Println("GetUserに失敗:", err)
	}

	newUser := model.User{
		Name:     "id4User",
		Password: "4pass",
		Age:      444,
	}
	new, err := model.UpdateUser(oldUser, newUser)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("変更前:", oldUser)
	fmt.Println("変更後:", new)
}
