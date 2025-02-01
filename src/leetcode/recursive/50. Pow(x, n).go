package recursive

// https://leetcode.com/problems/powx-n/description/
var m map[int]float64

func myPow(x float64, n int) float64 {
	m = make(map[int]float64)
	return dfs(x, n)
}

func dfs(x float64, n int) float64 {
	if v, ok := m[n]; ok {
		return v
	}
	if n == 0 {
		return 1
	}
	if n == -1 {
		return 1 / x
	}
	if n == 1 {
		return x
	}
	tmp := dfs(x, n/2) * dfs(x, n-(n/2))
	m[n] = tmp
	return tmp
}
