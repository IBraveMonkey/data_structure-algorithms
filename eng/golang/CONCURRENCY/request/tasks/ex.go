package tasks

import (
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"
)

/*
	There is a third-party weather service (its simulation is the WeatherForecast function)
	The third-party service takes a second to respond, which is long for us.
	Our service receives a high load. How to improve the current implementation?
	1. Propose and implement a solution.
	2. Additional task: The third-party service can provide data for more than one city.
	Refine the implementation from step 1 taking this fact into account.
*/

type Data struct {
	Temperatures map[string]int
	mu           sync.RWMutex
}

func NewData(interval time.Duration) *Data {
	ticker := time.NewTicker(interval)

	newData := &Data{}

	go func() {
		ticker.Stop()

		for range ticker.C {
			newData.UpdateTemperature()
		}
	}()

	return &Data{}
}

func (d *Data) UpdateTemperature() {
	wg := sync.WaitGroup{}

	for city := range d.Temperatures {
		wg.Add(1)
		go func() {
			defer wg.Done()
			tmp := WeatherForecast(city)

			d.mu.Lock()
			d.Temperatures[city] = tmp
			d.mu.Unlock()
		}()
	}

	wg.Wait()
}

func (d *Data) GetTemperature(city string) (int, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	t, ok := d.Temperatures[city]
	if !ok {
		return 0, fmt.Errorf("city - %s not found", city)
	}
	return t, nil
}

func WeatherForecast(city string) int {
	time.Sleep(1 * time.Second)
	return rand.Intn(70) - 30
}

func task_koz() {
	data := NewData(1 * time.Second)

	http.HandleFunc("/weather", func(w http.ResponseWriter, r *http.Request) {
		temp, err := data.GetTemperature("Kazan")
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		fmt.Fprintf(w, "{\"temperature\":%d\n", temp)
	})

	if err := http.ListenAndServe(":3333", nil); err != nil {
		panic(err)
	}
}
