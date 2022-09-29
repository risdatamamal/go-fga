package main

import (
	"fmt"
	"go-fga/user"
	"os"
	"strconv"
)

func main() {
	// 1. init data student pada init function
	// 2. program akan mengeluarkan pesan error, ketika id tidak ditemukan
	// "student dengan id xx tidak ada pada database"
	// 3. id harus positive integer
	var user1 user.User

	user1 = user.User{
		ID:   10,
		Name: "tara",
		DOB:  "1945-08-17",
		POB:  "Indonesia",
	}
	user1.CallName()

	user2 := user.User{
		ID:   10,
		Name: "tara2",
		DOB:  "1945-08-17",
		POB:  "Indonesia",
	}
	user2.CallName()

	argLen := len(os.Args[1:])
	fmt.Println("Argument length: ", argLen)

	// convert string menjadi int
	num, err := strconv.Atoi("1")
	fmt.Println(num, err)
}

func ChangeNum(source *int, destination int) {
	*source = destination
}

func ChangeName(src *user.User, name string) {
	// (*src).Name = name
	src.Name = name
}
