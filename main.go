package main

import (
    "fmt"
    "log"
    "net/http"
    "html"
    "time"
    "math/rand"
)

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Request received", html.EscapeString(r.URL.Path), "   ", time.Now())

    rand.Seed(time.Now().UnixNano())
    var randomTime = rand.Intn(10)
    fmt.Println("Sleeping for ", randomTime, "s")
    time.Sleep(time.Duration(randomTime) * time.Second)

    fmt.Println("Responding now to request.")
    fmt.Fprintf(w, "Hi")
}

func handleRequests() {
    http.HandleFunc("/", homePage);

    fmt.Println("Server started")
    log.Fatal(http.ListenAndServe(":8081", nil))
}

func main() {
    handleRequests()
}
