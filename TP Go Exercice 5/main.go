package main

import (
	"fmt"
	"sync"
)

func envoyer(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 1; i <= 5; i++ {
		ch <- i
	}
	close(ch)
}

func recevoir(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range ch {
		fmt.Println(num)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	wg.Add(2)

	go envoyer(ch, &wg)

	go recevoir(ch, &wg)

	wg.Wait()
	fmt.Println("Toutes les goroutines sont terminées.")
}
