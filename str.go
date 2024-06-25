package main

import (
	"fmt"
)

var m map[int]string = map[int]string{
	0: "0",
	1: "1",
	2: "2",
	3: "3",
	4: "4",
	5: "5",
	6: "6",
	7: "7",
	8: "8",
	9: "9",
}

func main() {
	fmt.Println(str(0))
	fmt.Println(str(666))
	fmt.Println(strtoint("420") + 420)
	fmt.Println('1' + '1')
	fmt.Println(is_identifier("a123"))
	fmt.Println(is_identifier("12a"))
	fmt.Println(is_identifier(""))
}

func str(n int) string {

	if n/10 == 0 {
		return m[n]
	}

	return str(n/10) + m[n%10]
}

func pow(n int, m int) int {
	if m == 0 {
		return 1
	}
	return n * pow(n, m-1)
}

func strtoint(s string) int {
	n := 0

	for i := 0; i < len(s); i++ {
		n += int(s[len(s)-i-1]-'0') * pow(10, i)
	}

	return n
}

//    65 90 97 122
// 354 = 3*10^2+5*10^1+4*10^0

//a123 -- identifier
//12a -- NOT identifier

func is_identifier(s string) bool {
	if s == "" {
		return false
	}
	prefix := int(s[0])
	rest := s[1:]
	if prefix >= 65 && prefix <= 90 || prefix >= 97 && prefix <= 122 {
		for i := 0; i < len(rest); i++ {
			if !(int(rest[0]) >= 48 && int(rest[0]) <= 57) {
				return false
			}
		}
	} else {
		return false
	}
	return true
}

//ad hoc
//конечный автомат
