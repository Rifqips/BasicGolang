package main

import "fmt"

func main() {
	fmt.Println("Hello, World")
	fmt.Println("==============")
	var a, b int = 1, 2
	var c = a + b
	fmt.Println(c)
	d := 10
	fmt.Println(d)
	e := "hello"
	fmt.Println(e)
	//var variable_name data_type(optional) = value(optional)
	//variable_name := value (shorthand)
	fmt.Println("==============")
	const f = 100
	fmt.Println(f)
	const s string = "hello"
	fmt.Println(s)
	//	const variable_name data_type (optional) = value (optional)
	fmt.Println("==============")
	if 3%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	if 5 > 3 {
		fmt.Println("5 is greater")
	} else if 5 < 3 {
		fmt.Println("5 is equal to three")
	} else {
		fmt.Println("5 is less than 3")
	}
	//	An if statement can be used without a corresponding else statement whereas an else statement must always follow an if statement.
	//	else statement cannot be used without a corresponding if.
	fmt.Println("==============")
	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	fmt.Println("==============")
	i := 0
	for i <= 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println("==============")

	var sum int = 0

	for i := 1; i <= 100; i++ {

		if i%2 == 0 {
			sum = sum + i
		}
	}
	fmt.Println(sum)
}
