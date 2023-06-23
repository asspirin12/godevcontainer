package user

import "fmt"

type User struct {
	Name string
	Id   int
}

var user1 = User{
	Name: "Maryam",
	Id:   1,
}

func Hello() string {
	return fmt.Sprintf("Hello, %s", user1.Name)
}
