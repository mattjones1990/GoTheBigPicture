package main

import "fmt"

func main() {
	const pi = 3.1415 //must be initialised at the same time
	fmt.Println(pi)

	const c = 3
	fmt.Println(c + 3)

	fmt.Println(c + 1.2) //interpreted each time.
	//If we declared it as an int at the start, the above would error.

	//Example with int declared
	const d int = 3
	fmt.Println(d + 3)

	fmt.Println(float32(d) + 1.2)
}
