package main

import "fmt"

type taskList struct {
	tasks []*task
}

func ( t *taskList ) agregarALista( tl *task ) {
	t.tasks = append( t.tasks, tl )
}

func ( t *taskList ) eliminarDeLista( index int) {
	t.tasks = append( t.tasks[:index], t.tasks[index+1:]... )
}

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

	t1 := &task{
		nombre:      "Completar mi curso de GO",
		descripcion: "Completar mi curso de GO en platzi en esta semana",
	}

	t2 := &task{
		nombre:      "Completar mi curso de Python",
		descripcion: "Completar mi curso de Python en platzi en esta semana",
	}

	t3 := &task{
		nombre:      "Completar mi curso de C#",
		descripcion: "Completar mi curso de C# en platzi en esta semana",
	}

	t4 := &task{
		nombre:      "Completar mi curso de Node",
		descripcion: "Completar mi curso de Node en platzi en esta semana",
	}

	lista := &taskList{
		tasks: []*task{
			t1, t2, t3,
		},
	}

	fmt.Println( lista.tasks[0] )
	lista.agregarALista(t4)
	fmt.Println( len(lista.tasks) )
	lista.eliminarDeLista(1)
	fmt.Println( len(lista.tasks) )

}
