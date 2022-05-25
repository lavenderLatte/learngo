// the file name main.go is not optional; it's what gets compiled; it's entry point
package main
import (
	"fmt"
	"strings"
)


func mul(a int, b int) int {
	return a*b
}


// a func can return many things
func lenAndUpper(name string) (int, string) {
	defer fmt.Println("lenAndUpper function done!") // defer gets executed after this func returns
	return len(name), strings.ToUpper(name)
}
// "naked return" way
// func lenAndUpper(name string) (namelen int, upper string) { // here defind that I will return len and upper
// 	defer fmt.Println("naked lenAndUpper done!")
// 	namelen = len(name)
// 	upper = strings.ToUpper(name)
// 	return // naked return!
// }


// take any number of args
func manyargs(args ...string) {
	fmt.Println(args) // prints in a form of array!
}


func sumAll(nums ...int) int {
	// nums is an array
	// range returns idx and item of the array
	// for idx, num := range nums {
	// 	// fmt.Println(idx, num)
	// }

	sum := 0
	for _, num := range nums {
		// fmt.Println(idx, num)
		sum += num
	}

	return sum
}




// this main func is mandatory
func main() {
//// PRINT
	// fmt.Println("hello world")
	
//// DEFIND CONST AND VARS
	const name string = "hana" // constants, can't change
	// var fruit string = "apple" // varaible
	// fruit = "banana" // can redefind var
	// car := "subaru" // shorthand  way of defining a var

//// FUNCTIONS 
	// fmt.Println(mul(2, 2))

//// GET MORE THAN ONE RETURN VAL
	length, uppercase := lenAndUpper(name)
	// length, _ := lenAndUpper(name) // to ignore uppercase return val
	fmt.Println(length, uppercase)

//// FUNCTIONS THAT TAKES ANY NUMBER OF STRING ARGS
	manyargs("hana", "ravi", "junco", "ninja", "scroll")

//// Forloop
	fmt.Println(sumAll(1, 2, 3, 4, 5))

}