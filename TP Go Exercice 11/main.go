package main

import (
    "errors"
    "fmt"
)

func diviser(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division par zéro")
    }
    return a / b, nil
}

func main() {

    result, err := diviser(10, 2)
    if err != nil {
        fmt.Printf("Erreur: %v\n", err)
    } else {
        fmt.Printf("Résultat de 10 / 2 = %.2f\n", result)
    }

    result, err = diviser(10, 0)
    if err != nil {
        fmt.Printf("Erreur: %v\n", err)
    } else {
        fmt.Printf("Résultat de 10 / 0 = %.2f\n", result)
    }

    result, err = diviser(15, 3)
    if err != nil {
        fmt.Printf("Erreur: %v\n", err)
    } else {
        fmt.Printf("Résultat de 15 / 3 = %.2f\n", result)
    }
}
