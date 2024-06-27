package main

import "fmt"

func somme(nombres []int) int {
    total := 0
    for _, nombre := range nombres {
        total += nombre
    }
    return total
}

func main() {

    nombres := []int{1, 2, 3, 4, 5}

    fmt.Println("Taille du tableau:", len(nombres))

    fmt.Println("Capacité du tableau:", cap(nombres))

    nombres = append(nombres, 6)
    fmt.Println("Tableau après ajout de 6:", nombres)

    nombres = append(nombres, 7, 8, 9)
    fmt.Println("Tableau après ajout de 7, 8, 9:", nombres)

    resultat := somme(nombres)
    fmt.Println("Somme des éléments du tableau:", resultat)
}
