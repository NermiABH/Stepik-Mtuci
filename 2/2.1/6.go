package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	if isAlmostHappy(n, 4) || isAlmostHappy(n, 7) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}

func isAlmostHappy(n int, happy int) bool {
	if happy == n || (!(happy > n/2) && (n%happy == 0 ||
		isAlmostHappy(n, happy*10+4) || isAlmostHappy(n, happy*10+7))) {
		return true
	}
	return false
}
