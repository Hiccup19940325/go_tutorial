package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	totalJobs := 5

	for i := 1; i <= totalJobs; i++ {
		processJobs(i)
	}
	fmt.Println("Total time ", time.Since(startTime))
}

func processJobs(i int) {
	fmt.Println("Started  job", i)
	time.Sleep(time.Second)
	fmt.Println("Finished job", i)
}
