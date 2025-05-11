package fetcher

import (
	"math/rand"
	"time"
)

func FetchPriceFromFirstSite() float64 {
	time.Sleep(1 * time.Second)

	return rand.Float64() * 100
}

func FetchPriceFromSecondSite() float64 {
	time.Sleep(3 * time.Second)

	return rand.Float64() * 100
}

func FetchPriceFromThirdSite() float64 {
	time.Sleep(2 * time.Second)

	return rand.Float64() * 100
}
