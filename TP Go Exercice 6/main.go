package main

import "fmt"

func main() {

    personnes := make(map[string]int)

    personnes["Alice"] = 30
    personnes["Bob"] = 25
    personnes["Charlie"] = 35

    for nom, age := range personnes {
        fmt.Printf("%s a %d ans\n", nom, age)
    }
}
