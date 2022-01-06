package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func HelloWorld(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello World - %s", time.Now().String())
	log.Println(r.URL)
}


func main() {
	http.HandleFunc("/", HelloWorld)

	fmt.Println("Servidor rodando na porta :8080")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalln(err)
	}
}