package channel

import "fmt"

func worker(ch chan int) {
	ch <- 42
}

func Channelmain() {
	ch := make(chan int)
	go worker(ch)
	value := <-ch
	fmt.Println("Recieved: ", value)
}

func Even(ch chan string, numb *int) {
	fmt.Println(numb)
	*numb += 1
	ch <- "OKEven"
}

func Odd(ch chan string, numb *int) {
	fmt.Println(numb)
	*numb += 1
	ch <- "OKOdd"
}

func ChannelPrint() {
	chanOK := make(chan string)
	num := 1
	for {
		if num > 100 {
			break
		}
		if num%2 == 0 {
			Even(chanOK, &num)
		} else {
			Odd(chanOK, &num)
		}
		ok := <-chanOK
		fmt.Println(ok)
	}

}
