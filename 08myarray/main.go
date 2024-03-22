package main

import "fmt"

func main() {
	fmt.Println("Welcome to arrays in golang")
	var fruitList [4]string
	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "peach"
	fmt.Println("Fruit list is: ", fruitList)
}
