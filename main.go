package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func ProxyHandler(hrp *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println(r.URL)
		rw.Header().Set("X-COOKIE", "Olá Mundo")
		hrp.ServeHTTP(rw, r)
	}

	
}


func main() {
	remote, err := url.Parse("http://localhost:8080")

	if err != nil {
		panic(err)
	}


	proxy := httputil.NewSingleHostReverseProxy(remote)

	http.HandleFunc("/", ProxyHandler(proxy))

	fmt.Println("Iniciando proxy na porta 80")

	if err =  http.ListenAndServe(":80", nil); err != nil {
		log.Fatalln("Erro ao iniciar o proxy")
	}
}
