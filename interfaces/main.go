package main

import (
	"fmt"
	"log"
)

type User struct {
	Name  string
	Email string
}

func (u *User) Notify() error {
	_,err:= fmt.Printf("user:%v", u)
	return err
}

type Notifier interface{
	Notify() error
}

func SendNotification(n Notifier) error {
	return n.Notify()
}

type Admin struct {
	User
	Level string
}

func (a *Admin) Notify() error {
	log.Printf("Admin: Sending Admin Email To %s<%s>\n",
		a.Name,
		a.Email)

	return nil
}

func main() {
	admin := Admin{User{Name: "silviu", Email: "silviu.eigel@thinslices.com"}, "high"}

	SendNotification(&admin)
}