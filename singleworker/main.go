package main

import (
	"fmt"
)

// Task struct represents a job task with name, duration, and dependencies
type Task struct {
	name     string
	duration int
	depends  []string
}

// Function to compute the minimum completion time and execution order for a single worker
func singleWorkerSchedule(tasks map[string]Task) (int, []string) {
	// Track the number of dependencies for each task
	inDegree := make(map[string]int)
	// Graph to store task dependencies
	graph := make(map[string][]string)

	for name, task := range tasks {
		if _, exists := inDegree[name]; !exists {
			inDegree[name] = 0
		}
		for _, dep := range task.depends {
			graph[dep] = append(graph[dep], name)
			inDegree[name]++
		}
	}

	// Queue to store available tasks that can be processed
	queue := []string{}
	for name, count := range inDegree {
		if count == 0 {
			queue = append(queue, name)
		}
	}

	timeElapsed := 0
	order := []string{}

	// Process tasks in sequential order (one at a time)
	for len(queue) > 0 {
		// Find the task with the shortest duration
		minIndex := 0
		for i := 1; i < len(queue); i++ {
			if tasks[queue[i]].duration < tasks[queue[minIndex]].duration {
				minIndex = i
			}
		}

		// Pick the shortest task and remove it from the queue
		current := queue[minIndex]
		queue = append(queue[:minIndex], queue[minIndex+1:]...)
		order = append(order, current)
		timeElapsed += tasks[current].duration

		// Unlock dependent tasks
		for _, next := range graph[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
		}
	}

	return timeElapsed, order
}

func main() {
	// Define tasks
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{}},
		"B": {name: "B", duration: 2, depends: []string{}},
		"C": {name: "C", duration: 4, depends: []string{}},
		"D": {name: "D", duration: 5, depends: []string{"A"}},
		"E": {name: "E", duration: 2, depends: []string{"B", "C"}},
		"F": {name: "F", duration: 3, depends: []string{"D", "E"}},
	}

	// Compute schedule for single worker
	minTime, order := singleWorkerSchedule(tasks)
	fmt.Printf("Minimum completion time with a single worker: %d units of time\n", minTime)
	fmt.Printf("Task execution order: %v\n", order)
}
