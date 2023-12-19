package main

import (
	"fmt"
	"time"
)

func main() {

	//Here’s a basic switch.
	i := 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	}

	//You can use commas to separate multiple expressions in the same case statement. We use the optional default case in this example as well.
	switch time.Now().Weekday() {
	case time.Sunday, time.Saturday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	//switch without an expression is an alternate way to express if/else logic. Here we also show how the case expressions can be non-constants.
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm an int")
		default:
			fmt.Printf("Don't know type %T\n", t)
		}
	}
	whatAmI(true)
	whatAmI(1)
	whatAmI("hello")

}
