package main

import (
	"fmt"
)

var heads = 0
var tails = 0

func main() {
	coinflip := "heads"
	switch coinflip {
	case "heads":
		heads++
	case "tails":
		tails++
	default:
		fmt.Println("landed on the edge")
	}
	fmt.Println("head is ", heads)
	fmt.Println("tails is ", tails)
}
