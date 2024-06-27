package main

import "fmt"


func factorielle(n int) int {
    if n == 0 {
        return 1
    }
    return n * factorielle(n-1)
}

func main() {

   fmt.Println(factorielle(6))
   fmt.Println(factorielle(3))
   fmt.Println(factorielle(9))
    
}
