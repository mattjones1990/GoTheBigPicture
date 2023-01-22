package main

import "fmt"

func main() {
	arrays()
	slices()
	slice2()
	mapFunc()
	structFunc()
}

// Arrays
func arrays() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	fmt.Println(arr)
	fmt.Println(arr[1])

	array := [3]int{1, 2, 3}
	fmt.Println(array)
	fmt.Println(array[1])
}

//Slices
func slices() {

	arrayForSlice := [3]int{1, 2, 3}
	slice := arrayForSlice[:]

	fmt.Println(arrayForSlice)
	fmt.Println(slice)
}

func slice2() {
	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 4) //resized itself. Don't need to worry about it.
	fmt.Println(slice)

	slice = append(slice, 6)
	fmt.Println(slice)

	// 0 1 2 3 4
	// 1 2 3 4 6
	s2 := slice[1:] //Will get 2,3,4,6
	fmt.Println(s2)

	s3 := slice[:2]
	fmt.Println(s3) //Will get 1,2

	s4 := slice[1:2]
	fmt.Println(s4) //Will get 2
}

//Map Data Type
func mapFunc() {
	m := map[string]int{"One": 1}
	fmt.Println(m)        //map[One:1]
	fmt.Println(m["One"]) //1

	delete(m, "One")
	fmt.Println(m) //map[] (empty)
}

//Structs
func structFunc() {
	//Can be any type of data we want
	//Fields are fixed at compile time
	type user struct {
		ID        int
		FirstName string
		LastName  string
	}

	var u user
	fmt.Println(u) //initalised to zero values. 0 "" "" (blank strings)

	u.ID = 1
	u.FirstName = "Matt"
	u.LastName = "Jones"

	fmt.Println(u)           //{1 Matt Jones}
	fmt.Println(u.FirstName) //Matt

	u2 := user{
		ID:        1,
		FirstName: "Matt",
		LastName:  "Jones",
	}

	fmt.Println(u2) //{1 Matt Jones}

	//struct can be moved to package level to use everywhere.
}
