package model

import (
	"fmt"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // ID uint, CreatedAt time.Time, UpdatedAt time.Time, DeletedAt gorm.DeletedAt
	Name       string `json:"name" xml:"name" form:"name" query:"name"`
	Password   string `json:"password" xml:"password" form:"password" query:"password"`
	Age        int    `json:"age" xml:"age" form:"age" query:"age"`
}

// Userを引数に Userとerrorを返す
func UserAdd(u User) error {
	db := Connect()
	result := db.Create(&u)
	return result.Error
}

// intを引数に、Userとerrorを返す
func GetUser(id int) (User, error) {
	db := Connect()
	u := User{}
	result := db.Where("id = ?", id).First(&u)
	//fmt.Println(u, result.Error)
	return u, result.Error
}

// 引数無しに、[]Userとerrorを返す
func GetListUsers() ([]User, error) {
	db := Connect()
	var users []User
	result := db.Find(&users)
	fmt.Println(users, result.Error)
	return users, result.Error
}

func GetOrder(serch User) (User, error) {
	db := Connect()
	fmt.Println(serch)
	u := User{}
	result := db.Debug().Where(serch).First(&u)
	fmt.Println(u, result)
	return u, result.Error
}

func DeleteUser(id int) error {
	db := Connect()
	result := db.Debug().Delete(&User{}, id)
	fmt.Println(result)
	return result.Error
}

func UpdateUser(u User) (User, error) {
	db := Connect()

	result := db.Debug().First(&u)

	return u, result.Error
}
