package main

// fmt = format package, used for formatted I/O operations, such as printing to the console.
import "fmt"

func main() {
	//Simple Print without newline
	fmt.Print("Hello World")
	//Println adds a newline at the end
	fmt.Println("I am learning Go")
	//Printf allows for formatted strings, similar to C's printf
	fmt.Printf("The value of Pi is approximately %.2f\n", 3.14159)
}
