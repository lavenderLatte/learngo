// the file name main.go is NOT optional; it's what gets compiled
// with different name, this file won't compile
// to run the file; >> go run main.go
// to build the file; >> go build main.go --> this will create an excutable
// >> go mod init
// >> go mod tidy

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

package main // file name should match with package name

import (
	"fmt"
	"reflect" // for TypeOf()
	"strings"

	accounts "github.com/lavenderLatte/learngo/banking"
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

	// *k    : dereferencing; k is a pointer/addr, *k will get value/content of what k is pointing at
	// *k = 0: store the 0 in the memory location k refers to
	// &q    : returns address/pointer of var q (which is *int: pointer to q!!!!); q could be int/float/struc obj

	// In the example below &x(in func main()) and xPtr(in func zero) refers to same memory location.
	// This is what allows us to modify the original variable.
	// func zero(xPtr *int) { // xPtr is pointer to int == *int == &x
	// 	*xPtr = 0             // store the int 0 in the memory location xPtr refers to
	// }
	// func main() {
	// 	x := 5                // x is an int
	// 	zero(&x)              // &x returns a *int (pointer/addr to an int)
	// 	fmt.Println(x)        // x is 0
	// }

	//// MAP
	// [key type]val type
	// value can't be other than what I defined, unlike python dict.
	bio := map[string]string{"name": "hana", "age": "23"}
	for key, val := range bio {
		fmt.Println(key, val)
	}

	//// STRUCT
	// == object --> can make it have differnt type of value like python dict
	// golang does not have class or objects only struct!
	// struct will also have methods
	// struct doesn't have constructor
	foodlist := []string{"ramen", "boba"}
	bio2 := person{name: "hana", age: 23, favfood: foodlist}
	fmt.Println(bio2)

	// create bank account using public struct
	hanaAccount := accounts.Account{Owner: "hana", Balance: 1000.99} // this is bad way: anybody can create an account or modify balance
	hanaAccount.Owner = "ravi"                                       // security risk: owner of account can be changed since it's public
	fmt.Println(hanaAccount)

	// when a struct and its data is public, there is a security problem
	// but there is no constructor in Golang
	// we need to use a public function (NewAccount) to create a private struct object
	raviAccount := accounts.NewAccount("ravi")
	fmt.Println(raviAccount) // this prints &{ravi 0}

	raviAccount.Deposit(10.0)
	fmt.Println(raviAccount)
	fmt.Println(raviAccount.ViewBalance()) // only shows balance part

	raviAccount.Withdraw(5.0)
	fmt.Println("bal after withdrawing 5: ", raviAccount.ViewBalance())
	raviAccount.Withdraw(15.0)
	fmt.Println("bal after withdrawing 15: ", raviAccount.ViewBalance())
	// to print out error message and kill the program, do below
	// err := raviAccount.Withdraw(15.0)
	// if err != nil {
	// 	// log.Fatalln(err)
	// }
	// fmt.Println("bal after withdrawing another 15: ", raviAccount.ViewBalance()) // this won't get printed cuz program is killed

	raviAccount.ChangeOwner("hana")
	fmt.Println(raviAccount.ViewOwner())

	// with toString() method, we can print out multiple info at the same time in a meaningful way
	fmt.Println(raviAccount.String())
}
