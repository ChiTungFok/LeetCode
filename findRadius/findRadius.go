package findradius

import (
	"math"
	"sort"
)

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	var ans int
	for _, house := range houses {
		i := sort.SearchInts(heaters, house)
		minDis := math.MaxInt
		if i < len(heaters) {
			minDis = heaters[i] - house
		}
		if i-1 >= 0 && house-heaters[i-1] < minDis {
			minDis = house - heaters[i-1]
		}
		ans = max(minDis, ans)
	}
	return ans
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
