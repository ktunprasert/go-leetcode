package bestbuystock

func maxProfit(prices []int) int {
	var min_idx, max_idx int
	var maxProfit, profit int
	for i, price := range prices {
		if price < prices[min_idx] {
			profit = prices[max_idx] - prices[min_idx]
			if profit > maxProfit {
				maxProfit = profit
			}
			min_idx = i
			max_idx = i
		}
		if price > prices[max_idx] && i > min_idx {
			max_idx = i
		}
	}
	profit = prices[max_idx] - prices[min_idx]
	if profit > maxProfit {
		return profit
	}
	return maxProfit
}
