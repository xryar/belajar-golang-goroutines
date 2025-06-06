package main

import (
	"fmt"
	"strconv"
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

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}

		close(channel)
	}()

	for data := range channel {
		fmt.Println("Menerima data", data)
	}

	fmt.Println("selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++
		}

		if counter == 2 {
			break
		}
	}
}

func TestDefaultSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)

	defer close(channel1)
	defer close(channel2)

	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)

	counter := 0
	for {
		select {
		case data := <-channel1:
			fmt.Println("Menerima data dari channel 1", data)
			counter++
		case data := <-channel2:
			fmt.Println("Menerima data dari channel 2", data)
			counter++
		default:
			fmt.Println("Menunggu data...")
			time.Sleep(2 * time.Second)
		}

		if counter == 2 {
			break
		}
	}
}
