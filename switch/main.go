// Example usages of switch case statement
package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Println("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s", os)
	}

	// switch case with no condition
	fmt.Println("When is it?")
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("Morning!")
	case t.Hour() < 17:
		fmt.Println("Afternoon!")
	default:
		fmt.Println("Evening!")
	}
}
