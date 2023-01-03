package main

import (
	"fmt"
)

type Forme interface {
	Air() float64
	Perimetre() float64
	Scaleup()
}

type Rectangle struct {
	largeur  float64
	longueur float64
}

/* Pas de méthode Air */

func (r Rectangle) Perimetre() float64 {
	return 2 * (r.largeur * r.longueur)
}

func (r *Rectangle) Scaleup() {
	r.largeur = 2 * r.largeur
}
func main() {
	f := Rectangle{5.0, 4.0}
	r := Rectangle{5.0, 4.0}
	fmt.Printf("Valeur de f : %v\n", f)
	fmt.Printf("Valeur de r : %v\n", r)
	r.Scaleup()
	fmt.Printf("Valeur de r après scaleup : %v\n", r)
}
