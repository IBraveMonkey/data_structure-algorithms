package main

import (
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

/*
 * The Errgroup pattern is used to manage a group of goroutines that can return errors.
 * It allows running multiple goroutines and waiting for their completion, returning the first error if one occurred.
 * Usage example: executing multiple HTTP requests concurrently and returning the result or error from the slowest request.
 */

// User represents a user to process
type User struct {
	Name string // User name
}

// fetch simulates fetching user data (e.g., from an API)
func fetch(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{}) // Channel for completion signal

	// Goroutine to simulate asynchronous work
	go func() {
		time.Sleep(10 * time.Millisecond) // Simulate request latency
		close(ch)                         // Signal completion
	}()

	// Wait for work completion or context cancellation
	select {
	case <-ch:
		return user.Name, nil // Successful completion
	case <-ctx.Done():
		return "", ctx.Err() // Context cancelled or timed out
	}
}

// process processes a list of users concurrently using errgroup
func process(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0) // Map for counting users
	mu := sync.Mutex{}                 // Mutex to protect map from data race

	// Create errgroup with context (automatic cancellation on error)
	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100) // Limit the number of concurrent goroutines

	// Start a goroutine for each user
	for _, u := range users {
		egroup.Go(func() error {
			// Get user data with errgroup context
			name, err := fetch(ectx, u)
			if err != nil {
				return err // Return error (will stop other goroutines via context)
			}

			// Safely update map with lock
			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1 // Increment counter for name
			return nil
		})
	}

	// Wait for all goroutines to complete or the first error
	if err := egroup.Wait(); err != nil {
		return nil, err // Return error if any
	}

	return names, nil // Return result
}

func main() {
	// Test data
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background() // Create base context

	start := time.Now()
	// Process users concurrently
	res, err := process(ctx, names)
	if err != nil {
		fmt.Println("an error occurred:", err.Error())
	}

	fmt.Print("time: ", time.Since(start)) // Execution time
	fmt.Println(res)                       // Result: map with count of each name
}

// ===================================================================================
// Advanced implementation with goroutine limit
// ===================================================================================
// This implementation uses errgroup.SetLimit to limit the number of concurrently running goroutines.
// This is useful when you need to avoid system overload or exceeding connection limits.

// fetchAdvanced simulates fetching data with explicit timeout handling
func fetchAdvanced(ctx context.Context, user User) (string, error) {
	ch := make(chan struct{}) // Channel for completion signal

	// Goroutine to simulate asynchronous work
	go func() {
		time.Sleep(10 * time.Millisecond) // Simulate request latency
		close(ch)                         // Signal completion
	}()

	// Wait for work completion or context cancellation
	select {
	case <-ch:
		return user.Name, nil // Successful completion
	case <-ctx.Done():
		return "", errors.New("timed out") // Explicit timeout error
	}
}

// processAdvanced shows advanced usage of errgroup with SetLimit
func processAdvanced(ctx context.Context, users []User) (map[string]int64, error) {
	names := make(map[string]int64, 0) // Map for counting users
	mu := sync.Mutex{}                 // Mutex to protect map from data race

	// Create errgroup with context
	egroup, ectx := errgroup.WithContext(ctx)
	egroup.SetLimit(100) // Limit the number of concurrently running goroutines

	// IMPORTANT: SetLimit means that maximum 100 goroutines will work simultaneously.
	// If more tasks are launched, the rest will wait for their turn.

	// Start a goroutine for each user
	for _, u := range users {
		egroup.Go(func() error {
			// Get user data with errgroup context
			name, err := fetchAdvanced(ectx, u)
			if err != nil {
				return err // Return error (will stop other goroutines)
			}

			// Safely update map with lock
			mu.Lock()
			defer mu.Unlock()
			names[name] = names[name] + 1 // Increment counter for name

			return nil
		})
	}

	// Wait for all goroutines to complete or the first error
	if err := egroup.Wait(); err != nil {
		return nil, err // Return error if any
	}

	return names, nil // Return result
}

// mainAdvanced demonstrates usage of the advanced version
func mainAdvanced() {
	// Test data
	names := []User{
		{"Ann"},
		{"Bob"},
		{"Cindy"},
		{"Bob"},
	}

	ctx := context.Background() // Create base context

	start := time.Now()
	// Process users concurrently with limit of 100 goroutines
	res, err := processAdvanced(ctx, names)
	if err != nil {
		fmt.Println("an error occurred: ", err.Error())
	}

	fmt.Println("time: ", time.Since(start)) // Execution time
	fmt.Println(res)                         // Result: map with count of each name
}
