package main

import (
	"fmt"
)

// Represents a task with dependencies
// Each task has a name, duration, and a list of dependent tasks

type Task struct {
	name     string
	duration int
	depends  []string
}

// Function to calculate the minimum completion time for all tasks and their execution order
func minCompletionTime(tasks map[string]Task) (int, []string) {
	// Maps to store task dependencies and execution details
	inDegree := make(map[string]int)       // Number of dependencies for each task
	graph := make(map[string][]string)     // Adjacency list representing task dependencies
	timeToComplete := make(map[string]int) // Earliest completion time for each task

	// Build the dependency graph
	for name, task := range tasks {
		if _, exists := inDegree[name]; !exists {
			inDegree[name] = 0
		}
		for _, dep := range task.depends {
			graph[dep] = append(graph[dep], name)
			inDegree[name]++
		}
	}

	// Initialize queue with independent tasks (tasks with no dependencies)
	queue := []string{}
	for name, count := range inDegree {
		if count == 0 {
			queue = append(queue, name)
			timeToComplete[name] = tasks[name].duration
		}
	}

	order := []string{}
	totalTime := 0

	// Perform topological sorting using Kahn's Algorithm
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		order = append(order, current)

		// Update dependent tasks
		for _, next := range graph[current] {
			inDegree[next]--
			if inDegree[next] == 0 {
				queue = append(queue, next)
			}
			// Update completion time for dependent tasks
			if timeToComplete[next] < timeToComplete[current]+tasks[next].duration {
				timeToComplete[next] = timeToComplete[current] + tasks[next].duration
			}
		}
		// Track the maximum completion time
		if timeToComplete[current] > totalTime {
			totalTime = timeToComplete[current]
		}
	}

	return totalTime, order
}

func main() {
	// Define tasks with their dependencies
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{}},
		"B": {name: "B", duration: 2, depends: []string{}},
		"C": {name: "C", duration: 4, depends: []string{}},
		"D": {name: "D", duration: 5, depends: []string{"A"}},
		"E": {name: "E", duration: 2, depends: []string{"B", "C"}},
		"F": {name: "F", duration: 3, depends: []string{"D", "E"}},
	}

	// Compute the minimum completion time and task order
	minTime, order := minCompletionTime(tasks)
	fmt.Printf("Minimum completion time: %d units of time\n", minTime)
	fmt.Printf("Task execution order: %v\n", order)
}
