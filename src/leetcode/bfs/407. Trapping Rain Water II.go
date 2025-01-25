package bfs

import (
	"container/heap"
	"leetcode/src/utils/mathhelper"
)

func trapRainWater(heightMap [][]int) int {
	m := len(heightMap)
	n := len(heightMap[0])

	directions := [][]int{
		{0, 1},
		{0, -1},
		{1, 0},
		{-1, 0},
	}
	// go through the margin
	var buckets MinHeap
	visited := make([][]bool, m)
	for i := 0; i < m; i++ {
		visited[i][0] = true
		visited[i][n-1] = true
		visited[i] = make([]bool, n)
		buckets = append(buckets, &Bucket{
			Row:    i,
			Col:    0,
			Height: heightMap[i][0],
		})
		buckets = append(buckets, &Bucket{
			Row:    i,
			Col:    n - 1,
			Height: heightMap[i][n-1],
		})
	}
	for i := 0; i < n; i++ {
		visited[0][i] = true
		visited[m-1][i] = true
		buckets = append(buckets, &Bucket{
			Row:    0,
			Col:    i,
			Height: heightMap[0][i],
		})
		buckets = append(buckets, &Bucket{
			Row:    m - 1,
			Col:    i,
			Height: heightMap[m-1][i],
		})
	}
	heap.Init(&buckets)
	var res int

	for buckets.Len() > 0 {
		lowestBucket := heap.Pop(&buckets).(*Bucket)
		for _, direction := range directions {
			nextR, nextC := lowestBucket.Row+direction[0], lowestBucket.Col+direction[1]
			if nextR >= m || nextR < 0 || nextC >= n || nextC < 0 || visited[nextR][nextC] {
				continue
			}
			visited[nextR][nextC] = true
			res += mathhelper.Max(0, lowestBucket.Height-heightMap[nextR][nextC])
			heap.Push(&buckets, &Bucket{
				Row:    nextR,
				Col:    nextC,
				Height: mathhelper.Max(heightMap[nextR][nextC], lowestBucket.Height),
			})
		}
	}
	return res
}

type MinHeap []*Bucket

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i].Height < h[j].Height }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *MinHeap) Push(x interface{}) {
	*h = append(*h, x.(*Bucket))
}

func (h *MinHeap) Pop() interface{} {
	x := (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	return x
}

type Bucket struct {
	Row    int
	Col    int
	Height int
}
