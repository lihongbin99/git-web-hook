package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data, err := ioutil.ReadAll(r.Body)
		if err != nil {
			return
		}

		fmt.Print("\n\n", string(data), "\n\n")
	})

	http.ListenAndServe(":13520", nil)
}
