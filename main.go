package main

import (
    "os"
    "encoding/json"
    "log"
    "net/http"
    "time"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    tasksProcessed = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "myapp_tasks_processed_total",
            Help: "Total number of tasks processed.",
        },
        []string{"type"},
    )
)

var (
    httpRequestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "http_requests_total",
            Help: "Total number of HTTP requests received.",
        },
        []string{"method", "path"},
    )
)

func init() {
    prometheus.MustRegister(tasksProcessed)
    prometheus.MustRegister(httpRequestsTotal)
}

func hostnameHandler(w http.ResponseWriter, r *http.Request) {
    httpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).Inc()
    hostname, err := os.Hostname()
    if err != nil {
        http.Error(w, "Failed to get hostname", http.StatusInternalServerError)
        return
    }
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Hostname: " + hostname))
}

func taskHandler(w http.ResponseWriter, r *http.Request) {
    var task struct {
        Type     string `json:"type"`
        Duration int    `json:"duration"`
    }

    if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    go func() {
    	switch task.Type {
    	case "cpu":
            simulateCPUTask(task.Duration)
            tasksProcessed.WithLabelValues("cpu").Inc()
            httpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).Inc()
        case "memory":
            simulateMemoryTask(task.Duration)
            tasksProcessed.WithLabelValues("memory").Inc()
            httpRequestsTotal.With(prometheus.Labels{"method": r.Method, "path": r.URL.Path}).Inc()
        default:
            http.Error(w, "Invalid task type", http.StatusBadRequest)
            return
        }

	log.Println("Task processed")
    }()

    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Task processed"))
}

func simulateCPUTask(duration int) {
    log.Println("Starting CPU-intensive task")
    endTime := time.Now().Add(time.Duration(duration) * time.Second)
    for time.Now().Before(endTime) {
        for i := 0; i < 1000000; i++ {
        }
    }
    log.Println("Completed CPU-intensive task")
}

func simulateMemoryTask(duration int) {
    log.Println("Starting memory-intensive task")
    numElements := 10000000
    data := make([]int, numElements)
    for i := range data {
        data[i] = i
    }
    time.Sleep(time.Duration(duration) * time.Second)
    log.Println("Completed memory-intensive task")
}

func main() {
    http.HandleFunc("/", hostnameHandler)
    http.Handle("/metrics", promhttp.Handler())
    http.HandleFunc("/task", taskHandler)

    port := ":8080"
    log.Printf("Starting server on %s", port)
    if err := http.ListenAndServe(port, nil); err != nil {
        log.Fatalf("Failed to start server: %v", err)
    }
}
