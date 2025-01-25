package _heap

import (
	"container/heap"
	"leetcode/src/define/myheap"
	"sort"
)

// https://leetcode.com/problems/course-schedule-iii/description/
func scheduleCourse(courses [][]int) int {
	sort.Slice(courses, func(i, j int) bool {
		return courses[i][1] < courses[j][1]
	})
	curDay := 0
	h := &myheap.MaxHeap{}
	heap.Init(h)
	for _, course := range courses {
		heap.Push(h, course[0])
		curDay += course[0]
		if curDay > course[1] {
			// remove a course cost time more than cur course, the cur day decreases and the last day extends, which make the courses in heap must meet the last day
			curDay -= heap.Pop(h).(int)
		}
	}
	return h.Len()
}
