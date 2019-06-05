package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	for _, url := range os.Args[1:] {
		// Add https:// as prefix if missing
		const prefix = "http://"
		if !strings.HasPrefix(url, prefix) {
			url = prefix + url
		}
		resp, err := http.Get(url)
		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: %v\n", err)
			os.Exit(1)
		}

		fmt.Fprintf(os.Stdout, "fetch http status code: %s\n", resp.Status)

		_, err = io.Copy(os.Stdout, resp.Body)

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: copying %s: %v\n", url, err)
			os.Exit(1)
		}

		err = resp.Body.Close()

		if err != nil {
			fmt.Fprintf(os.Stderr, "fetch: closing %s: %v\n", url, err)
			os.Exit(1)
		}
	}
}
