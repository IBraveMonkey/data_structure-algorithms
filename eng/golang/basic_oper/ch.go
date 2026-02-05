package main

import (
	"fmt"
)

// Error: Deadlock
func writeToNilChannel() {
	var ch chan int
	ch <- 1
}

// Error: Deadlock
func redToNilChannel() {
	var ch chan int
	<-ch
}

// Error: panic on write to closed channel
func writeToClosedChannel() {
	ch := make(chan int, 2)
	close(ch)
	ch <- 20
}

// Error: deadlock when ranging over nil channel
func rangeNilChannel() {
	var ch chan int
	for range ch {
	}
}

// Error: panic on closing nil channel
func closeNilChannel() {
	var ch chan int
	close(ch)
}

// Error: panic on closing closed channel
func closeChannelAnyTimes() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// Works: compares channels
func compareChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	equal1 := ch1 == ch2
	equal2 := ch1 == ch1

	fmt.Println(equal1)
	fmt.Println(equal2)
}

// Works: reads from channel, closes it, and reads from closed channel
func readFromChannel() {
	ch := make(chan int, 2)
	ch <- 10
	ch <- 20

	val, ok := <-ch
	fmt.Println(val, ok)

	close(ch)
	val, ok = <-ch
	fmt.Println(val, ok)

	val, ok = <-ch
	fmt.Println(val, ok)
}

// Works: reads from one of the channels
func readAnyChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go func() {
		ch1 <- 100
	}()

	go func() {
		ch2 <- 200
	}()

	select {
	case val1 := <-ch1:
		fmt.Println(val1)
	case val2 := <-ch2:
		fmt.Println(val2)
	}
}
