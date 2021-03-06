package code

//暴力搜索法
func lengthOfLongestSubstring1(s string) int {
	len := len(s)
	isexit := [128]int8{0}
	max := 0
	for i := 0; i < len; i++ {
		isexit[s[i]] = 1
		num := 1
		for j := i + 1; j < len; j++ {
			if isexit[s[j]] != 1 {
				isexit[s[j]] = 1
				num++
			} else {
				break
			}
		}
		if num > max {
			max = num
		}
		isexit = [128]int8{0}
	}
	return max
}
func lengthOfLongestSubstring2(s string) int {
	len := len(s)
	isexit := [128]int8{0}
	max := 0
	i, j := 0, 0
	for i < len && j < len {
		if isexit[s[j]] == 0 {
			isexit[s[j]] = 1
			j++
			if j-i > max {
				max = j - i
			}
		} else {
			isexit[s[i]] = 0
			i++
		}
	}
	return max
}

//优化滑动窗口
func lengthOfLongestSubstring3(s string) int {
	len := len(s)
	isexit := [128]int{0}
	max := 0
	i, j := 0, 0
	for ; j < len; j++ {
		i = mmax(isexit[s[j]], i)
		max = mmax(max, j-i+1)
		isexit[s[j]] = j + 1
	}
	return max
}
func mmax(a, b int) int {
	if a > b {
		return a
	}
	return b
}
