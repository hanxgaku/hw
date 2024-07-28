package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	inputfile := os.Args[1]
	outputfile := os.Args[2]
	fi, err := os.ReadFile(inputfile)
	if err != nil {
		panic(err)
	}

	instruction_arr := initialize_asm(string(fi))

	fo, err := os.Create(outputfile)
	if err != nil {
		panic(err)
	}

	defer fo.Close()

	for i := range len(instruction_arr) {
		fo.WriteString(instruction_arr[i] + "\n")
	}
}

func pow(a int, b int) int {
	if b == 0 {
		return 1
	}
	return a * pow(a, b-1)
}

func inttob(a int, bit int) string {
	res := make([]byte, bit)
	if a < 0 {
		res[0] = '1'
		a = -a
	} else {
		res[0] = '0'
	}
	for i := 1; i < bit; i++ {
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

func mnemonic_to_instruction(word string, labelmap map[string]int) string {
	reg := `^\w?\s+((?P<Op>(LOAD|STORE|CALL|BR|BREQ|BRGE|BRLT|ADD|SUB|MUL|DIV))\s+(?P<AddrMode>[=$@]?)(?P<Label>\w|(?P<AddrField>\d+)))|(?P<Halt>HALT)$`
	r := regexp.MustCompile(reg)
	matches := r.FindStringSubmatch(word)

	opindex := r.SubexpIndex("Op")
	addrmodeindex := r.SubexpIndex("AddrMode")
	addrfieldindex := r.SubexpIndex("AddrField")
	labelindex := r.SubexpIndex("Label")
	haltindex := r.SubexpIndex("Halt")

	//fmt.Println(matches)
	if len(matches) == 0 {
		fmt.Println(word)
		panic("syntax error")
	}

	if len(matches[haltindex]) != 0 {
		return "0000000000000000"
	} else {
		//fmt.Println(matches)
		//fmt.Println(opindex)
		//fmt.Println(matches[3])

		op := matches[opindex]
		addrmode := matches[addrmodeindex]
		addrfield := matches[addrfieldindex]

		res := ""
		switch op {
		case "LOAD":
			res = res + "0001"
		case "STORE":
			res = res + "0010"
		case "CALL":
			res = res + "0011"
		case "BR":
			res = res + "0100"
		case "BREQ":
			res = res + "0101"
		case "BRGE":
			res = res + "0110"
		case "BRLT":
			res = res + "0111"
		case "ADD":
			res = res + "1000"
		case "SUB":
			res = res + "1001"
		case "MUL":
			res = res + "1010"
		case "DIV":
			res = res + "1011"
		}

		if addrmode == "=" {
			res = res + "01"
		} else if addrmode == "$" {
			res = res + "10"
		} else if addrmode == "@" {
			res = res + "11"
		} else {
			res = res + "00"
		}

		label := matches[labelindex]
		var i int
		if len(label) != 0 {
			i = labelmap[label]
		} else {
			i, _ = strconv.Atoi(addrfield)
		}
		return res + inttob(i, 10)
	}
}

func init_labels(data_arr []string) map[string]int {
	res := make(map[string]int)

	for i := range len(data_arr) {
		line := data_arr[i]

		reg := `^(?P<Label>\w)\s+.*$`
		r := regexp.MustCompile(reg)
		match := r.FindStringSubmatch(line)
		if len(match) != 0 {
			fmt.Println("found match")

			label_index := r.SubexpIndex("Label")
			label := match[label_index]
			res[label] = i
		}
	}
	return res
}

// implement END and BSS
func initialize_asm(data string) []string {
	var address_block []string
	var org_block []string
	var org_address string
	var end_address string
	data_arr := strings.Split(data, "\n")
	label_map := init_labels(data_arr)
	fmt.Println(label_map)
	i := 0
	for i < len(data_arr) {
		line := data_arr[i]

		regdata := `^(?P<Label>\w?)(\s+DATA)\s+(?P<Numbers>(\d+(,\d+)*))$`
		r1 := regexp.MustCompile(regdata)
		data_matches := r1.FindStringSubmatch(line)

		regorg := `^(?P<Label>\w?)(\s+ORG)\s+(?P<AddrField>\d+)$`
		r2 := regexp.MustCompile(regorg)
		org_matches := r2.FindStringSubmatch(line)

		regend := `^(?P<Label>\w?)(\s+END)\s+(?P<AddrField>\d+)$`
		r3 := regexp.MustCompile(regend)
		end_matches := r3.FindStringSubmatch(line)

		if len(data_matches) != 0 {
			numbers_index := r1.SubexpIndex("Numbers")
			numbers := strings.Split(data_matches[numbers_index], ",")

			fmt.Println(numbers)
			for j := range len(numbers) {
				number, _ := strconv.Atoi(numbers[j])
				org_block = append(org_block, inttob(number, 16))
				i++
			}
		} else if len(org_matches) != 0 {
			if len(org_block) != 0 {
				address_block = append(address_block, org_address)
				address_block = append(address_block, inttob(len(org_block), 10))
				address_block = append(address_block, org_block...)
				org_block = nil
			}

			address_index := r2.SubexpIndex("AddrField")
			org_address = org_matches[address_index]
			i++
		} else if len(end_matches) != 0 {
			address_index := r3.SubexpIndex("AddrField")
			end_address = end_matches[address_index]
			i++
		} else {
			org_block = append(org_block, mnemonic_to_instruction(line, label_map))
			i++
		}
	}
	end_address_i, _ := strconv.Atoi(end_address)
	org_address_i, _ := strconv.Atoi(org_address)
	address_block = append(address_block, inttob(end_address_i, 10))
	address_block = append(address_block, inttob(org_address_i, 10))
	address_block = append(address_block, inttob(len(org_block), 10))
	address_block = append(address_block, org_block...)
	fmt.Println(address_block)
	return address_block
}
