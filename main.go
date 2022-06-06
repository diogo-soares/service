package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		d, _ := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Ops"))
			return
		}

		fmt.Fprintf(w, "hello %s", d)
	})

	http.ListenAndServe(":9000", nil)
}
