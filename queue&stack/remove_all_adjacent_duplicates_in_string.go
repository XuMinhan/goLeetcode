package main

import (
	"container/list"
)

func main() {
	println(removeDuplicates("abbaca"))
}
func removeDuplicates(s string) string {
	bytes := []byte(s)
	l := list.New()
	for _, char := range bytes {
		if l.Len() != 0 {
			if l.Back().Value.(byte) == char {
				l.Remove(l.Back())
			} else {
				l.PushBack(char)
			}
		} else {
			l.PushBack(char)
		}
	}
	retString := []byte{}
	for e := l.Front(); e != nil; e = e.Next() {
		retString = append(retString, e.Value.(byte))
	}
	return string(retString)
}
