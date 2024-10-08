package workerpool

import (
	"fmt"
	"strconv"
	"sync"
)

type Job struct {
	ID      int
	Payload string
}

type Result struct {
	JobID   int
	Outcome string
}

func worker(id int, jobs <-chan Job, results chan<- Result, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range jobs {
		// Process the job
		result := Result{JobID: job.ID, Outcome: "Processed by worker " + strconv.Itoa(id)}
		results <- result
	}
}

func main() {
	const numWorkers = 5
	jobs := make(chan Job, 100)
	results := make(chan Result, 100)
	var wg sync.WaitGroup

	// Start workers
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, jobs, results, &wg)
	}

	// Send jobs
	for j := 1; j <= 20; j++ {
		jobs <- Job{ID: j, Payload: "Data"}
	}
	close(jobs)

	// Wait for workers to finish
	wg.Wait()
	close(results)

	// Collect results
	for result := range results {
		fmt.Println("Job", result.JobID, ":", result.Outcome)
	}
}
