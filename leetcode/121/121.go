package main

import "fmt"

func main() {
	// prices := []int{7, 1, 5, 3, 6, 4} //5
	// prices := []int{7, 6, 4, 3, 1} //0
	// prices := []int{2, 4, 1} //2
	// prices := []int{2, 1, 4} //3
	// prices := []int{1, 2, 4} //3
	prices := []int{9, 5, 7, 4, 2, 4, 1, 6, 4} //5
	fmt.Println(maxProfit(prices))
}

func maxProfit(prices []int) int {
	maximumProfit := 0
	// currBuyPrice := 0
	currSalePrice := 0
	i := 0
	j := 1

	for j < len(prices) && i < len(prices) {

		if prices[j] < currSalePrice && j != len(prices)-1 {
			j++
		}

		// if prices[i] > currBuyPrice {
		// 	i++
		// }

		if (prices[i] < prices[j]) && (prices[j]-prices[i] > maximumProfit) {
			// currBuyPrice = prices[i]
			currSalePrice = prices[j]
			maximumProfit = prices[j] - prices[i]
		}

		if j == len(prices)-1 && i != len(prices)-1 {
			i++
			j = i + 1
			currSalePrice = 0
			continue
		}

		j++
	}

	return maximumProfit
}
