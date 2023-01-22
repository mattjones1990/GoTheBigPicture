package main

import "fmt"

func main() {
	//declare simple variables
	var i int
	i = 42
	fmt.Println(i)

	//another syntax
	var f float32 = 3.14
	fmt.Println(f)

	//simpler
	firstName := "Matt"
	fmt.Println(firstName)

	b := true
	fmt.Println(b)

	c := complex(3, 4)
	fmt.Print(c)

	r, im := real(c), imag(c)

	fmt.Println(r)
	fmt.Println(im)

}
