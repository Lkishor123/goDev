package datastruct

import "fmt"

type Node struct {
	val  int
	next *Node
}

func AddtoLL(head *Node, val int) {
	ptr := head
	for {
		if ptr.next == nil {
			n1 := Node{}
			n1.val = val
			ptr.next = &n1
			break
		} else {
			ptr = ptr.next
		}
	}
}

func PrintLL(head *Node) {
	ptr := head.next
	for {
		if ptr.next == nil {
			fmt.Printf("(%v)-> nil ", ptr.val)
			break
		} else {
			fmt.Printf("(%v)-> ", ptr.val)
			ptr = ptr.next
		}
	}
}

func LLmain() {
	head := Node{}
	AddtoLL(&head, 1)
	AddtoLL(&head, 2)
	AddtoLL(&head, 3)
	AddtoLL(&head, 4)

	PrintLL(&head)
}
