package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * Паттерн Worker Pool (Пул воркеров) используется для распределения задач между несколькими обработчиками.
 * Он позволяет эффективно использовать ресурсы, распределяя задачи между несколькими горутинами.
 * Пример использования: обработка изображений, где каждое изображение обрабатывается одним из воркеров.
 */
type Task struct {
	ID       int
	Filename string
}

func Process(task Task) string {
	time.Sleep(1 * time.Second)

	return fmt.Sprintf("FileID: %d done - %s \n", task.ID, task.Filename)
}

func Worker(taskCh <-chan Task, resCh chan<- string) {
	for task := range taskCh {
		resCh <- Process(task)
	}
}

func main() {
	const (
		numWorkers = 3
		numTasks   = 10
	)

	taskCh := make(chan Task, numTasks)
	resCh := make(chan string, numTasks)

	wg := sync.WaitGroup{}
	wg.Add(numWorkers)

	for range numWorkers {
		go func() {
			defer wg.Done()
			Worker(taskCh, resCh)
		}()
	}

	go func() {
		for i := range numTasks {
			taskCh <- Task{ID: i, Filename: fmt.Sprintf("file_%d.jpg", i)}
		}
		defer close(taskCh)
	}()

	go func() {
		wg.Wait()
		close(resCh)
	}()

	for res := range resCh {
		fmt.Println(res)
	}
}
