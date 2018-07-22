// Example usages of defer statement
package main

import "fmt"

func main() {
	defer fmt.Println("Ignition!")

	for x := 1; x <= 10; x++ {
		defer fmt.Printf("%d...\n", x)
	}

	fmt.Println("Countdown:")
}
