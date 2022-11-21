package main

import "net/http"

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>MeetUP - <strong>ArgoCD</strong></h1>"))
	})

	http.ListenAndServe(":8080", nil)
}