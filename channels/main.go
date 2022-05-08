package main

import "fmt"

func Example1() {
	ch1 := make(chan int, 1) // задал размер буфера для канала

	go func(in chan int) {
		val := <-in
		fmt.Println("GO get from chan1", val)
		fmt.Println("GO after read from chan")
	}(ch1)

	ch1 <- 42
	ch1 <- 42

	fmt.Println("MAIN after put to chan")
	fmt.Scanln()
}

func Example2() {
	in := make(chan int)

	go func(out chan<- int) {
		for i := 0; i < 4; i++ {
			fmt.Println("before", i)
			out <- i
			fmt.Println("after", i)
		}
		close(out)
		fmt.Println("generator finish")
	}(in)

	for i := range in {
		fmt.Println("\tget", i)
	}
}

func Example3() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int)
	ch1 <- 1
	select {
	case val := <-ch1:
		fmt.Println("ch1 val", val)
	case ch2 <- 1:
		fmt.Println("put val to ch2")
	default:
		fmt.Println("default case")
	}
}

func Example4() {
	ch1 := make(chan int, 2)
	ch1 <- 1
	ch1 <- 2
	ch2 := make(chan int, 2)
	ch2 <- 3
LOOP:
	for {
		select {
		case v1 := <-ch1:
			fmt.Println("ch1 val", v1)
		case v2 := <-ch2:
			fmt.Println("ch2 val", v2)
		default:
			break LOOP
		}
	}
}

func Example5() {
	cancelCh := make(chan struct{}) // канал для отмены
	dataCh := make(chan int)        // канал для данных

	go func(cancel chan struct{}, data chan int) {
		val := 0
		for {
			select {
			case <-cancelCh: // только если что-то прочитали из канала отмены
				return
			case dataCh <- val: // иначе будем писать в канал с данными
				val++
			}
		}
	}(cancelCh, dataCh)

	for curVal := range dataCh {
		fmt.Println("read", curVal)
		if curVal > 3 {
			fmt.Println("send cancel")
			cancelCh <- struct{}{}
			break
		}
	}
}

func main() {
	Example5()
}
