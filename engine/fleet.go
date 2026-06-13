package engine

import (
	"fmt"
	"net/http"
	"sync"
)

func LaunchFleet(
	client *http.Client,
	publisher *Publisher,
	runID string,
	botCount int,
	ordersPerBot int,
) []OrderResult {

	var wg sync.WaitGroup
	var mu sync.Mutex

	results := make([]OrderResult, 0)

	for i := 1; i <= botCount; i++ {

		botID := fmt.Sprintf(
			"bot-%03d",
			i,
		)

		wg.Add(1)

		go func(botID string) {
			defer wg.Done()

			botResults := RunBot(
				client,
				publisher,
				botID,
				runID,
				ordersPerBot,
			)

			mu.Lock()
			results = append(results, botResults...)
			mu.Unlock()

		}(botID)
	}

	wg.Wait()

	return results
}