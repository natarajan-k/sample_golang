package main

import (
  "fmt"
  "net/http"
  //        "github.com/prometheus/client_golang/prometheus"
//        "github.com/prometheus/client_golang/prometheus/promauto"
//        "github.com/prometheus/client_golang/prometheus/promhttp"

)

func homePage(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "My Awesome Go App")
}

func setupRoutes() {
  http.HandleFunc("/", homePage)
}

func main() {
  fmt.Println("Go Web App Started on Port 3000")
  setupRoutes()
  http.ListenAndServe(":3000", nil)
}
