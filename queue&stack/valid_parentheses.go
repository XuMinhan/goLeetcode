package main

import "container/list"

func main0() {
	println(isValid("()"))
}
func isValid(s string) bool {
	bytes := []byte(s)
	l := list.New()
	for _, char := range bytes {
		if char == '(' || char == '{' || char == '[' {
			l.PushBack(char)
		} else {
			if l.Len() == 0 {
				return false // 遇到右括号但栈为空
			}
			last := l.Back().Value.(byte) // 访问链表中的最后一个元素
			if (char == ')' && last == '(') || (char == '}' && last == '{') || (char == ']' && last == '[') {
				l.Remove(l.Back()) // 配对成功，移除栈顶元素
			} else {
				return false // 不匹配的情况
			}
		}
	}
	return l.Len() == 0 // 如果栈为空，表示所有括号匹配成功
}
