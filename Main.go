package main

func main() {
	println(maxProfit3([]int{1, 2, 3, 4}))
	println(maxProfit3([]int{7, 1, 5, 3, 6, 4}))
	println(maxProfit3([]int{7, 6, 4, 3, 1}))
	println(maxProfit3([]int{3, 7, 1, 5, 6, 3, 4, 2, 0, 7}))
	println(maxProfit3([]int{1, 0}))
	println(maxProfit3([]int{0}))
	println(maxProfit3([]int{4, 6, 8, 0, 5}))
}

func maxProfit(prices []int) int {
	var minimum = 0
	var profit = 0
	for i := 0; i < len(prices); i++ {
		if i != 0 && prices[i] >= minimum {
			continue
		}
		minimum = prices[i]
		for j := i + 1; j < len(prices); j++ {
			var maximum = prices[j]
			if (maximum - minimum) > profit {
				profit = maximum - minimum
			}
		}
	}
	return profit
}

func maxProfit2(prices []int) int {
	var minimum = 0
	var profit = 0
	var maximum = 0
	var maximumIndex = 0
	for i := 0; i < len(prices); i++ {
		if i != 0 && prices[i] >= minimum {
			continue
		}
		minimum = prices[i]
		if maximumIndex != 0 && maximumIndex > i {
			profit = maximum - minimum
			continue
		}
		for j := i + 1; j < len(prices); j++ {
			var localMaximum = prices[j]
			if (localMaximum - minimum) > profit {
				profit = localMaximum - minimum
				maximum = localMaximum
				maximumIndex = j
			}
		}
	}
	return profit
}

func maxProfit3(prices []int) int {
	var buyPrice = prices[0]
	profit := 0
	for i := 0; i < len(prices); i++ {
		if buyPrice > prices[i] {
			buyPrice = prices[i]
		}
		profit = max(profit, prices[i]-buyPrice)
	}
	return profit
}
