// echo2 prints its command-line arguments
package main

import (
	"fmt"
	"os"
)

func main(){
	lesson()
	exercise()
}

func lesson(){
	s, sep := "", ""
	for _, arg := range os.Args[1:]{
		s += sep + arg
		sep = " "
	}
	fmt.Println(s)
}

func exercise(){
	s, sep := "", ""
	for index, arg := range os.Args[1:]{
		fmt.Println(index, arg)
		s += sep + arg
		sep = " "
	}
}