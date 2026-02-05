package main

import (
	"fmt"
)

// Ошибка: Блокировка навсегда
func writeToNilChannel() {
	var ch chan int
	ch <- 1
}

// Ошибка: Блокировка навсегда
func redToNilChannel() {
	var ch chan int
	<-ch
}

// Ошибка: panic при записи в закрытый канал
func writeToClosedChannel() {
	ch := make(chan int, 2)
	close(ch)
	ch <- 20
}

// Ошибка: блокировка при итерации по nil каналу
func rangeNilChannel() {
	var ch chan int
	for range ch {
	}
}

// Ошибка: panic при закрытии nil канала
func closeNilChannel() {
	var ch chan int
	close(ch)
}

// Ошибка: panic при повторном закрытии канала
func closeChannelAnyTimes() {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// Работает: сравнивает каналы
func compareChannels() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	equal1 := ch1 == ch2
	equal2 := ch1 == ch1

	fmt.Println(equal1)
	fmt.Println(equal2)
}

// Работает: читаем из канала, закрывает его и читаем из закрытого канала
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

// Работает: читаем из одного из каналов
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
