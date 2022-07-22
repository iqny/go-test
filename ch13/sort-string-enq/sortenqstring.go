package sort_string_enq

import (
	"sort"
)

func isRegroup(s1, s2 string) bool {
	sl1 := []rune(s1)
	sl2 := []rune(s2)
	if len(sl1) > 5000 || len(sl2) > 5000 {
		return false
	}
	sort.Slice(sl1, func(i, j int) bool {
		return sl1[i] < sl1[j]
	})
	sort.Slice(sl2, func(i, j int) bool {
		return sl2[i] < sl2[j]
	})
	for i, _ := range sl1 {
		if sl1[i] != sl2[i] {
			return false
		}
	}
	return true
}

