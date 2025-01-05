package _sort

// https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/description/
// finish times: 1
func reversePairs(record []int) int {
	_, count := mergeSort(record)
	return count
}

func mergeSort(record []int) ([]int, int) {
	if len(record) <= 1 {
		return record, 0
	}
	mid := len(record) / 2
	a, countA := mergeSort(record[:mid])
	b, countB := mergeSort(record[mid:])
	res, count := mergeTwoSlice(a, b)
	return res, count + countA + countB
}

func mergeTwoSlice(a, b []int) ([]int, int) {
	res := make([]int, 0, len(a)+len(b))
	var count int
	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i] <= b[j] {
			res = append(res, a[i])
			i++
		} else {
			res = append(res, b[j])
			// 注意这里选择右边的数组，a[i] > b[j]，因为是升序排序，说明a[i:len(a)]都比b[j]大
			count += len(a) - i
			j++
		}
	}

	res = append(res, a[i:]...)
	res = append(res, b[j:]...)
	return res, count
}
