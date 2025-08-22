package main

import "fmt"


var a int = 10
var b int = 20

// Correct type is int, not init
func printf(num int) {
	fmt.Println("The number is:", num)
}

func add(a int, b int) int {
	res := a + b
	printf(res)
	return res
}

func main() {
	add(a, b)
}
