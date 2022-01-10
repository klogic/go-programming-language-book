package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	lesson()
	secs := time.Since(start)
	fmt.Println("with string.join", secs)

	exercise()
	secs2 := time.Since(start)
	fmt.Println("with normal for loop", secs2-secs)
}
func lesson() {
	fmt.Println(strings.Join(os.Args[1:], " "))
}
func exercise() {
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}
