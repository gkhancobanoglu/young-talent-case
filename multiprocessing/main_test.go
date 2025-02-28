package main

import (
	"testing"
)

// Helper function to verify if the task execution order is valid
func isValidOrder(order []string, tasks map[string]Task) bool {
	position := make(map[string]int)
	for i, task := range order {
		position[task] = i
	}

	for _, task := range tasks {
		for _, dep := range task.depends {
			if position[dep] > position[task.name] {
				return false // Dependency violated
			}
		}
	}
	return true
}

// TestMinCompletionTime checks if minCompletionTime function works correctly
func TestMinCompletionTime(t *testing.T) {
	// Define a set of tasks with dependencies
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{}},
		"B": {name: "B", duration: 2, depends: []string{}},
		"C": {name: "C", duration: 4, depends: []string{}},
		"D": {name: "D", duration: 5, depends: []string{"A"}},
		"E": {name: "E", duration: 2, depends: []string{"B", "C"}},
		"F": {name: "F", duration: 3, depends: []string{"D", "E"}},
	}

	// Expected minimum completion time
	expectedTime := 11

	// Run the function
	minTime, order := minCompletionTime(tasks)

	// Check if the calculated minimum time is correct
	if minTime != expectedTime {
		t.Errorf("Expected completion time %d, but got %d", expectedTime, minTime)
	}

	// Validate the task execution order
	if !isValidOrder(order, tasks) {
		t.Errorf("Invalid task execution order: %v", order)
	}
}

// TestCycleDetection ensures no circular dependencies exist in the task set
func TestCycleDetection(t *testing.T) {
	// Define a set of tasks with a circular dependency
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{"C"}},
		"B": {name: "B", duration: 2, depends: []string{"A"}},
		"C": {name: "C", duration: 4, depends: []string{"B"}}, // Circular dependency
	}

	// Run the function
	minTime, order := minCompletionTime(tasks)

	// If a circular dependency exists, the order should be incomplete
	if len(order) == len(tasks) {
		t.Errorf("Cycle detection failed, expected incomplete order but got: %v", order)
	}

	// Check if function returned zero time for cycle cases
	if minTime > 0 {
		t.Errorf("Expected minTime to be 0 for cyclic dependencies, got %d", minTime)
	}
}

// TestIndependentTasks ensures that tasks with no dependencies execute in parallel
func TestIndependentTasks(t *testing.T) {
	// Define independent tasks with no dependencies
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{}},
		"B": {name: "B", duration: 2, depends: []string{}},
		"C": {name: "C", duration: 4, depends: []string{}},
	}

	// Expected minimum completion time (max duration)
	expectedTime := 4

	// Run the function
	minTime, _ := minCompletionTime(tasks)

	// Check if the calculated time is correct
	if minTime != expectedTime {
		t.Errorf("Expected completion time %d, but got %d", expectedTime, minTime)
	}
}
