package main

import (
	"fmt"
	"price-search/internal/fetcher"
	"sync"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan float64)
	var wg sync.WaitGroup
	wg.Add(3)

	go func() {
		for price := range priceChannel {
			fmt.Printf("Price: R$ %.2f\n", price)
		}
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromFirstSite()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromSecondSite()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetcher.FetchPriceFromThirdSite()
	}()

	wg.Wait()
	close(priceChannel)

	fmt.Printf("\nTempo Total: %s", time.Since(start))
}
