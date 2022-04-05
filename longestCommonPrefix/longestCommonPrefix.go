package longestCommonPrefix

func longestCommonPrefix(strs []string) string {
	prefix := ""
	if len(strs) == 0 {
		return prefix
	}
	if len(strs) == 1 {
		return strs[0]
	}
	for index, r := range strs[0] {
		for i := 1; i < len(strs); i++ {
			if index >= len(strs[i]) {
				return prefix
			}
			if strs[i][index] != byte(r) {
				return prefix
			}
			continue
		}
		prefix += string(r)
	}
	return prefix
}
