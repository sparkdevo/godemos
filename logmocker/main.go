package main

import (
	"fmt"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	//createlogfile("abc.log")
	//writelogs("abc.log")
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		//fmt.Println("abc" + strconv.Itoa(i) + ".log")
		filename := "abc" + strconv.Itoa(i) + ".log"
		createlogfile(filename)
		go writelogs(filename, &wg)
	}
	wg.Wait()
	fmt.Println("Over!")
}

// create log file
func createlogfile(filename string) {
	f, err := os.Create(filename)
	defer f.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
}

// open file, write 10w lines data
func writelogs(filename string, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_TRUNC, 0600)
	defer f.Close()
	count := 0
	for i := 0; i < 50000; i++ {
		count = count + 1
		time.Sleep(3 * time.Millisecond)
		_, err = f.Write([]byte("[Trace]	[0be1085c-ae63-463b-a6df-50ca5b19d415]	[2019/05/23]	[11:38:28 928]	[2]	[Wang WenTing]	[{0a5f0805-ab00-4449-bc78-4789eb38fbe2}]	[{4589e8a5-8786-4495-a85d-63253c2dd751}]	[AdminName]	[GrapeCity.Leyser.Common.Contract.ILeyserMessageService, Leyser.Contract]	[GetUnifyMessageData]	[LeyserMessageService]	[6.0025]	[]\n"))
		if err != nil {
			fmt.Println(err.Error())
		}
		if count == 1000 {
			count = 0
			f.Sync()
		}
	}
	//f.Write([]byte("over!"))
}