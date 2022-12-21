package finalProject

import (
	"fmt"
	"time"
)

// Job represents a job to be executed, with a name and a number and a delay
type Job struct {
	Name   string
	Delay  time.Duration
	Number int
}

// Worker will be our concurrency-friendly worker
type Worker struct {
	Id         int
	JobQueue   chan Job
	WorkerPool chan chan Job
	QuitChan   chan bool
}

// Dispatcher is a dispatcher that will dispatch jobs to workers
type Dispatcher struct {
	WorkerPool chan chan Job
	MaxWorkers int
	JobQueue   chan Job
}

func NewWorker(id int, workerPool chan chan Job) *Worker {
	return &Worker{
		Id:         id,
		WorkerPool: workerPool,
		JobQueue:   make(chan Job),
		QuitChan:   make(chan bool),
	}
}

func (w *Worker) Start() {
	go func() {
		for {
			w.WorkerPool <- w.JobQueue // add job to pool

			// Multiplexing
			select {
			case job := <-w.JobQueue: // get job from queue
				fmt.Printf("worker%d: started %s, %d\n", w.Id, job.Name, job.Number)
				fib := Fibonacci(job.Number)
				time.Sleep(job.Delay)
				fmt.Printf("worker%d: finished %s, %d with result %d\n", w.Id, job.Name, job.Number, fib)
			case <-w.QuitChan: // quit if worker is told to do so
				fmt.Printf("Worker with id %d Stopped\n", w.Id)
				return
			}
		}
	}()
}

// Stop method stop the worker
func (w Worker) Stop() {
	go func() {
		w.QuitChan <- true
	}()
}

// Fibonacci calculates the fibonacci sequence
func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

// NewDispatcher returns a new Dispatcher with the provided maxWorkers
func NewDispatcher(jobQueue chan Job, maxWorkers int) *Dispatcher {
	pool := make(chan chan Job, maxWorkers)
	return &Dispatcher{
		WorkerPool: pool,
		MaxWorkers: maxWorkers,
		JobQueue:   jobQueue,
	}
}

// Dispatch will dispatch jobs to workers
func (d *Dispatcher) dispatch() {
	for {
		select {
		case job := <-d.JobQueue: // get job from queue
			// Asign the job to a worker
			go func() {
				jobChannel := <-d.WorkerPool // get worker from pool
				jobChannel <- job            // Workers will read from this channel
			}()
		}
	}
}

func (d *Dispatcher) Run() {
	for i := 0; i < d.MaxWorkers; i++ {
		worker := NewWorker(i+1, d.WorkerPool)
		worker.Start()
	}

	go d.dispatch()
}
