package anagram

type runeCounts map[rune]int

func countChars(s string) runeCounts {
	runeCounts := make(runeCounts)
	for _, c := range s {
		runeCounts[c] += 1
	}
	return runeCounts
}

func isCountsEqual(m1 runeCounts, m2 runeCounts) bool {
	if len(m1) != len(m2) {
		return false
	}
	for key, count := range m1 {
		if m2[key] != count {
			return false
		}
	}
	return true
}

func IsAnagram(s1 string, s2 string) bool {
	s1Counts := countChars(s1)
	s2Counts := countChars(s2)
	return isCountsEqual(s1Counts, s2Counts)
}
