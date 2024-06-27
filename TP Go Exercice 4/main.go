package main

import (
	"fmt"
	"sync"
)

func afficher() {
	for i := 1; i <= 10; i++ {
		fmt.Println(i)
	}
}

func main() {

	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		afficher()
	}()

	wg.Wait()

}
