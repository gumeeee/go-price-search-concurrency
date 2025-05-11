package main

import (
	"fmt"
	"price-search/internal/fetcher"
	"price-search/internal/processor"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	var consumerWg sync.WaitGroup

	consumerWg.Add(1)

	go func() {
		defer consumerWg.Done()
		processor.ShowPricesAndAVG(priceChannel)
	}()

	go fetcher.FetchPrices(priceChannel)

	consumerWg.Wait()

	fmt.Println("Consumer finished.")

	fmt.Printf("\nTempo Total: %s", time.Since(start))
}
