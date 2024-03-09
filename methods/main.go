package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
	oneAge int
}

func (u User) GetStatus() {
	fmt.Println("Is user active:", u.Status)
}

func (u User) NewEmail() {
	u.Email = "test@example.com"
	fmt.Println("The email of the user is:", u.Email)
}

func main() {
	name := User{"Hitesh", "hitesh@gmail.com", true, 16, 45}
	fmt.Println("Welcome to methods.")
	fmt.Println(name.Name)
	name.GetStatus()
	name.NewEmail()
}
