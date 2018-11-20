package common

import (
	"fmt"
	"net/http"
)

func StartServer(port int, fn func() string) error {
	fmt.Println("Listening on port", port)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, fn())
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "OK")
	})

	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
