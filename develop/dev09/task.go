package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	url := flag.String("url", "", "URL of the website to download")
	output := flag.String("output", "downloaded_site.html", "Output file name")

	flag.Parse()

	if *url == "" {
		fmt.Println("URL is required. Use -url flag.")
		return
	}

	downloadWebsite(*url, *output)

	fmt.Println("Website downloaded successfully.")
}

func downloadWebsite(url string, filename string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Error while downloading the website:", err)
		return
	}
	defer resp.Body.Close()

	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Error copying response to file:", err)
		return
	}
}
