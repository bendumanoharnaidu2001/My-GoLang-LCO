package main

import "fmt"

const LoginToken string = "aksjdhs" //public

func main() {
	var username string = "Manohar"
	fmt.Println(username)
	fmt.Printf("variable if of type: %T \n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n", isLoggedIn)

	var smallInt int16 = 255
	fmt.Println(smallInt)
	fmt.Printf("variable is of type: %T \n", smallInt)

	var smallFloat float32 = 255.987654532
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n", smallFloat)

	//default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable)
	fmt.Printf("variable is of type: %T \n", anotherVariable)

	var anotherVariable1 float32
	fmt.Println(anotherVariable1)
	fmt.Printf("variable is of type: %T \n", anotherVariable1)

	//implicit type
	var website = "Bendu Manohar Nadiu"
	fmt.Println(website)
	fmt.Printf("variable is of type %T \n", website)

	//no var style
	numberOfUser := 300000.0
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n", LoginToken)
}
