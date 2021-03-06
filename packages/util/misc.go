package util

func Short(s string) string {
	if len(s) <= 6 {
		return s
	}
	return s[:6] + ".."
}

func ContainsDuplicates(lst []string) bool {
	for i := range lst {
		for j := i + 1; j < len(lst); j++ {
			if lst[i] == lst[j] {
				return true
			}
		}
	}
	return false
}

func IntersectsLists(lst1, lst2 []string) bool {
	if len(lst1) == 0 || len(lst2) == 0 {
		return false
	}
	for _, s1 := range lst1 {
		for _, s2 := range lst2 {
			if s1 == s2 {
				return true
			}
		}
	}
	return false
}

func ContainsInList(elem string, lst []string) bool {
	for _, s := range lst {
		if s == elem {
			return true
		}
	}
	return false
}
