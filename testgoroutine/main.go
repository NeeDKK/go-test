package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var WaitGroup sync.WaitGroup
	WaitGroup.Add(1)
	fmt.Println("zzzzzzzzzz")
	go testPrint(&WaitGroup)
	fmt.Println("zzzzzzzzzzz")
	WaitGroup.Wait()
}

func testPrint(WaitGroup *sync.WaitGroup) {
	defer WaitGroup.Done()
	time.Sleep(5 * time.Second)
	fmt.Println("AAAAAA")
}
