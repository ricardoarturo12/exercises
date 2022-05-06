package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	Status int
}

func crawl(wId int, jobs <-chan Site, results chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", wId)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Println(err.Error())
		}
		results <- Result{Status: resp.StatusCode}
	}
}

func main() {
	fmt.Println("worker pools in go")

	// dos canales para enviar los trabajos y obtener resultados
	jobs := make(chan Site, 3)
	results := make(chan Result, 3)

	for w := 1; w <= 3; w++ {
		go crawl(w, jobs, results)
	}

	// qty de urls
	urls := [4]string{
		"http://www.google.com",
		"http://www.facebok.com",
		"http://www.kilka.com",
		"http://www.kilka.com",
	}

	// envía los trabajos a realizar
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	// cantidad de urls ->
	// imprime los resultados
	for a := 1; a <= 4; a++ {
		result := <-results
		log.Println(result.Status)
	}
}
