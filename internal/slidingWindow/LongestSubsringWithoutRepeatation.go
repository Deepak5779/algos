package slidingwindow

func findLongestSubStr(s string) string {
	maxLen := 0
	iStart := 0
	iEnd := 0

	lookup := map[byte]int{}
	for i := 0; i < len(s); i++ {
		if v, ok := lookup[s[i]]; ok {

		}
	}
}
