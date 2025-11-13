package main

import "net/http"

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("Hello World from Serg Fedorov!"))
    })
    http.ListenAndServe(":8083", nil)
}
