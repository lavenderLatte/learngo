// the file name main.go is not optional; it's what gets compiled; it's entry point
package main

import (
	"fmt"
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

func candrink(age int) bool {
	if age > 18 {
		return true
	} // else
	return false
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

	//// IF, ELSE
	fmt.Println(candrink(17))
	fmt.Println(candrinkKOR(17))

	//// SWITCH
	fmt.Println("candrinkSwitch: ", candrinkSwitch(17))
	fmt.Println("candrinkSwitch2: ", candrinkSwitch2(37))

	//// Pointers & Dereferencing
	a := 2
	b := a
	a = 10
	fmt.Println(b) // prints 2; b is deep copy of a

	c := 2
	d := &c
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

	//// struct == python dict
	foodlist := []string{"ramen", "boba"}
	bio2 := person{name: "hana", age: 23, favfood: foodlist}
	fmt.Println(bio2)

}
