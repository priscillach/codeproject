package recursive

var mFib map[int]int

func fib(n int) int {
	if mFib == nil {
		mFib = make(map[int]int)
	}
	if _, ok := mFib[n]; ok {
		return mFib[n]
	}
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	mFib[n] = fib(n-1) + fib(n-2)
	return mFib[n]
}
