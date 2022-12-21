package interfaces

import "fmt"

type figuras2d interface {
	area() float64
}

type cuadrado struct {
	base float64
}

type rectangulo struct {
	base   float64
	altura float64
}

func (c cuadrado) area() float64 {
	return c.base * c.base
}

func (r rectangulo) area() float64 {
	return r.base * r.altura
}

func calcular(f figuras2d) {
	fmt.Println("Area: ", f.area())
}

func FuncionesExample() {
	myCuadrado := cuadrado{base: 4}
	myRectangulo := rectangulo{base: 4, altura: 6}

	calcular(myCuadrado)
	calcular(myRectangulo)

	// Lista interfaces
	myInterface := []interface{}{"Hola", 12, 4.90}
	fmt.Println(myInterface...)

}
