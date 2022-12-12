package main

import "fmt"

type Rectangle struct {
	largeur float64
	longeur float64
}

func (p *Rectangle) init(longeur, largeur float64) {
	p.largeur = largeur
	p.longeur = longeur
}
func (p *Rectangle) scale(factor float64) {
	p.largeur *= factor
	p.longeur *= factor

}

func main() {
	var samir = Rectangle{}
	samir.init(1, 2)
	fmt.Println(samir)
	samir.scale(5)
	fmt.Println(samir)
}
