package main

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {

	cont, cancel := context.WithTimeout(context.Background(), time.Millisecond*30)

	defer cancel()

	url := "https://placehold.co/600x400"
	// Create Request
	req, err := http.NewRequestWithContext(cont, http.MethodGet, url, nil)

	if err != nil {
		panic(err)
	}

	// perform http request
	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Downloaded image of size %d\n", len(body))

}
