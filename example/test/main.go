package main

import "fmt"

func func1() {
	fmt.Println("func1 in")
	func2()
	fmt.Println("func1 out")
}

func func2() {
	fmt.Println("func2 in")
	func3()
	fmt.Println("func2 out")
}

func func3() {
	fmt.Println("handler")
}

func main() {
	func1()
}
