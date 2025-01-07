package greedy

func findPrimeFactors(num int) []int {
	var factors []int
	for num%2 == 0 {
		factors = append(factors, 2)
		num = num / 2
	}
	i := 3
	for i*i <= num {
		for num%i == 0 {
			factors = append(factors, i)
			num /= i
		}
		i += 2
	}
	if num > 2 {
		factors = append(factors, num)
	}
	return factors
}
