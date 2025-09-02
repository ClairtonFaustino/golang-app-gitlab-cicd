package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "ok - time=%s\n", time.Now().Format(time.RFC3339))
	})

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("ok"))
	})

	fmt.Println("Servidor rodando na porta 8080 🚀")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}
