package bit

import "errors"

func modWithBitCalculator(m, n int) (int, error) {
	if n == 0 {
		return 0, errors.New("除数不能为0")
	}

	// 如果 n 是 2 的幂，直接使用位运算
	if (n & (n - 1)) == 0 {
		// m%n=m−(m//n)∗n
		// n = 2^k，n-1则其低k位全为1
		// (m // n) * n -> (m >> k) << k，即舍弃了m的低k位值
		// m - (m // n) * n，即m的低k位值，
		return m & (n - 1), nil
	}

	div := 1
	for (n << div) <= m {
		div++
	}

	for div >= 0 {
		if m >= (n << div) {
			m -= n << div
		}
		div--
	}

	return m, nil
}
