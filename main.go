package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func rate(w http.ResponseWriter, r *http.Request) {
	const rate_api = "http://api.promasters.net.br/cotacao/v1/valores"
	base := r.URL.Query().Get("base")
	log.Println(r.Method, r.URL)

	if base == "" {
		base = "USD"
	}

	res, err := http.Get(rate_api + "?alt=json&moedas=" + base)

	if err != nil {
		log.Fatal(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	res.Body.Close()

	if err != nil {
		log.Fatal(err)
	}

	w.Write(body)
}

func main() {
	port := os.Getenv("PORT")
	http.HandleFunc("/rates", rate)
	http.ListenAndServe(":"+port, nil)
}