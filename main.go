package main

import "fmt"
import "net/http"

func main() {
handlerIndex := func(w http.ResponseWriter, r *http.Request) {
fmt. Fprintln(w, "Hallo nama saya Nabila")
}
http.HandleFunc("/", handlerIndex)

fmt.Println("server started at 43.173.13.18:9025")
http.ListenAndServe(":9025", nil)

}
