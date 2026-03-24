package main

import "fmt"

// Ini akan jalan DULUAN sebelum main()
var _ = func() bool {
	fmt.Println("Halo dari Variable Initialization!")
	return true
}()

func main() {
	fmt.Println("Halo dari Main!")
}
