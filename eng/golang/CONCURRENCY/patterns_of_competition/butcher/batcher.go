package batcher

import (
	"errors"
	"sync"
	"time"
)

/*
 * The Batcher pattern is used to combine multiple operations into a single group (batch) to reduce the number of calls or improve performance.
 * It accumulates messages until a certain size or time interval is reached, and then performs an action on all accumulated messages.
 * Usage example: sending logs to a server, where logs are collected into batches and sent every few seconds or when a certain number of logs are accumulated.
 */

// Batcher manages the accumulation and group processing of messages
type Batcher struct {
	maxSize       int            // Maximum batch size (how many messages to accumulate)
	flushInterval time.Duration  // Time interval for automatic sending
	flushAction   func([]string) // Function that processes a batch of messages
	ticker        *time.Ticker   // Timer for periodic sending

	mutex    sync.Mutex // Protection against data races when accessing messages
	messages []string   // Accumulated messages (current batch)

	closeCh     chan struct{} // Channel for close signal
	closeDoneCh chan struct{} // Channel for work completion confirmation
}

// NewBatcher creates a new batcher instance with parameter validation
func NewBatcher(action func([]string), size int, interval time.Duration) (*Batcher, error) {
	// Check that a processing function is passed
	if action == nil {
		return nil, errors.New("invalid action")
	}

	// Check batch size correctness
	if size <= 0 {
		return nil, errors.New("invalid size")
	}

	// Check interval correctness
	if interval <= 0 {
		return nil, errors.New("invalid interval")
	}

	// Create and return new batcher
	return &Batcher{
		maxSize:       size,
		flushAction:   action,
		flushInterval: interval,
		closeCh:       make(chan struct{}), // Initialize close channel
		closeDoneCh:   make(chan struct{}), // Initialize confirmation channel
	}, nil
}

// Append adds a message to the batch
func (b *Batcher) Append(message string) error {
	b.mutex.Lock()         // Lock for safe access to messages
	defer b.mutex.Unlock() // Unlock on exit

	// Check if batcher is closed (non-blocking check)
	select {
	case <-b.closeCh:
		return errors.New("batcher is closed")
	default:
	}

	// Add message to batch
	b.messages = append(b.messages, message)

	// If batch reached max size, send immediately
	if len(b.messages) == b.maxSize {
		b.flushLocked()                 // Send batch (mutex already acquired)
		b.ticker.Reset(b.flushInterval) // Reset timer for next batch
	}

	return nil
}

// Run starts the worker for periodic batch sending
func (b *Batcher) Run() {
	// Prevent multiple starts
	if b.ticker != nil {
		return
	}

	// Create timer for periodic sending
	b.ticker = time.NewTicker(b.flushInterval)

	// Start worker goroutine
	go func() {
		defer close(b.closeDoneCh) // Signal completion on exit

		for {
			// First select: check closure without blocking
			select {
			case <-b.closeCh:
				b.flush() // Send remaining messages
				return
			default:
			}

			// Second select: wait for events (timer or closure)
			select {
			case <-b.closeCh:
				b.flush() // Send remaining messages
				return
			case <-b.ticker.C:
				b.flush() // Send batch by timer
			}
		}
	}()
}

// flush sends accumulated messages (with locking)
func (b *Batcher) flush() {
	b.mutex.Lock()         // Lock for safe access
	defer b.mutex.Unlock() // Unlock on exit

	b.flushLocked() // Call locked version
}

// flushLocked sends accumulated messages (without locking, mutex already acquired)
func (b *Batcher) flushLocked() {
	// If batch is empty, do nothing
	if len(b.messages) == 0 {
		return
	}

	messages := b.messages // Save current batch
	b.messages = nil       // Clear batch for new messages

	// Run processing in a separate goroutine (do not block Append)
	go b.flushAction(messages)
}

// Close waits for worker completion and sends remaining messages
func (b *Batcher) Close() {
	// Check if batcher is already closed (non-blocking check)
	select {
	case <-b.closeCh:
		return // Already closed, exit
	default:
	}

	// Lock and close signal channel
	b.mutex.Lock()
	close(b.closeCh) // Signal worker to finish
	b.mutex.Unlock()

	<-b.closeDoneCh // Wait for worker completion
	b.ticker.Stop() // Stop timer
}
