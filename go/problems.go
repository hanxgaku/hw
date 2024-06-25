package main

import (
    "fmt"
)

var m map[int]string = map[int]string{
    0 : "0",
    1 : "1",
    2 : "2",
    3 : "3",
    4 : "4",
    5 : "5",
    6 : "6",
    7 : "7",
    8 : "8",
    9 : "9",
}

func main() {
    /*
    fmt.Println(str(0))
    fmt.Println(str(666))
    fmt.Println(strtoint("420") + 420)
    fmt.Println('1'+'1')
    fmt.Println(is_identifier2("a123"))
    fmt.Println(is_identifier2("12a"))
    fmt.Println(is_identifier2("a12b13"))
    */
    
    a := tokenize("    (+   12     (*   20 30))3   ")
    for i := 0; i < len(a); i++ {
        fmt.Println(a[i])
    }
}

func str(n int) string {

    
    if (n / 10 == 0) {
        return m[n]
    }
     
    return str(n / 10) + m[n % 10]
}

func pow(n int, m int) int {
    if m == 0 {
     return 1
    }
    return n * pow(n, m - 1)
}

func strtoint(s string) int {
    n := 0
    
    for i := 0; i < len(s); i++ {
       n += int(s[len(s) - i - 1] - '0') * pow(10, i)
    }
    
    return n
}
    
    
//    65 90 97 122
// 354 = 3*10^2+5*10^1+4*10^0


//a123 -- identifier
//a12b34 -- identifier
//12a -- NOT identifier

// ad-hoc
func is_identifier(s string) bool {
    if s == "" {
       return false
    }
    prefix := int(s[0])
    rest := s[1:]
    if prefix >= 65 && prefix <= 90 || prefix >= 97 && prefix <= 122 {
        for i := 0; i < len(rest); i++ {
            if !(int(rest[0]) >= 48 && int(rest[0]) <= 57 || (prefix >= 48 && prefix <= 57)) {
                return false
            }
        }
    } else {
        return false
    }
    return true
}

// finite state machine
func is_identifier2(s string) bool {
    if s == "" {
        return false
    }
    
    n := 0
    for i := 0; i < len(s); i++ {
        prefix := int(s[i])
        if n == 0 {
            if prefix >= 65 && prefix <= 90 || prefix >= 97 && prefix <= 122 {
                n = 1
            } else {
                return false
            }
        } else if n == 1 {
            if (prefix >= 65 && prefix <= 90 || prefix >= 97 && prefix <= 122) || (prefix >= 48 && prefix <= 57) {
                n = 1
            } else {
                return false
            }
        }
    }
    return true
}

//s0 -a-z-> s1 <-a-z-> s1

//"(+ 12(* 45 33))" -> "(","+","12","(","*","45","33",")",")"

//(+ foo bar)

func tokenize(s string) []string {
    arr := []string{}
    word := ""
    for i := 0; i < len(s); i++ {
        if s[i] == '(' || s[i] == ')' {
            if word != "" {
                arr = append(arr, word)
            }
            word = string(s[i])
            arr = append(arr, word)
            word = ""
        } else if s[i] == ' ' && word != "" {
            arr = append(arr, word)
            word = ""
        } else if s[i] == ' ' && word == "" {
            continue
        } else {
            word += string(s[i])
        }
    }
    if word != "" {
        arr = append(arr, word)
    }
    return arr
}