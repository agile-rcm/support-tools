package tools

// Levenshtein algorithm calculates and returns the Levenshtein distance between 2 strings and is used to compare
// the similarity between 2 strings. This can be used to determine if a provided name (string) maybe just a typo.
// https://www.golangprograms.com/golang-program-for-implementation-of-levenshtein-distance.html#:~:text=Golang%20program%20for%20implementation%20of,to%20transform%20s%20into%20t.
func Levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = minimum(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func minimum(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}
