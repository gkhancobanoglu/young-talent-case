package main

import (
	"testing"
)

// Helper function to check if two slices are equal
func slicesEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Test function for single worker scheduling
func TestSingleWorkerSchedule(t *testing.T) {
	tasks := map[string]Task{
		"A": {name: "A", duration: 3, depends: []string{}},
		"B": {name: "B", duration: 2, depends: []string{}},
		"C": {name: "C", duration: 4, depends: []string{}},
		"D": {name: "D", duration: 5, depends: []string{"A"}},
		"E": {name: "E", duration: 2, depends: []string{"B", "C"}},
		"F": {name: "F", duration: 3, depends: []string{"D", "E"}},
	}

	expectedTime := 19
	expectedOrder := []string{"B", "A", "C", "E", "D", "F"}

	minTime, order := singleWorkerSchedule(tasks)

	// Validate completion time
	if minTime != expectedTime {
		t.Errorf("Expected minimum completion time %d, got %d", expectedTime, minTime)
	}

	// Validate execution order
	if !slicesEqual(order, expectedOrder) {
		t.Errorf("Expected execution order %v, got %v", expectedOrder, order)
	}
}
