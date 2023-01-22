package main

import "fmt"

func main() {
	var firstName *string = new(string)
	*firstName = "Matthew"
	fmt.Println(*firstName)

	lastName := "Matt"
	fmt.Println(lastName)

	ptr := &lastName

	fmt.Println(ptr, *ptr)

	lastName = "Jones"

	fmt.Println(ptr, *ptr)
}
