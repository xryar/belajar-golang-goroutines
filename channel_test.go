package main

import (
	"fmt"
	"testing"
	"time"
)

/*
	channel adalah media komunikasi antar goroutine
	channel adalah tipe data yang bisa menampung data dari goroutine lain
*/

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		time.Sleep(2 * time.Second)
		channel <- "Arya Rizki Andaru"
		fmt.Println("Berhasil mengirim data ke channel")
	}()

	data := <-channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Arya Rizki Andaru"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <-channel
	fmt.Println(data)
}

// channel sebagai parameter input
func OnlyIn(channel chan<- string) {
	time.Sleep(2 * time.Second)
	channel <- "Arya Rizki Andaru"
}

// channel sebagai parameter output
func OnlyOut(channel <-chan string) {
	data := <-channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	channel := make(chan string, 3)
	defer close(channel)

	go func() {
		channel <- "Arya"
		channel <- "Rizki"
		channel <- "Andaru"
	}()

	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()

	fmt.Println(cap(channel))
	fmt.Println(len(channel))

	fmt.Println("selesai")
}
