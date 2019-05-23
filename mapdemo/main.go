package main

import "fmt"

func main() {
	// 创建一个映射，键的类型是string，值的类型是int
	// myMap := make(map[string]int)

	//for index, value := range myMap {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}

	// 创建一个映射，键和值的类型都是string
	// 使用两个键值对初始化映射
	// myMap := map[string]string{"Red": "#da1337", "Orange": "#e95a22"}

	// 使用映射字面量声明空映射
	// 创建一个映射，使用字符串切片作为映射的键
	// myMap := map[[]string]int{}

    // 声明一个存储字符串切片的映射
	// 创建一个映射，使用字符串切片作为值
	// myMap := map[int][]string{}


	// 创建一个空映射，用来存储颜色以及颜色对应的十六进制代码
	//myColors := map[string]string{}
	// 将 Red 的代码加入到映射
	//myColors["Red"] = "#da1337"


	// 通过声明映射创建一个 nil 映射
	//var myColors map[string]string
	// 将 Red 的代码加入到映射
	//myColors["Red"] = "#da1337"

	//Runtime Error:
//panic: runtime error: assignment to entry in nil map



	// 获取键Blue对应的值
	//value, exists := myColors["Blue"]
	// 这个键存在吗？
	//if exists {
	//	fmt.Println(value)
	//}


	//// 获取键 Blue 对应的值
	//value := myColors["Blue"]
	//// 这个键存在吗？
	//if value != "" {
	//	fmt.Println(value)
	//}

	// 创建一个映射，存储颜色以及颜色对应的十六进制代码
	myColors := map[string]string{
		"AliceBlue":"#f0f8ff",
		"Coral":"#ff7F50",
		"DarkGray":"#a9a9a9",
		"ForestGreen": "#228b22",
	}
	// 显示映射里的所有颜色
	for key, value := range myColors {
		fmt.Printf("Key: %s  Value: %s\n", key, value)
	}
    fmt.Println()

	// 调用函数来移除指定的键
	removeColor(myColors, "Coral")
	// 显示映射里的所有颜色
	for key, value := range myColors {
		fmt.Printf("Key: %s Value: %s\n", key, value)
	}

}

// removeColor 将指定映射里的键删除
func removeColor(colors map[string]string, key string) {
	delete(colors, key)
}
