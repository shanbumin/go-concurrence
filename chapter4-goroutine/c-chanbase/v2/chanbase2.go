package main

import (
	"fmt"
	"time"
)


//我重构了chanbase1.go,用单向通道约束了用于发送或接收的函数
var strChan = make(chan string, 3)

func main() {
	syncChan1 := make(chan struct{}, 1)
	syncChan2 := make(chan struct{}, 2)
	go receive(strChan, syncChan1, syncChan2) // 用于演示接收操作。
	go send(strChan, syncChan1, syncChan2)    // 用于演示发送操作。
	<-syncChan2
	<-syncChan2
}



//strChan  chan  string  ===>双向通道
//当<-在chan的右边，看起来像进入的样子，就是 ====>发送通道      strChan  chan<-  string
//当<-在chan的左边就是，看起来像出去的样子 ====>接收通道      strChan  <-chan  string

func receive(strChan <-chan string,syncChan1 <-chan struct{},syncChan2 chan<- struct{}) {

	<-syncChan1
	fmt.Println("Received a sync signal and wait a second... [receiver]")
	time.Sleep(time.Second)
	for {
		if elem, ok := <-strChan; ok {
			fmt.Println("Received:", elem, "[receiver]")
		} else {
			break
		}
	}
	fmt.Println("Stopped. [receiver]")
	syncChan2 <- struct{}{}
}

func send(strChan chan<- string, syncChan1 chan<- struct{},syncChan2 chan<- struct{}) {

	for _, elem := range []string{"a", "b", "c", "d"} {
		strChan <- elem
		fmt.Println("Sent:", elem, "[sender]")
		if elem == "c" {
			syncChan1 <- struct{}{}
			fmt.Println("Sent a sync signal. [sender]")
		}
	}
	fmt.Println("Wait 2 seconds... [sender]")
	time.Sleep(time.Second * 2)
	close(strChan)
	syncChan2 <- struct{}{}
}
