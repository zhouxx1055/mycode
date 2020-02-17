package main

import (
	"fmt"
	"log"
	"time"
)

func service1(ch chan string) {
	log.Println("BEEFORE")
	ch <- "from service1"
	log.Println("AFTER")
}

func main() {
	var ch chan string

	go service1(nil) // 写
	log.Printf("sleep at :%d\n", time.Now().Unix())
	time.Sleep(time.Second * 2)
	log.Printf("weak at :%d\n", time.Now().Unix())
	for {
		select {
		case str := <-ch: // 读
			fmt.Println(str)
		}
	}
}
