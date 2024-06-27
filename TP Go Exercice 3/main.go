package main

import "fmt"

func main() {
 
    x := 10
    fmt.Println("La valeur de x avant:", x)

    ptr := &x

    *ptr = 20

    fmt.Println("La nouvelle valeur de x après modification via le pointeur:", x)
}
