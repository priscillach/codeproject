package recursive

var parentheses []string

func generateParenthesis(n int) []string {
	parentheses = make([]string, 0)
	dfsGenerateParenthesis(n, "", 0, 0)
	return parentheses
}

func dfsGenerateParenthesis(n int, parenthesis string, left, right int) {
	if right == n {
		parentheses = append(parentheses, parenthesis)
	}
	if left < n {
		dfsGenerateParenthesis(n, parenthesis+"(", left+1, right)
	}
	if left > right {
		dfsGenerateParenthesis(n, parenthesis+")", left, right+1)
	}
}
