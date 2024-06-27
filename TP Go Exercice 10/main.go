package main

import (
    "fmt"
    "math"
)


type Forme interface {
    aire() float64
}


type Cercle struct {
    rayon float64
}


func (c Cercle) aire() float64 {
    return math.Pi * c.rayon * c.rayon
}


type Rectangle struct {
    largeur, hauteur float64
}

func (r Rectangle) aire() float64 {
    return r.largeur * r.hauteur
}

func main() {

    c := Cercle{rayon: 8}
    r := Rectangle{largeur: 9, hauteur: 12}


    fmt.Printf("L'aire du cercle est: %.2f\n", c.aire())
    fmt.Printf("L'aire du rectangle est: %f=\n", r.aire())
}
