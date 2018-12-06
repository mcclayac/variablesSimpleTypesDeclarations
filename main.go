package main

import "fmt"


const(
	message = "The answer to life is %d\n"     // declaration
	answer1 = iota * 2
	answer2
	message2 = "%d  %d \n"
	answer = 42
)

//var (
//	message = "The answer to life is %d\n"     // declaration
//	answer = 42
//)
func main () {

	//var message string
	//message := "The answer to life is %d\n"     // declaration
	//answer := 42
	//message = "Hello, world !\n"
	//answer ++
	fmt.Printf(message, answer )
	fmt.Printf(message2, answer1, answer2 )

	//var pi float64 = 3.14
	//var pi float32 = 3.14
	//pi := 3.14
	pi := float64(3.14)
	fmt.Printf("Value Float: %.2f\n", pi)


	//nine := 9
	//nine := uint(9)
	nine := uint8(9)
	//nine := uint32(9)
	//nine := uint64(9)

	fmt.Printf("Value Integer: %d\n", nine)

	var isTrue bool
	fmt.Printf("Value Boolean: %t\n", isTrue)

	abyte := byte(65)
	fmt.Printf("Value Byte: %x\n", abyte)


}
