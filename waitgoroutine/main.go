package main

import (
	"time"
	"fmt"
	"sync"
	"net/http"
)



func main() {
    // demo1()
    // demo2()
    // demo3()
    demo4()
}

func demo4() {
	var urls = []string{
		"https://www.baidu.com/",
		"https://www.cnblogs.com/",
	}

	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}

	wg.Wait()
}
func fetch(url string, wg *sync.WaitGroup) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	wg.Done()
	fmt.Println(resp.Status)
	return resp.Status, nil
}

func demo3() {
	var wg sync.WaitGroup
	wg.Add(2)
	say2("hello", &wg)
	say2("world", &wg)
	fmt.Println("over!")
}

func say2(s string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()

	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func demo2() {
	done := make(chan bool)
	go func() {
		for i := 0; i < 3; i++ {
			//time.Sleep(100 * time.Millisecond)
			fmt.Println("hello")
		}
		done <- true
	}()
	go func() {
		for i := 0; i < 3; i++ {
			//time.Sleep(100 * time.Millisecond)
			fmt.Println("world")
		}
		done <- true
	}()
	<-done
	<-done
	fmt.Println("over!")
}

func demo1() {
	go say("hello world")
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("over!")
}

func say(s string) {
	for i := 0; i < 3; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}