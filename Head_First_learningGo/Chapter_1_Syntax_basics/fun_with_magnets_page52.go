// I started with 10 apples.
// Some jerk ate 4 apples.
// There are 6 apples left.

package main

import "fmt"

func main() {
	var originalCount int
	var eatenCount int
	originalCount = 10
	eatenCount = 4
	fmt.Println("I started with", originalCount, "apples.")
	fmt.Println("Some jerk ate", eatenCount, "apples.")
	fmt.Println("There are", originalCount-eatenCount, "apples left.")
}
