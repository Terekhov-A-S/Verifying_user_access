package main

import "fmt"

func main() {
	var age int64
	var role, status string

	getInput(age, role, status)
	fmt.Println(accessUser(age, role, status))

}

func getInput(int64, string, string) (int64, string, string) {
	var getAge int64
	var getRole, getStatus string
	fmt.Print("Введите возраст: ")
	fmt.Scanln(&getAge)
	fmt.Print("Введите роль: ")
	fmt.Scanln(&getRole)
	fmt.Print("Введите статус: ")
	fmt.Scanln(&getStatus)
	return getAge, getRole, getStatus
}

func accessUser(accessAge int64, accessRole string, accessStatus string) bool {
	switch accessStatus {
	case "active":
		if accessRole == string("admin") {
			return true
		} else if accessRole == string("moderator") {
			return true
		} else if accessRole == "user" && accessAge >= 18 {
			return true
		}
	}
	return false
}

/*
if status == "active" {true}
if age <18 {false}
if role == "admin" || "moderator" {true}
*/
