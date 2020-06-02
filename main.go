package main

import (
   "github.com/prometheus/client_golang/prometheus"
   "github.com/prometheus/client_golang/prometheus/promhttp"
   "net/http"
)

var ordersPlaced = prometheus.NewGauge(
   prometheus.GaugeOpts{
      Name: "orders_placed",
      Help: "Total number of orders placed",
   },
)

func main() {

   http.HandleFunc("/", Index)

   http.HandleFunc("/PlaceOrder", PlaceOrder)
   http.HandleFunc("/CancelOrder", CancelOrder)

   //Here we are saying that any call to /metrics should be handled
   // by prometheus's http handler for metrics
   http.Handle("/metrics", promhttp.Handler())

   //Here we are telling prometheus to keep track of ordersPlaced metric
   prometheus.MustRegister(ordersPlaced)

   // start server
   if err := http.ListenAndServe(":8080", nil); err != nil {
      panic(err)
   }

}

func Index(res http.ResponseWriter, req *http.Request) {

}

func PlaceOrder(res http.ResponseWriter, req *http.Request) {
   //Increment counter
   ordersPlaced.Inc()
}

func CancelOrder(res http.ResponseWriter, req *http.Request) {
   //Decrement counter
   ordersPlaced.Dec()
}
