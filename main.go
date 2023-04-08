// at least, when you run golang file, there's one <package main> exists
package main

import "fmt"
import "strings"
import "math"

func main() {
	fmt.Println("Hello, world!")
	fmt.Printf("Hello, %s!\n", "world")

	//* To declare variable, <var> <variable-name> <type> = <value>
	var firstName string = "John"
	var lastName string = "Doe"
	fmt.Printf("Hello, %s %s!\n", firstName, lastName)

	//* To declare variable without data type, use <variable-name> <:=> instead
	// without var
	age := 20
	// with var
	var age2 int = 20
	fmt.Printf("Hello, %s %s! You are %d %d years old.\n", firstName, lastName, age2, age)

	//* To declare multiple variables
	// without initial value
	var first, second, third string
	first, second, third = "first", "second", "third"
	fmt.Println(first, second, third)

	// with initial value
	var fourth, fifth, sixth string = "fourth", "fifth", "sixth"
	fmt.Println(fourth, fifth, sixth)

	// if you want more efficient way and without declare data type, use <:=> instead
	seventh, eighth, ninth := "seventh", "eighth", "ninth"
	fmt.Println(seventh, eighth, ninth)

	//* underscore is a special variable name, it means "I don't care about this value"
	_, _ = "halo", "halo"
	// but if you want to use it, you can use it like this
	name, _ := "halo", "halo"
	fmt.Printf("Hello, %s!\n", name)

	//* Also, you can declare variable with "new" keyword
	// but it's not recommended to use it's because it's not efficient
	newVar := new(string)
	*newVar = "Hello, world!"
	// If you print newVar, it will print the memory address
	fmt.Println(newVar)
	// newVar is a pointer, so you need to use * to get the valued
	fmt.Println(*newVar)

	//* To declare constant, use <const> <variable-name> <type> = <value>
	// you can't change the value of constant
	const pi float64 = 3.14
	fmt.Println(pi)
	// you can declare multiple constants like this
	const (
		pi2 float64 = 3.14
		e  float64 = 2.71
	)
	fmt.Println(pi2, e)
	// or like this
	const pi3, e2 float64 = 3.14, 2.71
	fmt.Println(pi3, e2)

	//* Temporary Variables in conditional statements
	// you can use temporary variables in conditional statements
	if tmp := 10; tmp > 5 {
		fmt.Println("tmp is greater than 5")
	}
	if(pi < 5) {
		fmt.Println("tmp is less than 5")
	}
	// but you can't use it outside the conditional statements
	//! fmt.Println(tmp)

	// Fallthrough in switch statements
	// you can use fallthrough in switch statements
	// it will execute the next case
	switch number := 10; number {
	case 10:
		fmt.Println("10")
		fallthrough
	// it will execute the next case, even if the condition is not true
	case 20:
		fmt.Println("20")
	case 30:
		fmt.Println("30")
	}

	//* To declare array, use <var> <variable-name> [<length>] <type>
	// you can't change the length of array
	var arr [5]int
	arr[0] = 1
	arr[1] = 2
	arr[2] = 3
	arr[3] = 4
	arr[4] = 5
	fmt.Println(arr)
	// you can declare array with initial value
	var arr2 [6]int = [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(arr2)
	// or like this
	arr3 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(arr3)
	// you can declare array with ellipsis / without length
	arr4 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println(arr4)
	// Also, you can declare array with make function
	arrMake := make([]int, 5)
	fmt.Println(arrMake)

	//* you can declare array with multi-dimensional
	fmt.Println("\nMULTI-DIMENSIONAL ARRAY")
	var arr5 [2][3]int
	arr5[0][0] = 1
	arr5[0][1] = 2
	arr5[0][2] = 3
	arr5[1][0] = 4
	arr5[1][1] = 5
	arr5[1][2] = 6
	fmt.Println(arr5)
	// you can declare array with multi-dimensional with initial value
	var arr6 [2][3]int = [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr6)
	// or like this
	arr7 := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr7)
	// you can declare array with multi-dimensional with ellipsis / without length
	arr8 := [...][3]int{{1, 2, 3}, {4, 5, 6}}
	fmt.Println(arr8)

	//* you can use range to get the index and value of array
	fmt.Println("\nRANGE")
	for i, value := range arr8 {
		fmt.Println("index:", i, "value:", value)
	}
	// you can use _ to ignore the index
	for _, value := range arr8 {
		fmt.Println("value:", value)
	}

	//* To declare slice, use <var> <variable-name> []<type>
	// you can change the length of slice
	fmt.Println("\nSLICE")
	var slice []int
	slice = append(slice, 1)
	slice = append(slice, 2)
	slice = append(slice, 3)
	slice = append(slice, 4)
	slice = append(slice, 5)
	fmt.Println(slice)
	// you can declare slice with initial value
	var slice2 []int = []int{1, 2, 3, 4, 5}
	fmt.Println(slice2)
	// or like this
	slice3 := []int{1, 2, 3, 4, 5}
	fmt.Println(slice3)
	// you can declare slice with ellipsis / without length
	slice4 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(slice4)

	//* Connection between slice and array
	// slice is a reference to array
	// so if you change the value of slice, it will change the value of array
	fmt.Println("\nCONNECTION BETWEEN SLICE AND ARRAY")
	arr9 := [5]int{1, 2, 3, 4, 5}
	slice5 := arr9[0:3]
	fmt.Println(slice5)

	//* To declare map, use <var> <variable-name> map[<key-type>]<value-type>
	// you can't change the length of map
	fmt.Println("\nMAP")
	var map1 map[string]int
	map1 = make(map[string]int)
	map1["one"] = 1
	map1["two"] = 2
	map1["three"] = 3
	fmt.Println(map1)
	// you can declare map with initial value
	var map2 map[string]int = map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(map2)
	// or like this
	map3 := map[string]int{"one": 1, "two": 2, "three": 3}
	fmt.Println(map3)
	fmt.Println(map3["one"])
	// you can use len function to get the length of map
	fmt.Println(len(map3))
	// you can use delete function to delete the value of map
	delete(map3, "one")
	fmt.Println(map3)
	// you can use range to get the key and value of map
	for key, value := range map3 {
		fmt.Println("key:", key, "value:", value)
	}
	// you can use _ to ignore the key
	for _, value := range map3 {
		fmt.Println("value:", value)
	}
	// You can use comma ok idiom to check if the key is exist
	value, isExist := map3["two"]
	if isExist {
		fmt.Println("Value exists, that is:", value)
	} else {
		fmt.Println("key is not exist")
	}
	//* You can combine slice and map
	fmt.Println("\nCOMBINE SLICE AND MAP")
	var chickens = []map[string]string{
		{"name": "chicken blue", "gender": "male"},
		{"name": "chicken red", "gender": "female"},
	}
	fmt.Println(chickens)
	fmt.Println(chickens[0]["name"])
	for _, chicken := range chickens {
		fmt.Println(chicken["name"], chicken["gender"])
	}
	// Also, you can define different type of data in map
	var data = []map[string]string{
    {"name": "chicken blue", "gender": "male", "color": "brown"},
    {"address": "mangga street", "id": "k001"},
    {"community": "chicken lovers"},
	}
	fmt.Println(data)
	fmt.Println(data[0]["name"])
	fmt.Println(data[1]["address"])

	//* To declare function, use <func> <function-name> (<parameter-name> <parameter-type>) <return-type>
	// Function without return value
	fmt.Println("\nFUNCTION")
	sayHello([]string{"John", "Wick"}, []int{20, 21})
	// Function with return value
	fmt.Println(sum(10, 20))
	// Function with multiple return value
	fmt.Println("\nMULTIPLE RETURN VALUE")
	additionSubstraction := func(value1 int, value2 int) (int, int) {
		return value1 + value2, value1 - value2
	}
	result1, result2 := additionSubstraction(10, 5)
	fmt.Println(result1, result2)
	// Function with predefined arguments
	fmt.Println("\nPREDEFINED ARGUMENTS")
	area, circumference := calculate(10)
	fmt.Println(area, circumference)
	// Variadic function
	fmt.Println("\nVARIADIC FUNCTION")
	var avg = calculateVariadic(1, 2, 3, 4, 5)
	fmt.Println(avg)
	// Function as parameter
	fmt.Println("\nFUNCTION AS PARAMETER")
	var callbackData = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(callbackData, func(each string) bool {
		return strings.Contains(each, "o")
	})
	fmt.Println("data asli", callbackData)
	fmt.Println("data contains 'o'", dataContainsO)


	//* To declare struct, use <type> <struct-name> struct {<field-name> <field-type>}
	//* you can use struct to create your own data type
	fmt.Println("\nSTRUCT")
	type Person struct {
		name string
		age  int
	}
	var person1 Person
	person1.name = "John"
	person1.age = 20
	fmt.Println(person1)
}


// you can declare function outside the main function
func sayHello(name []string, age []int) {
	fmt.Println("Hello", strings.Join(name, ", "), "your age is", age)
}

// Function with return value
func sum(value1 int, value2 int) int {
	return value1 + value2
}

// Function with multiple return value and predefined return value
func calculate(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d / 2, 2)
	circumference = math.Pi * d

	return
}

// variadic function
func calculateVariadic(numbers ...int) float64 {
	total := 0
	for _, number := range numbers {
		total += number
	}
	var avg = float64(total) / float64(len(numbers))
	return avg
}

// Callback function
func filter(data []string, callback func(string) bool) []string {
	var result []string
	for _, each := range data {
			if filtered := callback(each); filtered {
					result = append(result, each)
			}
	}
	return result
}