package main

import(
	"math/rand"
	"sync"
	"time"
	"fmt"
)

const(
	numberGoroutines = 2 // 要使用的goroutine的数量
	taskLoad = 5 // 要处理的工作的数量
)

// wg用来等待程序结束
var wg sync.WaitGroup

func init() {
	rand.Seed(time.Now().UnixNano())
}

// main是所有Go程序的入口
func main() {
    //demo1()
    demo2()
}

// 这个示例程序展示如何用无缓冲的通道来模拟
//2个goroutine间的网球比赛
func demo1() {
	// 创建一个无缓冲的通道
	court := make(chan int)

	// 计数加2，表示要等待两个goroutine
	wg.Add(2)

	// 启动两个选手
	go player("Nick", court)
	go player("Jack", court)

	// 发球
	court <- 1

	// 等待游戏结束
	wg.Wait()
}

// player 模拟一个选手在打网球
func player(name string, court chan int) {
	// 在函数退出时调用Done来通知main函数工作已经完成
	defer wg.Done()

	for{
		// 等待球被击打过来
		ball, ok := <-court
		if !ok {
			// 如果通道被关闭，我们就赢了
			fmt.Printf("Player %s Won\n", name)
			return
		}

		// 选随机数，然后用这个数来判断我们是否丢球
		n := rand.Intn(100)
		if n%5 == 0 {
			fmt.Printf("Player %s Missed\n", name)

			// 关闭通道，表示我们输了
			close(court)
			return
		}

		// 显示击球数，并将击球数加1
		fmt.Printf("Player %s Hit %d\n", name, ball)
		ball++

		// 将球打向对手
		court <- ball
	}
}

// 这个示例程序展示如何使用
// 有缓冲的通道和固定数目的
// goroutine来处理一堆工作
func demo2() {
	// 创建一个有缓冲的通道来管理工作
	tasks := make(chan string, taskLoad)

	// 启动goroutine来处理工作
	wg.Add(numberGoroutines)
	for gr := 1; gr <= numberGoroutines; gr++ {
		go worker(tasks, gr)
	}

	// 增加一组要完成的工作
	for post := 1; post <= taskLoad; post++ {
		tasks <- fmt.Sprintf("Task: %d", post)
	}

	// 当所有工作都处理完时关闭通道
	// 以便所有goroutine退出
	close(tasks)

	// 等待所有工作完成
	wg.Wait()
}

// worker作为goroutine启动来处理
// 从有缓冲的通道传入的工作
func worker(tasks chan string, worker int){
	// 通知函数已经返回
	defer wg.Done()

	for{
		// 等待分配工作
		task, ok := <-tasks
		if !ok{
			// 这意味着通道已经空了，并且已被关闭
			fmt.Printf("Worker: %d: Shutting Down\n", worker)
			return
		}

		// 显示我们开始工作了
		fmt.Printf("Worker: %d: Started %s\n", worker, task)

		// 随机等一段时间来模拟工作
		sleep := rand.Int63n(100)
		time.Sleep(time.Duration(sleep)* time.Millisecond)

		// 显示我们完成了工作
		fmt.Printf("Worker: %d: Completed %s\n", worker, task)
	}
}