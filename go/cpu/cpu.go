package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	filename := os.Args[1]
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	mainloop(strings.TrimSuffix(string(dat), "\n"))

	//fmt.Println(btoint(sub(inttob(2, 16), inttob(1, 16))))
	//fmt.Println(btoint(subc(inttob(373, 16), inttob(185, 16))))
	//subcn(inttob(425, 16), inttob(400, 16))
}

func pow(a int, b int) int {
	if b == 0 {
		return 1
	}
	return a * pow(a, b-1)
}

func btoint_helper(a string) int {
	bit := len(a)
	res := 0
	for i := range bit {
		if a[i] == '1' {
			res = res + pow(2, (bit-i-1))
		}
	}
	return res
}

func btoint(a string) int {
	if a[0] == '1' {
		res_positive := two_complement(a)
		return -btoint_helper(res_positive)
	}
	return btoint_helper(a)
}

// TODO: account for sign bit
func inttob(a int, bit int) string {
	res := make([]byte, bit)
	for i := range bit {
		curr := pow(2, bit-i-1)
		if curr <= a {
			res[i] = '1'
			a = a - curr
		} else {
			res[i] = '0'
		}
	}
	return string(res)
}

func append0s(a string) string {
	bit := 16
	alen := len(a)
	res := make([]byte, bit)
	for i := range bit {
		if i >= alen {
			res[bit-i-1] = '0'
		} else {
			res[bit-i-1] = a[alen-i-1]
		}
	}
	return string(res)
}

func add(a string, b string) string {
	carry := false
	bit := len(a)
	res := make([]byte, bit)

	if len(b) != bit {
		panic("add: size must be the same")
	}

	for i := range bit {
		a_digit := a[bit-i-1]
		b_digit := b[bit-i-1]

		if a_digit == '1' && b_digit == '1' && carry {
			res[bit-i-1] = '1'
			carry = true
		} else if a_digit == '1' && b_digit == '1' {
			res[bit-i-1] = '0'
			carry = true
		} else if (a_digit == '0' && b_digit == '1' || a_digit == '1' && b_digit == '0') && carry {
			res[bit-i-1] = '0'
			carry = true
		} else if a_digit == '0' && b_digit == '1' || a_digit == '1' && b_digit == '0' {
			res[bit-i-1] = '1'
			carry = false
		} else if a_digit == '0' && b_digit == '0' {
			if carry {
				res[bit-i-1] = '1'
			} else {
				res[bit-i-1] = '0'
			}
			carry = false
		} else {
			panic("add: no conditions were met")
		}
	}
	return string(res)
}

func sub(a string, b string) string {
	carry := false
	bit := len(a)
	res := make([]byte, bit)

	if len(b) != bit {
		panic("sub: size must be the same")
	}

	for i := range bit {
		a_digit := a[bit-i-1]
		b_digit := b[bit-i-1]

		if a_digit == '1' && b_digit == '1' && carry {
			res[bit-i-1] = '1'
			carry = true
		} else if a_digit == '1' && b_digit == '1' {
			res[bit-i-1] = '0'
			carry = false
		} else if a_digit == '0' && b_digit == '1' && carry {
			res[bit-i-1] = '0'
			carry = true
		} else if a_digit == '0' && b_digit == '1' {
			res[bit-i-1] = '1'
			carry = true
		} else if a_digit == '1' && b_digit == '0' && carry {
			res[bit-i-1] = '0'
			carry = false
		} else if a_digit == '1' && b_digit == '0' {
			res[bit-i-1] = '1'
			carry = false
		} else if a_digit == '0' && b_digit == '0' {
			if carry {
				res[bit-i-1] = '1'
				carry = true
			} else {
				res[bit-i-1] = '0'
				carry = false
			}
		} else {
			panic("sub: no conditions were met")
		}
	}
	return string(res)
}

func increment(a string) string {
	bit := len(a)
	one := make([]byte, bit)
	for i := range bit {
		if i == bit-1 {
			one[i] = '1'
		} else {
			one[i] = '0'
		}
	}
	return add(a, string(one))
}

func two_complement(a string) string {
	bit := len(a)
	res := make([]byte, bit)

	for i := range bit {
		a_digit := a[i]
		if a_digit == '1' {
			res[i] = '0'
		} else {
			res[i] = '1'
		}
	}
	return increment(string(res))
}

func subc(a string, b string) string {
	bit := len(a)

	if len(b) != bit {
		panic("sub: size must be the same")
	}

	twoc := two_complement(b)

	return add(a, twoc)
}

func subcn(a string, b string) string {
	bit := len(a)

	if len(b) != bit {
		panic("sub: size must be the same")
	}

	resc := subc(a, b)

	if resc[0] == '1' {
		ressigned := two_complement(resc)
		fmt.Printf("-%d\n", btoint(ressigned))
		return ressigned
	} else {
		fmt.Println(btoint(resc))
		return resc
	}
}

func make_zero(length int) string {
	str := ""
	for range length {
		str = str + "0"
	}
	return str
}

func mul(a string, b string) string {
	X := 0
	bit := len(a)
	temp := make([]byte, bit)
	res := make_zero(16)
	for i := range bit {
		for j := range bit {
			if bit-j-1-X < 0 {
				break
			}
			if a[bit-j-1] == '1' && b[bit-i-1] == '1' {
				temp[bit-j-1-X] = '1'
			} else {
				temp[bit-j-1-X] = '0'
			}
		}
		X++
		res = add(res, string(temp))
		for j := range bit {
			temp[j] = '0'
		}
	}
	return res
}

func div(a string, b string) string {
	/*
				Step 1: Divide the bits of the dividend and record the quotient.
				Step 2: Multiply the divisor by the quotient and write the product.
				Step 3: Subtract the product from the dividend and write the difference.
				Step 4: Bring down the next digit and repeat.
				LOAD X
		 		ADD Y
		 		HALT
				X DATA 32
				Y DATA 33
	*/
	return " "
}

func iszero(a string) bool {
	for i := range len(a) {
		if a[i] == '1' {
			return false
		}
	}
	return true
}

func decode(word string) (op string, addrmode string, addrfield string) {
	return word[:4], word[4:6], word[6:]
}

func instruction_to_mnemonic(word string) string {
	res := ""
	op, addrmode, addrfield := decode(word)

	switch op {
	case "0000":
		res = "HALT"
		return res
	case "0001":
		res = "LOAD "
	case "0010":
		res = "STORE "
	case "0011":
		res = "CALL"
	case "0100":
		res = "BR "
	case "0101":
		res = "BREQ "
	case "0110":
		res = "BRGE"
	case "0111":
		res = "BRLT "
	case "1000":
		res = "ADD "
	case "1001":
		res = "SUB "
	case "1010":
		res = "MUL "
	case "1011":
		res = "DIV "
	}

	switch addrmode {
	case "01":
		res = res + "="
	case "10":
		res = res + "$"
	case "11":
		res = res + "@"
	}

	return res + strconv.Itoa(btoint(addrfield))
}

func print_mem(memory []string) {
	for i := range len(memory) {
		if memory[i] != "0000000000000000" {
			fmt.Printf("[%d] = %s\n", i, memory[i])
		}
	}
}

func mainloop(data string) int {
	loop := true

	var addrfield_mem = make([]string, 1024)
	for i := range len(addrfield_mem) {
		addrfield_mem[i] = "0000000000000000"
	}
	data_arr := strings.Split(data, "\n")

	fmt.Println(data_arr)

	i := 0
	for i < len(data_arr) {
		org_address := btoint(data_arr[i])
		word_count := btoint(data_arr[i+1])
		fmt.Println("word count " + strconv.Itoa(word_count))
		for j := range word_count - 1 {
			fmt.Println(i + 2 + j)
			addrfield_mem[org_address+j] = data_arr[i+2+j]
		}
		i = i + word_count + 2
	}

	AC := "0000000000000000"
	PC := "0000000000"
	MAR := "0000000000"
	MBR := "0000000000000000"
	IR := "0000000000000000"

	/* addrfield_mem[0] = mnemonic_to_instruction("LOAD @3")
	// addrfield_mem[1] = mnemonic_to_instruction("ADD 4")
	// addrfield_mem[2] = mnemonic_to_instruction("HALT")
	// addrfield_mem[3] = inttob(5, 16)  // 5
	// addrfield_mem[4] = inttob(13, 16) // 13
	// addrfield_mem[5] = inttob(42, 16) // 42

	// addrfield_mem[0] = "0001" + "01" + "0000000100" // LOAD =4
	// addrfield_mem[1] = "1001" + "01" + "0000000001" // SUB =1
	// addrfield_mem[2] = "0101" + "00" + "0000000100" // BREQ 4
	// addrfield_mem[3] = "0100" + "00" + "0000000001" // BR 1
	// addrfield_mem[4] = "0000" + "00" + "0000000000" // HALT

	// addrfield_mem[0] = "0001" + "01" + "0000000011" // LOAD =3
	// addrfield_mem[1] = "0010" + "00" + "0000000101" // STORE 5
	// addrfield_mem[2] = "0001" + "01" + "0000001101" // LOAD =13
	// addrfield_mem[3] = "0001" + "00" + "0000000101" // LOAD 5
	// addrfield_mem[4] = "0000" + "00" + "0000000000" // HALT

	// addrfield_mem[0] = mnemonic_to_instruction("CALL 3")
	// addrfield_mem[1] = mnemonic_to_instruction("HALT")
	// addrfield_mem[4] = mnemonic_to_instruction("LOAD =4")
	// addrfield_mem[5] = mnemonic_to_instruction("STORE 8")
	// addrfield_mem[6] = mnemonic_to_instruction("MUL 8")
	// addrfield_mem[7] = mnemonic_to_instruction("BR @3") */

	for loop {
		IR = addrfield_mem[btoint(PC)]
		op, addrmode, addrfield := decode(IR)

		switch addrmode {
		case "00": // Direct
			MAR = addrfield
		case "01": // Immediate
			MAR = addrfield
			MBR = append0s(MAR)
		case "10": // Indexed
			// MAR = addrfield
			// MAR = add(MAR, XR)
			panic("not implemented")
		case "11": // Indirect
			MAR = addrfield
			MBR = addrfield_mem[btoint(MAR)]
			MAR = MBR[6:]
		}

		PC = add(PC, "0000000001")

		switch op {
		case "0000": // HALT
			loop = false
		case "0001": // LOAD
			if addrmode != "01" {
				MBR = addrfield_mem[btoint(MAR)]
			}
			AC = MBR
		case "0010": // STORE
			MBR = AC
			addrfield_mem[btoint(MAR)] = MBR
		case "0011": // CALL
			PC = MBR
			addrfield_mem[btoint(MAR)] = PC
			PC = MAR
			PC = add(PC, "0000000001")
		case "0100": // BR
			PC = MAR
		case "0101": // BREQ
			if iszero(AC) {
				PC = MAR
			}
		case "0110": // BRGE
			if AC[0] == '0' || iszero(AC) {
				PC = MAR
			}
		case "0111": // BRLT
			if AC[0] == '1' {
				PC = MAR
			}
		case "1000": // ADD
			if addrmode != "01" {
				MBR = addrfield_mem[btoint(MAR)]
			}
			AC = add(AC, MBR)
		case "1001": // SUB
			if addrmode != "01" {
				MBR = addrfield_mem[btoint(MAR)]
			}
			AC = subc(AC, MBR)
		case "1010": // MUL
			if addrmode != "01" {
				MBR = addrfield_mem[btoint(MAR)]
			}
			AC = mul(AC, MBR)
		case "1011": // DIV
			panic("not implemented")
		}

		fmt.Printf("AC  = %s (%d) | ", AC, btoint(AC))
		fmt.Printf("PC  = %s       (%d) | ", PC, btoint(PC))
		//fmt.Printf("MAR = %s       (%d)\n", MAR, btoint(MAR))
		//fmt.Printf("MBR = %s (%d)\n", MBR, btoint(MBR))
		fmt.Printf("IR  = %s (%s)\n", IR, instruction_to_mnemonic(IR))
	}

	fmt.Println(btoint(AC))
	return 0
}
