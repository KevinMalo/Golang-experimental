package main

import "fmt"

type task struct {
	nombre      string
	descripcion string
	completado  bool
}

func (t *task) marcarCompetla() {
	t.completado = true
}

func (t *task) actualizarDescripcion(descripcion string) {
	t.descripcion = descripcion
}

func (t *task) actualizarNombre(nombre string) {
	t.nombre = nombre
}

func main() {

	t := &task{
		nombre:      "Completar mi curso de GO",
		descripcion: "Completar mi curso de GO en platzi en esta semana",
	}

	fmt.Printf("%+v\n", t)
	t.marcarCompetla()
	t.actualizarNombre("Finalizar mi curso de GO")
	t.actualizarDescripcion("Completar mi curso cuanto antes")
	fmt.Printf("%+v\n", t)

}
