package main

import "fmt"

func main() {
	hanoi(1, 3, 2, 3)
}

// 1 -> 2 (2), 1 -> 3 (1), 2 -> 3 (2)
func hanoi(src int, dest int, temp int, count int) {
	if count == 0 {
		return
	}
	hanoi(src, temp, dest, count-1)
	fmt.Printf("%d -> %d\n", src, dest)
	hanoi(temp, dest, src, count-1)
}
