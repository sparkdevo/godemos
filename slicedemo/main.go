package main

import "fmt"

func main() {
	//myNum := make([]int, 5)
	//
	//for index, value := range myNum {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}

	// 创建一个整型切片
	// 使其长度大于容量
	//myNum := make([]int, 5, 3)
	//
	//for index, value := range myNum {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}

	//myNum := []int{10, 20, 30, 40, 50}
	//newNum := myNum[1:3]
	//newNum[3] = 45


	//myNum := []int{10, 20, 30, 40, 50}
	//// 创建一个新切片
	//// 其长度为 2 个元素，容量为 4 个元素
	//newNum := myNum[1:3]
	//// 使用原有的容量来分配一个新元素
	//// 将新元素赋值为 60
	//newNum = append(newNum, 60)
	//
	//for index, value := range myNum {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}
	//
	//for index, value := range newNum {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}

	//fruit := []string{"Apple", "Orange", "Plum", "Banana", "Grape"}
	//myFruit := fruit[2:3:6]
	//for index, value := range myFruit {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}


	//// 创建两个切片，并分别用两个整数进行初始化
	//num1 := []int{1, 2}
	//num2 := []int{3, 4}
	//// 将两个切片追加在一起，并显示结果
	//fmt.Printf("%v\n", append(num1, num2...))
	//

	//myNum := []int{10, 20, 30, 40, 50}
	//
	//for index, _ := range myNum {
	//	myNum[index] += 1
	//}
	//
	//for index, value := range myNum {
	//	fmt.Printf("index: %d value: %d\n", index, value)
	//}


	// 创建一个整型切片
	// 其长度和容量都是4个元素
	//myNum := []int{10, 20, 30, 40, 50}
	//// 从第三个元素开始迭代每个元素
	//for index := 2; index < len(myNum); index++ {
	//	fmt.Printf("Index: %d Value: %d\n", index, myNum[index])
	//}


	num1 := []int{10, 20, 30}
	num2 := make([]int, 5)
	count := copy(num2, num1)
	fmt.Println(count)
	fmt.Println(num2)

}
