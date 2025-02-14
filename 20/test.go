package main

type Node struct {
	data rune
	next *Node
}

type LinkedList struct {
	head *Node
}

func (list *LinkedList) pop() {
	if list.head == nil || list.head.next == nil {
		list.head = nil
		return
	}

	current := list.head
	for current != nil {
		if current.next.next == nil {
			current.next = nil
			return
		}
		current = current.next
	}
}

func (list *LinkedList) push(value rune) {
	newNode := &Node{data: value, next: nil}
	if list.head == nil {
		list.head = newNode
		return
	}
	current := list.head
	for current != nil {
		if current.next == nil {
			current.next = newNode
			return
		}
		current = current.next
	}
}

func (list *LinkedList) peek() rune {
	if list.head == nil {
		return ' '
	}
	current := list.head
	for current != nil {
		if current.next == nil {
			return current.data
		}
		current = current.next
	}
	return ' '
}

func isValid(s string) bool {
	rune := []rune(s)
	if len(rune) == 1 {
		return false
	}
	if len(rune)%2 != 0 {
		return false
	}
	list := &LinkedList{head: nil}
	for i := range rune {
		if rune[i] == '{' || rune[i] == '(' || rune[i] == '[' {
			list.push(rune[i])
		} else {
			currentrune := list.peek()
			if rune[i] == '}' {
				if currentrune == '{' {
					list.pop()
				} else {
					return false
				}
			}
			if rune[i] == ')' {
				if currentrune == '(' {
					list.pop()
				} else {
					return false
				}
			}
			if rune[i] == ']' {
				if currentrune == '[' {
					list.pop()
				} else {
					return false
				}
			}
		}
	}
	if list.peek() != ' ' {
		return false
	}
	return true
}
