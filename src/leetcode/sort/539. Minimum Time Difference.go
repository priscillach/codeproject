package _sort

import (
	"leetcode/src/utils/mathhelper"
	"math"
	"sort"
	"strconv"
	"strings"
)

// https://leetcode.com/problems/minimum-time-difference/description/
func findMinDifference(timePoints []string) int {
	sort.Strings(timePoints)
	minDistance := math.MaxInt
	for i := 0; i < len(timePoints); i++ {
		if i == len(timePoints)-1 {
			minDistance = mathhelper.Min(minDistance, computeMinDifference(timePoints[0], timePoints[i]))
			continue
		}
		minDistance = mathhelper.Min(minDistance, computeMinDifference(timePoints[i], timePoints[i+1]))
	}
	return minDistance
}

func computeMinDifference(time1, time2 string) int {
	time1Slice := strings.Split(time1, ":")
	time2Slice := strings.Split(time2, ":")
	t1Hour, _ := strconv.Atoi(time1Slice[0])
	t1Min, _ := strconv.Atoi(time1Slice[1])
	t2Hour, _ := strconv.Atoi(time2Slice[0])
	t2Min, _ := strconv.Atoi(time2Slice[1])
	time1Sum := t1Hour*60 + t1Min
	time2Sum := t2Hour*60 + t2Min
	dis := mathhelper.Abs(time1Sum - time2Sum)
	return mathhelper.Min(dis, 24*60-dis)
}
