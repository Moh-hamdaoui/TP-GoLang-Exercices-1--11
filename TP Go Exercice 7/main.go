package main

import "fmt"

func main() {
  
    for i := 1; i <= 10; i++ {
        fmt.Println(i)
    }

    for i := 1; i <= 10; i++ {
        if i%2 == 0 {
            fmt.Printf("%d est pair\n", i)
        } else {
            fmt.Printf("%d est impair\n", i)
        }
    }


    var dayOfWeek int = 3
    switch dayOfWeek {
    case 1:
        fmt.Println("L'hiver")
    case 2:
        fmt.Println("Le printemps")
    case 3:
        fmt.Println("L'été")
    case 4:
        fmt.Println("L'automne")
    default:
        fmt.Println("Season invalide")
    }
}
