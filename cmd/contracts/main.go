package main

import (
	"log"

	"dogovorsBackEnd/internal"
)

func main() {
	log.Fatalln("Ошибка запуска веб-сервера ", internal.App())
}
