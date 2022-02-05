package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	var decision string

	fmt.Println("Please input for calling API from `db` y or n :")

	fmt.Scan(&decision)

	url := "http://localhost:8080/wordcounter"

	if decision == "db" {
		url = "http://localhost:8080/wordcounterdb"
	}

	method := "POST"

	payload := strings.NewReader(FileReader())

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "text/plain")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func FileReader() string {
	fileName := "GoLang_Test.txt"
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}

	return string(data)
}
