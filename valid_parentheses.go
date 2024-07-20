package main

// # VALID PARENTHESES
//
// Given a string s containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.
//
// An input string is valid if:
//
// Open brackets must be closed by the same type of brackets.
// Open brackets must be closed in the correct order.
// Every close bracket has a corresponding open bracket of the same type.
//
// Example 1:
//
// Input: s = "()"
// Output: true
// Example 2:
//
// Input: s = "()[]{}"
// Output: true
// Example 3:
//
// Input: s = "(]"
// Output: false
func isValid(s string) bool {
	st := Stack{}

	for _, r := range s {
		top := st.Top()
		if r == '{' || r == '[' || r == '(' {
			st.Push(r)
		} else if r == '}' && top == '{' {
			st.Pop()
		} else if r == ']' && top == '[' {
			st.Pop()
		} else if r == ')' && top == '(' {
			st.Pop()
		} else {
			return false
		}
	}
	if st.IsEmpty() {
		return true
	}
	return false
}

type Stack []rune

func (st *Stack) Push(r rune) {
	*st = append(*st, r)
}

func (st *Stack) IsEmpty() bool {
	return len(*st) == 0
}

func (st *Stack) Pop() rune {
	if st.IsEmpty() {
		return rune(0)
	}
	top := (*st)[len(*st)-1]
	*st = (*st)[:len(*st)-1]
	return top
}

func (st *Stack) Top() rune {
	if st.IsEmpty() {
		return rune(0)
	}
	return (*st)[len(*st)-1]
}
