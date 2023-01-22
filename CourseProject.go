package main

import (
	"fmt"

	"example.com/basic/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Matt",
		LastName:  "Jones",
	}

	fmt.Println(u)
}
