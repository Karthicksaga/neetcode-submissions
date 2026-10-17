func maxProfit(prices []int) int {

	if len(prices) == 0{
		return 0
	}
	minimumPrice, maxProfit := prices[0], 0

	for _, price := range prices{

		if price < minimumPrice{
			minimumPrice = price 
		}

		currentProfit := price - minimumPrice 

		if currentProfit > maxProfit{
			maxProfit = currentProfit 
		}

	}
	
	return maxProfit
}
