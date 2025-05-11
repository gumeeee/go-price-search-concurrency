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
		var totalPrice float64
		countPrices := 0.0
		for price := range priceChannel {
			totalPrice += price
			countPrices++
			avgPrice := totalPrice / countPrices
			fmt.Printf("Price received: R$ %.2f | Average Price at the moment: R$ %.2f\n", price, avgPrice)
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
