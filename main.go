package main

import "fmt"

func main() {
	var age int64
	var role, status string

	// Получаем входные данные и присваиваем их переменным
	age, role, status = getInput()
	fmt.Println(accessUser(age, role, status))
}

// Функция возвращает полученные значения, не принимает параметров
func getInput() (int64, string, string) {
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
	// Проверяем, что роль допустима
	if accessRole != "admin" && accessRole != "moderator" && accessRole != "user" {
		return false
	}

	// Admin и moderator имеют доступ всегда
	if accessRole == "admin" || accessRole == "moderator" {
		return true
	}

	// Для user проверяем возраст и статус
	if accessRole == "user" && accessStatus == "active" && accessAge >= 18 {
		return true
	}

	return false
}

/*
if status == "active" {true}
if age <18 {false}
if role == "admin" || "moderator" {true}
*/
