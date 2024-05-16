package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

const URL = "https://jsonplaceholder.typicode.com/"

type todo struct {
	ID        int    `json:id`
	Title     string `json:title`
	Completed bool   `json:completed`
}

func main() {
	res, err := http.Get(URL + "/todos/1")

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(-1)
	}

	defer res.Body.Close()

	if res.StatusCode == http.StatusOK {
		body, err := io.ReadAll(res.Body)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		var item todo

		err = json.Unmarshal(body, &item)

		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}

		fmt.Printf("%#v", item)
	}
}
