package main
import (
        "net/http"
        "fmt"
        "math/rand"
//        "github.com/prometheus/client_golang/prometheus"
        "github.com/prometheus/client_golang/prometheus/promauto"
        "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    goRandomValue = promauto.NewGauge(prometheus.GaugeOpts{
	Name: "go_random_value",
	Help: "randomly generated value in Go",
    })
)

func main() {
    http.HandleFunc("/", handle)
    http.Handle("/metrics", promhttp.Handler())
}

func handle(w http.ResponseWriter, r *http.Request) {
    goRandomValue.Set(rand.Float64())
    fmt.Fprintf(w, "Hello1")
}
