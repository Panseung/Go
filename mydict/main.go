package main

import (
	"fmt"

	"github.com/panseung/learngo/mydict"
)

func main() {
	dictionary := mydict.Dictionary{}
	baseword := "hello"
	dictionary.Add(baseword, "First")
	err := dictionary.Update(baseword, "Second")
	if err != nil {
		fmt.Println(err)
	}
	word, _ := dictionary.Search(baseword)
	fmt.Println(word)
}
 