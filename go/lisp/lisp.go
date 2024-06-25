package main

import (
	"fmt"
	"os"
	"strconv"
)

type LispNode struct {
	data interface{}
	next *LispNode
}

type Closure struct {
	lmd   *LispNode
	scope map[string]interface{}
}

var envg map[string]interface{}
var env map[string]interface{}

func main() {
	envg = make(map[string]interface{})
	env = make(map[string]interface{})

	filename := os.Args[1]
	dat, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	printLispList(evaluate(parseTokensToLispList(tokenize(string(dat))), env))

}

func printLispListHelper1(i interface{}) {
	switch i := i.(type) {
	case int:
		fmt.Printf("%d", i)
	case string:
		fmt.Printf(i)
	case bool:
		fmt.Printf(strconv.FormatBool(i))
	case *LispNode:
		printLispListHelper2(i)
	case Closure:
		fmt.Printf("{closure}")
	}
}

func printLispListHelper2(lispNode *LispNode) {
	fmt.Printf("(")
	for lispNode != nil {
		printLispListHelper1(lispNode.data)
		if lispNode.next != nil {
			fmt.Printf(" ")
		}
		lispNode = lispNode.next
	}
	fmt.Printf(")")
}

func printLispList(i interface{}) {
	printLispListHelper1(i)
	fmt.Printf("\n")
}

func tokenize(s string) []string {
	op := 0
	cp := 0
	arr := []string{}
	word := ""
	comment := false
	for i := 0; i < len(s); i++ {
		if comment {
			if s[i] == '\n' {
				comment = false
			}
			continue
		}
		if s[i] == ';' {
			comment = true
			continue
		}
		if s[i] == '(' || s[i] == ')' {
			if s[i] == '(' {
				op++
			} else if s[i] == ')' {
				cp++
			}
			if word != "" {
				arr = append(arr, word)
			}
			word = string(s[i])
			arr = append(arr, word)
			word = ""
		} else if (s[i] == ' ' || s[i] == '\n' || s[i] == '\t') && word != "" {
			arr = append(arr, word)
			word = ""
		} else if (s[i] == ' ' || s[i] == '\n' || s[i] == '\t') && word == "" {
			continue
		} else {
			word += string(s[i])
		}
	}
	if op != cp {
		panic("uneven parens")
	}
	if word != "" {
		arr = append(arr, word)
	}
	return arr
}

func isNumeral(s string) bool {
	for i := 0; i < len(s); i++ {
		if !(s[i] <= 57 && s[i] >= 48) {
			return false
		}
	}
	return true
}

func parseTokensToLispList(tokenArr []string) interface{} {
	var stack *LispNode = nil
	for i := 0; i < len(tokenArr); i++ {
		if tokenArr[i] == ")" {
			var sublist *LispNode = nil
			for stack.data != "(" {
				sublist = &LispNode{stack.data, sublist}
				stack = stack.next
			}
			stack = stack.next
			stack = &LispNode{sublist, stack}
		} else {
			if isNumeral(tokenArr[i]) {
				intVal, err := strconv.Atoi(tokenArr[i])
				if err != nil {
					panic(err)
				}
				stack = &LispNode{intVal, stack}
			} else {
				stack = &LispNode{tokenArr[i], stack}
			}
		}
	}
	return stack.data
}

func mapCopy(old map[string]interface{}) map[string]interface{} {
	newMap := make(map[string]interface{})
	for key, value := range old {
		newMap[key] = value
	}
	return newMap
}

func printMap(env map[string]interface{}) {
	for key, value := range env {
		fmt.Printf("key: %s, val: ", key)
		printLispList(value)
	}
}

func listlen(i interface{}) int {
	count := 0
	l, ok := i.(*LispNode)
	if !ok {
		panic("cannot take length of non-list")
	}
	for l != nil {
		count++
		l = l.next
	}
	return count
}

func listeq(l1 *LispNode, l2 *LispNode) bool {
	if l1 == nil && l2 == nil {
		return true
	}
	switch d1 := l1.data.(type) {
	case int:
		d2, ok := l2.data.(int)
		if !ok || d1 != d2 {
			return false
		}
	case string:
		d2, ok := l2.data.(string)
		if !ok || d1 != d2 {
			return false
		}
	case *LispNode:
		d2, ok := l2.data.(*LispNode)
		if !ok || !listeq(d1, d2) {
			return false
		}
	}
	return listeq(l1.next, l2.next)
}

func evaluate(exp interface{}, env map[string]interface{}) interface{} {
	printLispList(exp)
	printMap(env)
	fmt.Println("")
	switch i := exp.(type) {
	case int:
		return i
	case string:
		if i == "t" {
			return true
		} else if i == "f" {
			return false
		}
		val1, ok1 := env[i]
		val2, ok2 := envg[i]
		if ok1 {
			return val1
		} else if ok2 {
			return val2
		} else {
			panic(fmt.Sprintf("variable %s undefined", i))
		}
	case *LispNode:
		switch i.data {
		case "+":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 + val2
			} else {
				panic("non-int values to operand")
			}
		case "-":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 - val2
			} else {
				panic("non-int values to operand")
			}
		case "*":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 * val2
			} else {
				panic("non-int values to operand")
			}
		case "/":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 / val2
			} else {
				panic("non-int values to operand")
			}
		case ">":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 > val2
			} else {
				panic("non-int values to operand")
			}
		case "<":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 < val2
			} else {
				panic("non-int values to operand")
			}
		case ">=":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 >= val2
			} else {
				panic("non-int values to operand")
			}
		case "<=":
			val1, ok1 := evaluate(i.next.data, env).(int)
			val2, ok2 := evaluate(i.next.next.data, env).(int)
			if ok1 && ok2 {
				return val1 <= val2
			} else {
				panic("non-int values to operand")
			}
		case "=":
			res1 := evaluate(i.next.data, env)
			res2 := evaluate(i.next.next.data, env)
			val1, ok1 := res1.(int)
			val2, ok2 := res2.(int)
			val3, ok3 := res1.(string)
			val4, ok4 := res2.(string)
			val5, ok5 := res1.(*LispNode)
			val6, ok6 := res2.(*LispNode)
			if ok1 && ok2 {
				return val1 == val2
			} else if ok3 && ok4 {
				return val3 == val4
			} else if ok5 && ok6 {
				return listeq(val5, val6)
			} else {
				return false
			}
		case "/=":
			res1 := evaluate(i.next.data, env)
			res2 := evaluate(i.next.next.data, env)
			val1, ok1 := res1.(int)
			val2, ok2 := res2.(int)
			val3, ok3 := res1.(string)
			val4, ok4 := res2.(string)
			if ok1 && ok2 {
				return val1 != val2
			} else if ok3 && ok4 {
				return val3 != val4
			} else {
				panic("bad values to operand /=")
			}
		case "quote":
			return i.next.data
		case "car":
			val, ok := evaluate(i.next.data, env).(*LispNode)
			if !ok {
				panic("car argument is not a list")
			}
			if val == nil {
				panic("cannot take car from null")
			}
			return val.data
		case "cdr":
			val, ok := evaluate(i.next.data, env).(*LispNode)
			if !ok {
				panic("cdr argument is not a list")
			}
			if val == nil {
				panic("cannot take cdr from null")
			}
			return val.next
		case "cons":
			val1 := evaluate(i.next.data, env)
			val2, ok := evaluate(i.next.next.data, env).(*LispNode)
			if !ok {
				panic("cons: 2nd argument must be list")
			}
			return &LispNode{val1, val2}
		case "list":
			var newListRev *LispNode = nil
			var newList *LispNode = nil
			i = i.next
			for i != nil {
				newListRev = &LispNode{evaluate(i.data, env), newListRev}
				i = i.next
			}
			for newListRev != nil {
				newList = &LispNode{newListRev.data, newList}
				newListRev = newListRev.next
			}
			return newList
		case "null":
			val1, ok := evaluate(i.next.data, env).(*LispNode)
			if ok {
				if val1 == nil {
					return true
				}
				return false
			} else {
				return false
			}

		case "if":
			condition, ok := evaluate(i.next.data, env).(bool)
			if !ok {
				panic("if: bad condition")
			}
			if condition {
				return evaluate(i.next.next.data, env)
			} else {
				return evaluate(i.next.next.next.data, env)
			}
		case "cond":
			i = i.next
			for i != nil {
				branch := i.data.(*LispNode)
				condition := evaluate(branch.data, env).(bool)
				if condition {
					return evaluate(branch.next.data, env)
				}
				i = i.next
			}
			panic("cond: must have valid branch")
		case "let":
			vars := i.next.data.(*LispNode)
			newEnv := mapCopy(env)
			for vars != nil {
				assignment := vars.data.(*LispNode)
				varName := assignment.data.(string)
				varVal := evaluate(assignment.next.data, env)
				newEnv[varName] = varVal
				vars = vars.next
			}
			return evaluate(i.next.next.data, newEnv)
		case "define":
			if listlen(i) != 3 {
				panic("wrong number of args to define func")
			}
			varName := i.next.data.(string)
			varVal := evaluate(i.next.next.data, env)
			envg[varName] = varVal
			return varName
		case "setq":
			varName := i.next.data.(string)
			varVal := evaluate(i.next.next.data, env)
			if env[varName] != nil {
				env[varName] = varVal
			} else {
				envg[varName] = varVal
			}
			return varName
		case "print":
			result := evaluate(i.next.data, env)
			printLispList(result)
			return result
		case "begin":
			i = i.next
			for i != nil {
				if i.next == nil {
					return evaluate(i.data, env)
				}
				evaluate(i.data, env)
				i = i.next
			}
		case "closure?":
			_, ok := evaluate(i.next.data, env).(Closure)
			return ok
		case "number?":
			_, ok := evaluate(i.next.data, env).(int)
			return ok
		case "string?":
			_, ok := evaluate(i.next.data, env).(string)
			return ok
		}
		printLispList(exp)
		closureExp, ok := i.data.(Closure)
		if ok {
			vars := closureExp.lmd.next.data.(*LispNode)
			newEnv := mapCopy(closureExp.scope)
			i = i.next
			if listlen(vars) != listlen(i) {
				printLispList(vars)
				printLispList(i)
				panic(fmt.Sprintf("invalid number of args to lambda expression."))
			}
			for vars != nil {
				varName := vars.data.(string)
				varVal := evaluate(i.data, env)
				newEnv[varName] = varVal
				vars = vars.next
				i = i.next
			}
			return evaluate(closureExp.lmd.next.next.data, newEnv)
		} else {
			execVar, ok1 := i.data.(string)
			lambdaExp, ok2 := i.data.(*LispNode)
			if !ok1 && !ok2 {
				panic("invalid 1st element of expression")
			} else if ok1 && execVar == "lambda" {
				returnClosure := Closure{i, env}
				return returnClosure
				// ((if x car cdr) '(1 2 3 4))
			} else if ok2 {
				vars := lambdaExp.next.data.(*LispNode)
				i = i.next
				if listlen(vars) != listlen(i) {
					printLispList(vars)
					printLispList(i)
					panic(fmt.Sprintf("invalid number of args to lambda expression."))
				}
				for vars != nil {
					varName := vars.data.(string)
					varVal := evaluate(i.data, env)
					env[varName] = varVal
					vars = vars.next
					i = i.next
				}
				return evaluate(lambdaExp.next.next.data, env)
			} else {
				val1, ok1 := env[execVar]
				val2, ok2 := envg[execVar]
				if ok1 {
					newLambda := &LispNode{val1, i.next}
					return evaluate(newLambda, env)
				} else if ok2 {
					newLambda := &LispNode{val2, i.next}
					return evaluate(newLambda, env)
				} else {
					panic(fmt.Sprintf("variable %s undefined", execVar))
				}
			}
		}
	}
	return 0
}
