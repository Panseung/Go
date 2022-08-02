package main

import "fmt"

func chV(str *string) {
	fmt.Println(*str)
}


func main() {
	myS := "hello"
	fmt.Println("1: ", &myS)
	chV(&myS)

}