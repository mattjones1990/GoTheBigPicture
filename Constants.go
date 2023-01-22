package main

import "fmt"

const (
	numberOne   = iota //0
	numberTwo   = iota //1
	numberThree = iota //2

	first  = "one"
	second = "second"

	numberFour = iota //will by 5 as there's two other constants in the way ^
	numberFive        //assumes iota from above = 6

)

const (
	numberSix = iota //resets to zero
)

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

	fmt.Println(first, second)
	fmt.Println(numberOne, numberTwo, numberThree, numberFour, numberFive, numberSix)
}
