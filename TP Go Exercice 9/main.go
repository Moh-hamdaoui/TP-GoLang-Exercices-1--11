package main

import "fmt"

func compteur() func() int {
    count := 0 

    return func() int {
        count++
        return count
    }
}

func main() {

    cpt := compteur()

    fmt.Println(cpt()) 
    fmt.Println(cpt())
    fmt.Println(cpt())
    fmt.Println(cpt())
    fmt.Println(cpt())
}
