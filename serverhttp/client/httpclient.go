package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func main() {
	res, err := http.Get("http://localhost:8080/" + os.Args[1])

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

		fmt.Println(string(body))
	}
}
