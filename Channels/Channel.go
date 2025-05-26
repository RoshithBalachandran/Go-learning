package main

import (
	"fmt"
	"time"
)

func Producer(Ch chan int){
	for i:=0;i<=5;i++{
		fmt.Println("Producer : ",i)
		Ch<- i
		time.Sleep(time.Second*1)
	}
}
func Consumer(Ch chan int){
	for item :=range Ch{
		fmt.Println("Consumer : ",item)
	}
}


func main(){
	Ch:=make(chan int)
	go Producer(Ch)
	go Consumer(Ch)

	time.Sleep(7*time.Second)
}