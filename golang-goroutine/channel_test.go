package golang_goroutine

import (
	"fmt"
	"strconv"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {
	// cara membuat channel
	channel := make(chan string)
	// pastikan channel selalu ditutup
	defer close(channel)

	// running as a goroutine
	go func (){
		time.Sleep(2 * time.Second)

		// kirim data ke channel
		channel <- "Dzikri Nur Akbar"
		fmt.Println("Done")
	}()

	// terima/ambil data dari channel
	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func TestChannelAsParams(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go GiveMeResponse(channel)

	data := <- channel
	fmt.Println(data)

	time.Sleep(5 * time.Second)
}

func GiveMeResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Udindindun"
}

// channel hanya untuk mengirim data (memasukkan data ke channel)
func OnlyIn(channel chan<- string) {
	time.Sleep(1 * time.Second)
	channel <- "Data dari only in channel"
}

// channel hanya untuk mengambil/menerima data dari channel
func OnlyOut(channel <-chan string) {
	data := <- channel
	fmt.Println(data)
}

func TestInOutChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go OnlyIn(channel)
	go OnlyOut(channel)

	time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t * testing.T) {
	// menambahkan buffered channel
	channel := make(chan string, 3)
	defer close(channel)

	channel <- "Kiki"
	channel <- "Kinyun"
	channel <- "Udindindun"

	fmt.Println(<- channel)
	fmt.Println(<- channel)
	fmt.Println(<- channel)

	time.Sleep(2 * time.Second)
	fmt.Println("Done")
}

func TestRangeChannel(t *testing.T) {
	channel := make(chan string)

	go func() {
		for i := 0; i< 10; i++ {
			channel <- "Perulangan ke " + strconv.Itoa(i)
		}
		close(channel)
	} ()

	for data := range channel {
		fmt.Println("Menerima data :", data)
	}

	fmt.Println("Selesai")
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
			case data := <- channel1:
				fmt.Println("Data dari channel 1: ", data)
				counter++
			case data := <- channel2:
				fmt.Println("Data dari channel 2: ", data)
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
			case data := <- channel1:
				fmt.Println("Data dari channel 1: ", data)
				counter++
			case data := <- channel2:
				fmt.Println("Data dari channel 2: ", data)
				counter++
			default :
				fmt.Println("Menunggu data")
		}

		if counter == 2 {
			break
		}
	}
}