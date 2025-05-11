package fetcher

import (
	"math/rand"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- float64) {
	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		defer wg.Done()
		priceChannel <- fetchPriceFromFirstSite()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetchPriceFromSecondSite()
	}()

	go func() {
		defer wg.Done()
		priceChannel <- fetchPriceFromThirdSite()
	}()

	go func() {
		defer wg.Done()
		fetchAndSendMultiplePrices(priceChannel)
	}()

	wg.Wait()
	close(priceChannel)
}

func fetchPriceFromFirstSite() float64 {
	time.Sleep(1 * time.Second)

	return rand.Float64() * 100
}

func fetchPriceFromSecondSite() float64 {
	time.Sleep(3 * time.Second)

	return rand.Float64() * 100
}

func fetchPriceFromThirdSite() float64 {
	time.Sleep(2 * time.Second)

	return rand.Float64() * 100
}

func fetchAndSendMultiplePrices(priceChannel chan<- float64) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- price
	}
}
