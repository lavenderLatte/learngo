// the file name main.go is NOT optional; it's what gets compiled
// with different name, this file won't compile
// to run the file, >> go run main.go

/* ––––––––––––––––––––– SUMMARY –––––––––––––––––––––

- func main() {}  // entry point of file

- VARIABLES
	const varName type = "hana" // constants, can't change
	var varName type = "apple" // varaible
	varName = "banana"           // can redefind var
	car := "subaru"            // shorthand  way of defining a var

- PRINT
	fmt.Println("to print sth", "can print multiple items")

- reflect.TypeOf(car) // type()

- ARRAY & FOR-LOOP
	foodlist := []string{"ramen", "boba"} // array of strings
	nums := []int{1, 2, 3, 4} // array ints
	for idx, num := range nums {  // for-loop
		fmt.Println(idx, num)
	}

- FUNCTIONS
	• func fnName(args) (rtn type) {} // structure of functions
	• rtnVal1, rtnVal2 := returnTwo(name) // can return multiple items from one function
	• "naked return"
	• "variadic functions": functions take any number of args
	•	func sumAll(nums ...int) int {
			sum := 0
			for idx, num := range nums {
				fmt.Println(idx, num)
				sum += num
			}
			return sum
		}
		sumAll(1, 2, 3, 4, 5) --> will print 15; (1, 2, 3, 4, 5) becomes an array
	• "variable expression"
	•	func candrinkKOR(age int) bool {
		if koreanAge := age + 2; koreanAge > 18 { // creating var(koreanAge) right inside of if-stmt only for if-stmt
			return true
		}
		return false
	}
	• if-else
	• switch-case

- POINTERS & DEREFERENCING
	ref later part of this doc

- MAP (== DICT)
	mapName := map[keyType]valType{"name": "hana", "age": "23"}

- STRUCT (== object --> can make it have differnt type of value like python dict)
	foodlist := []string{"ramen", "boba"}  // create array to use it as value of favfood-key
	bio2 := person{name: "hana", age: 23, favfood: foodlist}

*/

package main

import (
	"fmt"
	"reflect" // for TypeOf()
	"strings"
)

type person struct {
	name    string
	age     int
	favfood []string
}

func mul(a int, b int) int {
	return a * b
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
	for idx, num := range nums {
		fmt.Println(idx, num)
	}

	sum := 0
	for _, num := range nums {
		// fmt.Println(idx, num)
		sum += num
	}
	return sum
}

func candrink(age int) bool {
	// if age > 18 {
	// 	return true
	// } // else
	// return false

	return age > 18 // shorter version
}

func candrinkKOR(age int) bool {
	// create var right inside of if-stmt; "variable expression"
	// I can create "koreanAge := age + 2" outside of if-stmt too
	// but doing this way indicates that I am using koreanAge var for if-stmt only
	if koreanAge := age + 2; koreanAge > 18 {
		return true
	} // else
	return false
}

func candrinkSwitch(age int) bool {
	switch age {
	case 10:
		return false
	case 18:
		return true
	}
	return false
}

func candrinkSwitch2(age int) bool {
	switch {
	case age > 50:
		return false
	case age > 18:
		return true
	}
	return false
}

// ENTRY POINT!!
// this main func is mandatory
// it's the entry point when I do >> go run main.go
func main() {
	//// PRINT
	fmt.Println("hello world")

	//// DEFIND CONST AND VARS
	const name string = "hana" // constants, can't change
	var fruit string = "apple" // varaible
	fruit = "banana"           // can redefind var
	car := "subaru"            // shorthand  way of defining a var
	qty := 1
	fmt.Println(name, fruit, reflect.TypeOf(car), reflect.TypeOf(qty))

	//// FUNCTIONS
	fmt.Println(mul(2, 2))

	//// GET MORE THAN ONE RETURN VAL
	length, uppercase := lenAndUpper(name)
	// length, _ := lenAndUpper(name) // to ignore uppercase return val
	fmt.Println(length, uppercase)

	//// FUNCTIONS THAT TAKES ANY NUMBER OF STRING ARGS "variadic functions"
	manyargs("hana", "ravi", "junco", "ninja", "scroll")

	//// Forloop
	fmt.Println(sumAll(1, 2, 3, 4, 5))

	//// IF, ELSE
	fmt.Println(candrink(17))
	fmt.Println(candrinkKOR(17))

	//// SWITCH
	fmt.Println("candrinkSwitch: ", candrinkSwitch(17))
	fmt.Println("candrinkSwitch2: ", candrinkSwitch2(37))

	//// Pointers & Dereferencing
	a := 2
	b := a // b points what a is pointing at
	a = 10
	fmt.Println(b) // prints 2; b is deep copy of a

	c := 2
	d := &c // d stores address of c
	c = 10
	fmt.Println(d)  // prints 0xc08809808; d is copy of ref to c
	fmt.Println(*d) // prints 10; *d to deference
	*d = 20
	fmt.Println(c) // prints 20; 0xc08809808 got updated by *d

	//// Map == dict
	// [key type]val type
	// value can't be other than what I defined, unlike python dict.
	bio := map[string]string{"name": "hana", "age": "23"}
	for key, val := range bio {
		fmt.Println(key, val)
	}

	//// struct == object --> can make it have differnt type of value like python dict
	// golang does not have class or objects only struct!
	// struct will also have methods
	// struct doesn't have constructor
	foodlist := []string{"ramen", "boba"}
	bio2 := person{name: "hana", age: 23, favfood: foodlist}
	fmt.Println(bio2)

}
