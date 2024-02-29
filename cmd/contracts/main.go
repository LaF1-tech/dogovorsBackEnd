package main

import (
	"dogovorsBackEnd/internal"
	"log"
)

func main() {
	log.Fatalln("Ошибка запуска веб-сервера ", internal.App())
}
