package utils

import (
	"strings"
	"unicode"

	"github.com/antzucaro/matchr"
)

func StringSimilarity(a, b string) float64 {
	na := normalize(a)
	nb := normalize(b)
	jw := matchr.JaroWinkler(na, nb, true)
	jc := jaccard(na, nb)
	return 0.6*jw + 0.4*jc
}

func RemoveParentheses(raw string) string {
	s := strings.ReplaceAll(raw, "(", "")
	s = strings.ReplaceAll(s, ")", "")
	s = strings.ReplaceAll(s, "[", "")
	s = strings.ReplaceAll(s, "]", "")
	s = strings.ReplaceAll(s, "{", "")
	s = strings.ReplaceAll(s, "}", "")
	return s
}

func NormalizeHyphens(s string) string {
	s = strings.ReplaceAll(s, "–", "-")
	s = strings.ReplaceAll(s, "—", "-")
	s = strings.ReplaceAll(s, "-", "-")
	s = strings.ReplaceAll(s, ": ", "- ")
	s = strings.ReplaceAll(s, " : ", "- ")
	return s
}

func CollapseSpaces(s string) string {
	return strings.Join(strings.Fields(s), " ")
}

func normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.Map(func(r rune) rune {
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			return r
		}
		return -1
	}, s)
	return CollapseSpaces(s)
}

func jaccard(a, b string) float64 {
	wa := strings.Fields(a)
	wb := strings.Fields(b)
	setA, setB := map[string]struct{}{}, map[string]struct{}{}
	for _, t := range wa {
		setA[t] = struct{}{}
	}
	for _, t := range wb {
		setB[t] = struct{}{}
	}

	inter, union := 0, len(setA)
	for k := range setB {
		if _, ok := setA[k]; ok {
			inter++
		} else {
			union++
		}
	}
	if union == 0 {
		return 0
	}
	return float64(inter) / float64(union)
}
