package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProfit(prices []int) int {

	var max int

	for i := 0; i < len(prices); i++ {
		for j := i + 1; j < len(prices); j++ {
			if prices[i] < prices[j] {
				if (prices[j]-prices[i] >= max) {
					max = prices[j] - prices[i]
				}
			}
		}
	}

	return max

}

func TestMaxProfit(t *testing.T) {
	cases := []struct {
		Prices []int
		Want   int
	}{
		{Prices: []int{7, 1, 5, 3, 6, 4}, Want: 5},
		{Prices: []int{7, 6, 4, 3, 1}, Want: 0},
	}

	for _, c := range cases {
		got := maxProfit(c.Prices)
		msg := fmt.Sprintf("expected %d, got %d, prices: %v", c.Want, got, c.Prices)
		assert.Equal(t, c.Want, got, msg)
	}
}
