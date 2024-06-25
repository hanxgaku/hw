package main

import (
	"fmt"
)

type Pair struct {
	p1 int
	p2 int
}

type Node struct {
	data int
	next *Node
}

var envg map[string]interface{}
var env map[string]interface{}

func main() {
	envg = make(map[string]interface{})
	env = make(map[string]interface{})

	/* a := 1
	// apointer := &a
	// fmt.Println(*apointer)
	// p := Pair{1, 2}
	// fmt.Println(p)
	// n3 := Node{3, nil}
	// n2 := Node{2, &n3}
	// n1 := Node{1, &n2}
	// fmt.Println(n1.data)
	// fmt.Println((*n1.next).data)
	// fmt.Println((*(*n1.next).next).data)
	// printLL(&n1)
	// fmt.Println(lastElement(&n1))
	//delLast(&n1)
	//printLL(&n1)
	// c := []int{1, 2, 3}
	// b := []int{4, 5, 7}
	// l1 := arrayToLinkedList(c)
	// l2 := arrayToLinkedList(b)
	// l1 = concat(l1, l2)
	// printLL(l1)
	// l1 = insertElementSorted(6, l1)
	// printLL(l1)
	// d := []int{4, 2, 3, 1}
	// l3 := arrayToLinkedList(d)
	// printLL(l3)
	// printLL(sortLL(l3))

	// fmt.Println("----")
	// d2 := []int{5, 1, 7, 3, 2, 9, 6}
	// l4 := arrayToLinkedList(d2)
	// printLL(l4)
	// printLL(qsort(l4))

	// fmt.Println("----")
	// printLL(l4)
	// div1, div2 := divide(l4)
	// printLL(div1)
	// printLL(div2)

	// fmt.Println("-- msort --")
	// sortedarr1 := []int{1, 3, 5, 7}
	// sortedarr2 := []int{2, 4, 6}
	// sortedlist1 := arrayToLinkedList(sortedarr1)
	// sortedlist2 := arrayToLinkedList(sortedarr2)
	// printLL(msort_helper(sortedlist1, sortedlist2))

	// notsorted := arrayToLinkedList([]int{2, 1, 5, 4, 3})
	// printLL(msort(notsorted)) */
}

func printLL(n *Node) {
	fmt.Printf("{")
	for n != nil {
		fmt.Printf("%d ", (*n).data)
		n = n.next
	}
	fmt.Printf("}\n")
}

func insertElementSorted(k int, n *Node) *Node {
	ninit := n
	m := 0
	for true {
		if n == nil || (k < n.data && m == 0) {
			ninit = &Node{k, n}
			break
		} else if k >= n.data && n.next == nil {
			n.next = &Node{k, nil}
			break
		} else if k < n.next.data {
			n.next = &Node{k, n.next}
			break
		}
		n = n.next
		m++
	}
	return ninit
}

func sortLL(n *Node) *Node {
	var sorted *Node = nil
	for n != nil {
		sorted = insertElementSorted(n.data, sorted)
		n = n.next
	}
	return sorted
}

func lastElement(n *Node) int {
	if (*n).next == nil {
		return (*n).data
	}
	return lastElement((*n).next)
}

func delLast(n *Node) {
	for true {
		if (*(*n).next).next == nil {
			(*n).next = nil
			return
		}
		n = (*n).next
	}
}

func concat(n1 *Node, n2 *Node) *Node {
	ninit := n1
	if n1 == nil {
		return n2
	} else {
		for n1.next != nil {
			n1 = n1.next
		}
		n1.next = n2
	}
	return ninit
}

func arrayToLinkedList(a []int) *Node {
	var n *Node = nil
	for i := 0; i < len(a); i++ {
		n = &Node{a[len(a)-i-1], n}
	}
	return n
}

func qsort_helper(n *Node, k int, op func(int, int) bool) *Node {
	var list *Node = nil
	for n != nil {
		if op(n.data, k) {
			list = &Node{n.data, list}
		}
		n = n.next
	}
	return list
}

func qsort(n *Node) *Node {
	if n == nil {
		return n
	}
	mid := &Node{n.data, nil}
	return concat(qsort(qsort_helper(n, n.data, func(k int, m int) bool { return k < m })), concat(mid, qsort(qsort_helper(n, n.data, func(k int, m int) bool { return k > m }))))
}

func divide(n *Node) (*Node, *Node) {
	k := 0
	var n1 *Node = nil
	var n2 *Node = nil
	for n != nil {
		if k%2 == 0 {
			n1 = &Node{n.data, n1}
		} else {
			n2 = &Node{n.data, n2}
		}
		n = n.next
		k++
	}
	return n1, n2
}

func msort_helper(n1 *Node, n2 *Node) *Node {
	var sorted *Node = nil
	if n1.data >= n2.data {
		sorted = n2
		n2 = n2.next
	} else {
		sorted = n1
		n1 = n1.next
	}
	var init *Node = sorted
	for true {
		if n1 == nil {
			sorted.next = n2
			break
		} else if n2 == nil {
			sorted.next = n1
			break
		} else if n1.data >= n2.data {
			sorted.next = &Node{n2.data, nil}
			n2 = n2.next
		} else {
			sorted.next = &Node{n1.data, nil}
			n1 = n1.next
		}
		sorted = sorted.next
	}
	return init
}

func msort(n *Node) *Node {
	if n.next == nil {
		return n
	}
	div1, div2 := divide(n)
	return msort_helper(msort(div1), msort(div2))
}

func add1(x int) int {
	return x + 1
}

func mul2(x int) int {
	return 2 * x
}

func compose(f func(int) int, g func(int) int) func(int) int {
	return func(x int) int {
		return f(g(x))
	}
}
