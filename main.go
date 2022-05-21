package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		d, err := ioutil.ReadAll(r.Body)

		if err == nil {
			http.Error(w, "Oops", http.StatusBadRequest)
			return
		}

		fmt.Fprint(w, "hello %s", d)
	})

	http.ListenAndServe(":9090", nil)
}
