package main

import (
    "fmt"
    "os"
    "strconv"
    "log"
    "net/http"
    "encoding/json"
    "github.com/google/uuid"
)

type response struct {
    Hostname   string `json:"hostname"`
    Uuid string `json:"uuid"`
}

func handler(w http.ResponseWriter, r *http.Request) {
    uuid_random, _ := uuid.NewRandom()
    hostname, _ := os.Hostname()

    data := &response{
        Hostname:   hostname,
        Uuid:       fmt.Sprintf("%s", uuid_random)}
    response, _ := json.Marshal(data)

    fmt.Fprintf(w, string(response))
}

func main() {
    port := 80

    fmt.Printf("Listening on :%s\n", strconv.Itoa(port))

    http.HandleFunc("/", handler)

    log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", port), nil))
}
