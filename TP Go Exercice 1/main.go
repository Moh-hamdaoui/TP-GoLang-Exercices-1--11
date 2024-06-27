package main

import "fmt"

var age int = 25

func double(n int) int {
    return n * 2
}

func main() {

    result := double(age)
    

    fmt.Println("Double de l'âge:", result)
}
