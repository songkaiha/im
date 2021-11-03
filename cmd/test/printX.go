package main

import (
	"fmt"
	"sync"
)

func PrintX(wg *sync.WaitGroup, ch chan string) {
	for i := 0; i < 100; i++ {

		//time.Sleep(1 * time.Second)
		switch <-ch {
		case "A":
			fmt.Println("A")
			ch <- "B"
		case "B":
			fmt.Println("B")
			ch <- "C"
		case "C":
			fmt.Println("C")
			ch <- "D"
		case "D":
			fmt.Println("D")
			if i != 99 {
				ch <- "A"
			}

		}

		wg.Done()
	}
}

func main() {
	var wg = new(sync.WaitGroup)
	var ch = make(chan string)
	for i := 1; i <= 4; i++ {
		wg.Add(100)
		go PrintX(wg, ch)
	}
	ch <- "A"
	wg.Wait()
	close(ch)
}
