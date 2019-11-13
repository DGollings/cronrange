package cronrange_test

import (
	"fmt"
	"time"

	"github.com/1set/cronrange"
)

// This example shows greeting according to your local box time.
func ExampleCronRange_IsWithin() {
	crGreetings := make(map[*cronrange.CronRange]string)
	crExprGreetings := map[string]string{
		"DR=360; 0 6 * * *":  "Good morning!",
		"DR=360; 0 12 * * *": "Good afternoon!",
		"DR=240; 0 18 * * *": "Good evening!",
		"DR=180; 0 22 * * *": "Good night!",
		"DR=300; 0 1 * * *":  "ZzzZZzzzZZZz...",
	}

	// create cronrange from expressions
	for crExpr, greeting := range crExprGreetings {
		if cr, err := cronrange.ParseString(crExpr); err == nil {
			crGreetings[cr] = greeting
		} else {
			fmt.Println("got parse err:", err)
			return
		}
	}

	// check if current time fails in any time range
	current := time.Now()
	for cr, greeting := range crGreetings {
		if isWithin, err := cr.IsWithin(current); err == nil {
			if isWithin {
				fmt.Println(greeting)
			}
		} else {
			fmt.Println("got check err:", err)
			break
		}
	}
}