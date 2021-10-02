package main

import "fmt"

//in go if we declare a variable and not use it we get error
//to get the type use format specifiers: %T - for string
func main() {
	fmt.Println("hello world")
	s := make([]string, 3)
	fmt.Println("emp:", s)

}
