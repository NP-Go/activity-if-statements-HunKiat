package main

import "fmt"

func main() {
	number1 := 17
	number2 := 24
	resultMessage := "No outcome."
	//Insert your code here
	//Hint: You may wish to make use of strconv.Itoa to convert int to string
	fmt.Println("Performing 1st `if` comparison...")
	if number1, number2 := 7, 14; number1 < number2 {
		// number1 and number2 are scoped in this function
		fmt.Println("Number1 & Number 2 are ", number1, number2)
		fmt.Println("Number 1 is less than Number 2")
	}

	// number1 is 17 // global
	// number2 is 24 // global
	var ptrNum1, ptrNum2 *int
	ptrNum1 = &number1
	ptrNum2 = &number2
	fmt.Println("")
	fmt.Println("Performing 2nd `if` comparison...")
	fmt.Println("Number1 & Number 2 are ", *ptrNum1, *ptrNum2)
	ptrNum1, ptrNum2 = ptrNum2, ptrNum1
	if number1 != number2 {
		fmt.Println(resultMessage + " " + "The numbers compared are not the same.")
	} else {
		fmt.Println(resultMessage + " " + "The numbers compared are the same.")
	}
}
