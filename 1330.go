// 1330: 두 수 비교하기
// Level: Bronze5
// Link:https://www.acmicpc.net/problem/1330

package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Scanln(&a, &b)
	if sum_v := a - b; sum_v > 0 {
		fmt.Println(">")
		return
	} else if sum_v == 0 {
		fmt.Println("==")
		return
	} 
	fmt.Println("<")
}

