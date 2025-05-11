package main

import (
	"fmt"
	"price-search/internal/fetcher"
	"price-search/internal/models"
	"price-search/internal/processor"
	"time"
)

func main() {
	start := time.Now()

	priceChannel := make(chan models.PriceDetail)
	done := make(chan bool)

	go fetcher.FetchPrices(priceChannel)
	go processor.ShowPricesAndAVG(priceChannel, done)

	<-done

	fmt.Printf("\nTempo Total: %s", time.Since(start))
}
