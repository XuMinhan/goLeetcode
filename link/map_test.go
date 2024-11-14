package main

import "fmt"

func main() {
	// 创建一个 map
	fruits := make(map[string]int)

	// 添加键值对
	fruits["apple"] = 5
	fruits["banana"] = 3
	fruits["orange"] = 8

	// 访问和更新值
	fmt.Println("Apple count:", fruits["apple"]) // 输出: Apple count: 5
	fruits["apple"] = 10
	fmt.Println("Updated apple count:", fruits["apple"]) // 输出: Updated apple count: 10

	// 检查键是否存在
	if val, ok := fruits["banana"]; ok {
		fmt.Println("Banana count:", val) // 输出: Banana count: 3
	} else {
		fmt.Println("Banana not found")
	}

	// 删除一个键值对
	delete(fruits, "banana")

	// 遍历 map
	for key, value := range fruits {
		fmt.Printf("%s: %d\n", key, value)
	}
}
