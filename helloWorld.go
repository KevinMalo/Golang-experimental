package main

import (
	"fmt"
	"math"
)

// Declaracion de variables
var x, y, z int = 1, 2, 3
var c, python, java = true, false, "no!"

// Cosntantes
const Pi = 3.14

const (
	Big   = 1 << 100
	Small = Big >> 99
)

// Declaracion de funciones
func add(x, y int) int {
	return x + y
}

// Declaracion struct
type Vertex struct {
	X int
	Y int
}

func experimental() {

	// Declaracion de vaiables corta
	typescript, javascript, kevinscript := true, false, "yes!"

	fmt.Printf("Now you have %g problems.",
		math.Nextafter(1, 2))

	fmt.Println(add(42, 13))

	fmt.Println(x, y, z, c, python, java, typescript, javascript, kevinscript, Pi)

	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
		if i < 3 {
			fmt.Println("IFFFF")
		}
	}
	fmt.Println(sum)

	v := Vertex{1, 2}
	v.X = 4
	fmt.Println(v.X)

}
