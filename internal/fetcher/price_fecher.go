package fetcher

import (
	"math/rand"
	"price-search/internal/models"
	"sync"
	"time"
)

func FetchPrices(priceChannel chan<- models.PriceDetail) {
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

func fetchPriceFromFirstSite() models.PriceDetail {
	time.Sleep(1 * time.Second)

	return models.PriceDetail{
		StoreName: "Site A",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func fetchPriceFromSecondSite() models.PriceDetail {
	time.Sleep(3 * time.Second)

	return models.PriceDetail{
		StoreName: "Site B",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func fetchPriceFromThirdSite() models.PriceDetail {
	time.Sleep(2 * time.Second)

	return models.PriceDetail{
		StoreName: "Site C",
		Value:     rand.Float64() * 100,
		Timestamp: time.Now(),
	}
}

func fetchAndSendMultiplePrices(priceChannel chan<- models.PriceDetail) {
	time.Sleep(6 * time.Second)
	prices := []float64{
		rand.Float64() * 100,
		rand.Float64() * 100,
		rand.Float64() * 100,
	}

	for _, price := range prices {
		priceChannel <- models.PriceDetail{
			StoreName: "Site D",
			Value:     price,
			Timestamp: time.Now(),
		}
	}
}
