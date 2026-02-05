package scheduler

import (
	"errors"
	"sync"
	"time"
)

/*
 * The Scheduler pattern is used to schedule the execution of functions in the future.
 * It allows setting a timer for function execution and cancelling execution if necessary.
 * Usage example: performing scheduled tasks, such as sending notifications or clearing cache.
 */

// Scheduler manages delayed execution of functions
type Scheduler struct {
	mutex   sync.Mutex          // Protection against data races when accessing actions and closed
	closed  bool                // Scheduler state flag (closed or not)
	actions map[int]*time.Timer // Map of active timers by keys
}

// NewScheduler creates a new scheduler instance
func NewScheduler() *Scheduler {
	return &Scheduler{
		actions: make(map[int]*time.Timer), // Initialize timer map
	}
}

// SetTimeout starts the function after the specified delay
func (s *Scheduler) SetTimeout(key int, delay time.Duration, action func()) error {
	// Check delay correctness
	if delay < 0 {
		return errors.New("invalid delay")
	}

	// Check that function is passed
	if action == nil {
		return errors.New("invalid action")
	}

	s.mutex.Lock()         // Lock for safe access
	defer s.mutex.Unlock() // Unlock on exit

	// Check if scheduler is closed
	if s.closed {
		return errors.New("scheduler is closed")
	}

	// If timer with such key already exists, stop it
	if timer, found := s.actions[key]; found {
		timer.Stop() // Stop old timer
	}

	// Create new timer that will execute action after delay
	s.actions[key] = time.AfterFunc(delay, action)
	return nil
}

// CancelTimeout cancels scheduled function execution
func (s *Scheduler) CancelTimeout(key int) {
	s.mutex.Lock()         // Lock for safe access
	defer s.mutex.Unlock() // Unlock on exit

	// Search timer by key
	if timer, found := s.actions[key]; found {
		timer.Stop()           // Stop timer
		delete(s.actions, key) // Remove from map
	}
}

// Close stops all scheduled functions and closes the scheduler
func (s *Scheduler) Close() {
	s.mutex.Lock()         // Lock for safe access
	defer s.mutex.Unlock() // Unlock on exit

	s.closed = true // Mark scheduler as closed

	// Stop and remove all active timers
	for key, timer := range s.actions {
		timer.Stop()           // Stop timer
		delete(s.actions, key) // Remove from map
	}
}
