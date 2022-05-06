package main

import "fmt"

type User struct {
	Name       string
	FamilyName string
	Surname    string
}

type Worker struct {
	User
	Profession string
}

// ChangeName изменит оригинальную структуру
func (arg *User) ChangeName() {
	arg.Name = "new name"
}

type SubWorker struct {
	Worker
	Id string
}

func main() {
	user := User{Name: "Nikolai", Surname: "Sviridov", FamilyName: "Nikolaevich"}
	worker := Worker{user, "teacher"}
	fmt.Println(worker)
	worker.ChangeName()
	fmt.Println(worker)
	fmt.Println(user)
	subworker := SubWorker{Id: "abcd", Worker: worker}
	fmt.Println(subworker)
	subworker.ChangeName()
	fmt.Println(subworker)
}
