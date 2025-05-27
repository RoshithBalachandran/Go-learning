package main

import (
	"fmt"
	"sync"
)

var count=0
var mutex sync.Mutex

func increment(wg *sync.WaitGroup){
	mutex.Lock()
	count++
	mutex.Unlock()
	wg.Done()
}


func main(){
	var wg sync.WaitGroup
		for i:=0;i<50;i++{
			wg.Add(1)
			go increment(&wg)
		}
		wg.Wait()
		fmt.Println("Final Result is ",count)
}