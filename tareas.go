package main

import "fmt"

type taskList struct {
	tasks []*task
}

func (t *taskList) agregarALista(tl *task) {
	t.tasks = append(t.tasks, tl)
}

func (t *taskList) eliminarDeLista(index int) {
	t.tasks = append(t.tasks[:index], t.tasks[index+1:]...)
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

func ( t *taskList ) imprimirLista() {
	for _, tarea := range t.tasks{
		fmt.Println( "Nombre", tarea.nombre )
		fmt.Println( "Descriptcion", tarea.descripcion )
	}
}

func ( t *taskList ) imprimirListaCompletadas() {
	for _, tarea := range t.tasks{
		if tarea.completado == true {
			fmt.Println( "Nombre", tarea.nombre )
			fmt.Println( "Descriptcion", tarea.descripcion )
		}
	}
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

	lista.agregarALista(t4)
	lista.imprimirLista()
	lista.tasks[0].marcarCompetla()
	fmt.Println("Tareas completadas")
	lista.imprimirListaCompletadas()

	// for i := 0; i < len(lista.tasks); i++ {
	// 	fmt.Println( "Index", i, "Nombre", lista.tasks[i].nombre )
	// }

	// for index, tarea := range lista.tasks {
	// 	fmt.Println("Index", index, "Nombre", tarea.nombre)
	// }

	// Uso de break
	// for i := 0; i < 10; i++ {
		
	// 	if i == 5 {
	// 		break
	// 	}
	// 	fmt.Println(i)

	// }

	// Uso de continue
	// for i := 0; i < 10; i++ {
		
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
		
	// }

}
