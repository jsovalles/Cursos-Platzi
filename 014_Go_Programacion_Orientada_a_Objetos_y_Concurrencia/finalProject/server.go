package finalProject

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"time"
)

const (
	maxWorkers   = 4
	maxQueueSize = 20
	port         = ":8080"
)

func Project() {
	jobQueue := make(chan Job, maxQueueSize)
	dispatcher := NewDispatcher(jobQueue, maxWorkers)
	dispatcher.Run()

	router := mux.NewRouter()
	router.HandleFunc("/fib", func(w http.ResponseWriter, r *http.Request) {
		RequestHandler(w, r, jobQueue)
	}).Methods(http.MethodPost)
	log.Println("Starting server, listening on port %s", port)

	// Start the server, and log any errors
	log.Fatal(http.ListenAndServe(port, router))
}
func RequestHandler(w http.ResponseWriter, r *http.Request, jobQueue chan Job) {
	if r.Method != "POST" {
		w.Header().Add("Allow", "POST")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Parse the request
	delay, err := time.ParseDuration(r.FormValue("delay"))
	if err != nil {
		http.Error(w, "Invalid delay", http.StatusBadRequest)
		return
	}

	value, err := strconv.Atoi(r.FormValue("value"))
	if err != nil {
		http.Error(w, "Invalid value", http.StatusBadRequest)
		return
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Invalid name", http.StatusBadRequest)
		return
	}

	// Create the job
	job := Job{
		Name:   name,
		Delay:  delay,
		Number: value,
	}

	// Add the job to the queue
	jobQueue <- job // send the job
	w.WriteHeader(http.StatusOK)
}
