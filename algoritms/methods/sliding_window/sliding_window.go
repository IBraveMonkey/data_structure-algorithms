package main

/*
Шаблон для sliding window
func findMaxAverage(nums []int, k int) float64 {
    begin := 0
    windowState
    result

    for end := 0; end < len(nums);end++ {
  			windowState = end-begin + 1 // window size
        if //window condition
        		result
          	windowState
            begin += 1 // shrink window
    }

    return result;
}

*/

func main() {
	nums := []int{}
	k := 4
	begin := 0
	window_state := 0
	result := 0.0

	for end := 0; end < len(nums); end++ {
		window_state += nums[end]

		if end-begin+1 == k {
			result = max(result, float64(window_state))

			window_state -= nums[begin]
			begin++
		}
	}

	// return result / float64(k)
}
