package channel

import (
	"fmt"
	"sync"
)

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
	fmt.Println(*numb)
	*numb += 1
	ch <- "OKEven"
}

func Odd(ch chan string, numb *int) {
	fmt.Println(*numb)
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
			go Even(chanOK, &num)
		} else {
			go Odd(chanOK, &num)
		}
		ok := <-chanOK
		fmt.Println(ok)
	}

}

func support(ch chan string, wt *sync.WaitGroup) {
	fmt.Println("Entered support using go routine")
	wt.Done()
	ch <- "Done"
}
func ChannelWorkerMain() {
	var wt sync.WaitGroup
	wt.Add(1)
	ch := make(chan string)
	go support(ch, &wt)
	wt.Wait()
	cp := <-ch
	fmt.Println(cp)
}
