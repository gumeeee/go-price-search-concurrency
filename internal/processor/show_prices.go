package processor

import (
	"fmt"
)

func ShowPricesAndAVG(priceChannel chan float64) {
	var totalPrice float64
	countPrices := 0.0
	for price := range priceChannel {
		totalPrice += price
		countPrices++
		avgPrice := totalPrice / countPrices
		fmt.Printf("Price received: R$ %.2f | Average Price at the moment: R$ %.2f\n", price, avgPrice)
	}
	fmt.Println("Consumer goroutine finished. Channel closed.")
}
