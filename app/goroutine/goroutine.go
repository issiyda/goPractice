package main

import (
	"fmt"
	"time"
)

func oneSecondWait() {
	time.Sleep(1000 * time.Millisecond)
	fmt.Println("1 second passed")
}

func twoSecondWait() {
	time.Sleep(2000 * time.Millisecond)
	fmt.Println("2 second passed")
}

func threeSecondWait() {
	time.Sleep(3000 * time.Millisecond)
	fmt.Println("1 second passed")
}

func main() {
	fmt.Println(time.Now())
	go threeSecondWait()
	go twoSecondWait()
	oneSecondWait()
	fmt.Println(time.Now())
}
