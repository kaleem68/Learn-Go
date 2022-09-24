package main

import (
	"fmt"
)

var println = fmt.Println

func nums1(mychannel chan int) {
	mychannel <- 1
	mychannel <- 2
}
func nums2(mychannel2 chan int) {
	mychannel2 <- 4
	mychannel2 <- 3
}

// concurrency / go routine
func main() {
	channel1 := make(chan int)
	go nums1(channel1)

	channel2 := make(chan int)
	go nums2(channel2)

	println(<-channel1)
	println(<-channel2)
	println(<-channel1)
	println(<-channel2)

}
