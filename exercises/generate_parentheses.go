package exercises

// var branchLeft []byte = []byte("(")
// var branchRight []byte = []byte(")")

func GenerateParentheses(n int) []string {

	return getBranchValidCombinations("(", 1, n)

}

func getBranchValidCombinations(current string, remainingOpenParentheses int, targetSize int) []string {

	if remainingOpenParentheses < 0 {
		return []string{}
	}

	if targetSize-len(current) < remainingOpenParentheses {
		return []string{}
	}

	if len(current) == targetSize {
		return []string{current}
	}

	return append(
		getBranchValidCombinations(current+"(", remainingOpenParentheses+1, targetSize),
		getBranchValidCombinations(current+")", remainingOpenParentheses-1, targetSize)...,
	)

}
