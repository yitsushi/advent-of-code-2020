package day08

import "fmt"

// UnknownOperation occurs when the source code contains an unknown operation.
type UnknownOperation struct {
	Operation string
	Value     int64
}

func (e UnknownOperation) Error() string {
	return fmt.Sprintf("unknown operation: %s %d", e.Operation, e.Value)
}

// ExecutionError occurs when the machine meets an unexpected error.
type ExecutionError struct{}

func (e ExecutionError) Error() string {
	return "execution error"
}

// InfiniteLoopDetected occurs when the machine detects an infinite loop.
type InfiniteLoopDetected struct{}

func (e InfiniteLoopDetected) Error() string {
	return "infinite loop detected"
}
