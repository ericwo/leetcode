package main

func main() {}

func IsValid(s string) bool {
	stack := []rune{}

	for i := 0; i < len(s); i++ {
		ch := rune(s[i])
		if ch == '(' || ch == '{' || ch == '[' {
			stack = append(stack, ch)
		} else {
			if len(stack) < 1 {
				return false
			}
			if ch == ')' && stack[len(stack)-1] != '(' {
				return false
			}
			if ch == '}' && stack[len(stack)-1] != '{' {
				return false
			}
			if ch == ']' && stack[len(stack)-1] != '[' {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) < 1
}
