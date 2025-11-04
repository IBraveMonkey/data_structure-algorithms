package fan_out

import (
	"fmt"
	"sync"
	"time"
)

func worker(
	id int,
	input <-chan int,
	wg *sync.WaitGroup,
) {
	defer wg.Done()
	for num := range input {
		fmt.Printf("Worker %d processing %d\n", id, num)
		time.Sleep(500 * time.Millisecond)
		fmt.Printf("Worker %d finished %d, result: %d\n", id, num, num*2)
	}
}

func FanOut() {
	const numWorkers = 3
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8}

	inputs := make([]chan int, numWorkers)
	var wg sync.WaitGroup

	for i := 0; i < numWorkers; i++ {
		inputs[i] = make(chan int)
		wg.Add(1)
		go worker(i+1, inputs[i], &wg)
	}

	go func() {
		for i, num := range numbers {
			inputs[i%numWorkers] <- num
		}

		for _, in := range inputs {
			close(in)
		}
	}()

	wg.Wait()
	fmt.Println("All workers done")
}

func JoinChannels(chs ...<-chan int) <-chan int {
	mergedCh := make(chan int)

	go func() {
		wg := &sync.WaitGroup{}

		wg.Add(len(chs))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)

			wg.Wait()
			close(mergedCh)
		}
	}()

	return mergedCh
}
