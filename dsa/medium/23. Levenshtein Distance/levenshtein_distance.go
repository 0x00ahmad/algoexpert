package main

// O(nm) time | O(nm) space
// we can optimize the timespace further by having min(n, m) but I just went with this one
// iterative approach using tabulation
func LevenshteinDistance(a, b string) int {
	n := len(a)
	m := len(b)

	if n == 0 {
		return m
	}

	if m == 0 {
		return n
	}

	d := make([][]int, n+1)
	for i := 0; i <= n; i++ {
		d[i] = make([]int, m+1)
		d[i][0] = i
	}

	for j := 0; j <= m; j++ {
		d[0][j] = j
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= m; j++ {
			if a[i-1] == b[j-1] {
				d[i][j] = d[i-1][j-1]
			} else {
				d[i][j] = min(d[i-1][j], d[i][j-1], d[i-1][j-1]) + 1
			}
		}
	}

	return d[n][m]
}

func min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
		return c
	}
	if b < c {
		return b
	}
	return c
}
