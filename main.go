package main

import (
	"fmt"
)


type User struct {
	Name 			string
	Age int
	PhoneNumber string
	IsClose bool
	Rating 			float64
}

func NewUser(
	name string,
	age int,
	phoneNumber string,
	isClose bool,
	rating float64,
) User {
	if name == "" {
		return User{}
	}

	if age <= 0 || age >= 150 {
		return User{}
	}

	if phoneNumber == "" {
		return User{}
	}

	if rating < 0.0 || rating > 10.0 {
		return User{}
	}

	return User{
		Name: name,
		Age: age,
		PhoneNumber: phoneNumber,
		IsClose: isClose,
		Rating: rating,
	}
}

func (u *User) chengeName(newName string) {
	if newName != ""{
		u.Name = newName
	}
}

func (u *User) chenheAge(newAge int) {
	if newAge > 0 && newAge < 150 {
		u.Age = newAge
	}
}

func (u *User) chengeNumb(newNumb string) {
	if u.PhoneNumber != "" {
		u.PhoneNumber = newNumb
	}
}

func (u *User) CloseAcc() {
	u.IsClose = true
}

func (u *User) OpenAcc() {
	u.IsClose = false
}


func (u *User) RatingUp(raiting float64) {
	if u.Rating+raiting <= 10.0 {
		u.Rating += raiting
	}
}

func (u *User) RatingDown(raiting float64) {
	if u.Rating+raiting >= 0 {
		u.Rating += raiting
	}
}

func (u User) GetAge() int {
	return u.Age
}

func main()  {
	user := NewUser(
		"DAnu",
		55,
		"777 777 77 77",
		false,
		4.5,
	)

	user.chenheAge(5)

	fmt.Println(user.GetAge())

	fmt.Println(user)


	// ДОБАВИЛ КОД И ЗАПУШИЛ
	// Проверка 
}