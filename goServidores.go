package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	inicio := time.Now()

	canal := make(chan string)

	servidores := []string{
		"http://platzi.com",
		"http://google.com",
		"http://facebook.com",
		"http://instagram.com",
	}

	i := 0

	for {

		if i > 10 {
			break
		}

		for _, servidor := range servidores {
			go revisarServidor(servidor, canal)
		}

		time.Sleep(1 * time.Second)
		fmt.Println(<-canal)

		i++
	}

	tiempoTranscurrido := time.Since(inicio)
	fmt.Printf("Tiempo de ejecucion %s\n", tiempoTranscurrido)

}

func revisarServidor(servidor string, canal chan string) {
	_, err := http.Get(servidor)
	if err != nil {
		canal <- servidor + " no esta disponible :("
	} else {
		canal <- servidor + " servidor disponible :)"
	}
}
