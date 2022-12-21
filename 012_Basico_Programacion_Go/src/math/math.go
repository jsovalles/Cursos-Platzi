package mathclass

import "fmt"

func MathExamples() {
	//Declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.14
	fmt.Println("First pi", pi)
	fmt.Println("Second pi", pi2)

	//Declaracion variables enteras
	base := 12
	var altura int = 14
	var area int
	fmt.Println(base, altura, area)

	//Zero values
	var a int
	var b float32
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseCuadrado = 10
	areaCuadrado := baseCuadrado * baseCuadrado
	fmt.Println("Area al cuadrado", areaCuadrado)

	//Operadores aritmeticos
	x := 10
	y := 50
	result := x + y
	fmt.Println("suma", result)
	result = y - x
	fmt.Println("resta", result)

	// Rectángulo
	baseRectangulo := 20
	alturaRectangulo := 10

	areaRectangulo := baseRectangulo * alturaRectangulo

	fmt.Println("El Area del Rectángulo es :", areaRectangulo)

	// Circulo : AreaCirculo = pi por radio al cudrado
	const PI float64 = 3.14 // Constant
	var radioCirculo float64 = 10

	areaCirculo := PI * radioCirculo * radioCirculo

	fmt.Println("El Area del Circulo es :", areaCirculo)

	// Trapecio
	var baseUno float64 = 6
	var baseDos float64 = 15
	var alturaTrapecio float64 = 25

	areaTrapecio := ((baseUno + baseDos) * alturaTrapecio) / 2

	fmt.Println("El Area del Trapecio es :", areaTrapecio)

	//Datos primitivos

	//Numeros enteros
	//int = Depende del OS (32 o 64 bits)
	//int8 = 8bits = -128 a 127
	//int16 = 16bits = -2^15 a 2^15-1
	//int32 = 32bits = -2^31 a 2^31-1
	//int64 = 64bits = -2^63 a 2^63-1

	//Optimizar memoria cuando sabemos que el dato simpre va ser positivo
	//uint = Depende del OS (32 o 64 bits)
	//uint8 = 8bits = 0 a 127
	//uint16 = 16bits = 0 a 2^15-1
	//uint32 = 32bits = 0 a 2^31-1
	//uint64 = 64bits = 0 a 2^63-1

	//numeros decimales
	// float32 = 32 bits = +/- 1.18e^-38 +/- -3.4e^38
	// float64 = 64 bits = +/- 2.23e^-308 +/- -1.8e^308

	//textos y booleanos
	//string = ""
	//bool = true or false

	//numeros complejos
	//Complex64 = Real e Imaginario float32
	//Complex128 = Real e Imaginario float64
	//Ejemplo : c:=10 + 8i

	return
}
