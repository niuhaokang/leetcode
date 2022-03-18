package main

func maxProduct(words []string) int {
	binArr := make([]int64, len(words))

	for i, word := range words {
		for _, c := range word {
			binArr[i] |= 1 << (c - 'a')
		}
	}

	ans := 0
	for i:=0; i<len(binArr); i++ {
		for j:=i+1; j<len(binArr); j++ {
			if (binArr[i] & binArr[j]) == 0 {
				ans = max(ans, len(words[i]) * len(words[j]))
			}
		}
	}

	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}