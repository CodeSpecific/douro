package main

import (
	"github.com/CodeSpecific/douro/infra/algo"
)

func main() {
	count, amount := 10, 10000
	remainCount, remainAmount := count, amount
	sumAmount := 0
	for i := 0; i < count; i++ {
		curAmount := algo.RandomValue(remainCount-i, remainAmount)
		remainAmount -= curAmount
		sumAmount += curAmount
	}
}
