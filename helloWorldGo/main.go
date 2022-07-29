//every file in a go program has to start with a package declaration
package main

//import package that going to be used
import (
	"fmt"
	"myApp/doctor"
)

//every file in a go program has to have a main function with no parameter in the parentheses
func main() {
	var whatToSay string = doctor.Intro()
	fmt.Println(whatToSay)
}
