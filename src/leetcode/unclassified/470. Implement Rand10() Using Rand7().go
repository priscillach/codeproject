package unclassified

import (
	"math/rand"
	"time"
)

// https://leetcode.com/problems/implement-rand10-using-rand7/description/
func rand10() int {
	// generate 1 - 49 fairly
	x := (rand7()-1)*7 + rand7()
	for x > 40 {
		x = (rand7()-1)*7 + rand7()
	}
	return x%10 + 1
}

func rand7() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(7) + 1
}
