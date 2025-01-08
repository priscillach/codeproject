package string_resolve

func swapString(a, b string) (string, string) {
	a = a + b
	b = a[:len(a)-len(b)]
	a = a[len(b):]
	return a, b
}
