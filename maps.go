package main

import (
	"fmt"
)

// Map es una estrucgtura de 
// llave valor parecida a los diccionarios en python

func main() {

	m1 := make(map[string]int)

	m1["a"] = 8

	fmt.Println(m1)
	fmt.Println(m1.a)

	
}