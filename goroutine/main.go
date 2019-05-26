package main

import (
	"runtime"
	"sync"
	"fmt"
)

func main() {
    //demo1()
    //demo2()
    demo3()
}

func demo1() {
	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// wg用来等待程序完成
	// 计数加2，表示要等待两个goroutine
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Start Goroutines")

	// 声明一个匿名函数，并创建一个goroutine
	go func(){
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count< 3; count++{
			for char := 'a'; char< 'a'+26; char++{
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()
	// 声明一个匿名函数，并创建一个goroutine
	go func(){
		// 在函数退出时调用Done来通知main函数工作已经完成
		defer wg.Done()

		// 显示字母表3次
		for count := 0; count< 3; count++{
			for char := 'A'; char< 'A'+26; char++{
				fmt.Printf("%c ", char)
			}
			fmt.Println()
		}
	}()

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

func demo2() {
	// wg用来等待程序完成
	var wg sync.WaitGroup

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(1)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A", &wg)
	go printPrime("B", &wg)

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}

// printPrime 显示5000以内的素数值
func printPrime(prefix string, wg *sync.WaitGroup){
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

next:
	for outer := 2; outer < 5000; outer++ {
		for inner := 2; inner < outer; inner++ {
			if outer % inner == 0 {
				continue next
			}
		}
		fmt.Printf("%s:%d\n", prefix, outer)
	}
	fmt.Println("Completed", prefix)
}

func demo3() {
	// wg用来等待程序完成
	var wg sync.WaitGroup

	// 分配一个逻辑处理器给调度器使用
	runtime.GOMAXPROCS(2)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 创建两个goroutine
	fmt.Println("Create Goroutines")
	go printPrime("A", &wg)
	go printPrime("B", &wg)

	// 等待goroutine结束
	fmt.Println("Waiting To Finish")
	wg.Wait()

	fmt.Println("Terminating Program")
}