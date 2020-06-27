package utils

func isInterleave(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	mem := make([][]bool, l1+1)
	for i := 0; i < l1+1; i++ {
		mem[i] = make([]bool, l2+1)
	}
	mem[0][0] = true
	for i := 0; i < l3; i++ {
		for j := 0; j <= i; j++ {
			if j < l1 && (i-j) < l2+1 && s1[j] == s3[i] && mem[j][i-j] {
				mem[j+1][i-j] = true
			}
			if j < l2 && (i-j) < l1+1 && s2[j] == s3[i] && mem[i-j][j] {
				mem[i-j][j+1] = true
			}
		}
	}
	return mem[l1][l2]
}

func isInterleave3(s1 string, s2 string, s3 string) bool {
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	mem := make([][][]bool, l1+1)
	for i := 0; i < l1+1; i++ {
		mem[i] = make([][]bool, l2+1)
		for j := 0; j < l2+1; j++ {
			mem[i][j] = make([]bool, l3+1)
		}
	}
	mem[0][0][0] = true
	for i := 0; i < l3; i++ {
		for j := 0; j <= i; j++ {
			if j < l1 && (i-j) < l2+1 && s1[j] == s3[i] && mem[j][i-j][i] {
				mem[j+1][i-j][i+1] = true
			}
			if j < l2 && (i-j) < l1+1 && s2[j] == s3[i] && mem[i-j][j][i] {
				mem[i-j][j+1][i+1] = true
			}
		}
	}
	return mem[l1][l2][l3]
}

func isInterleave2(s1 string, s2 string, s3 string) bool {
	var potentialInterleaves [][]int
	l1 := len(s1)
	l2 := len(s2)
	l3 := len(s3)
	if (l1 + l2) != l3 {
		return false
	}
	potentialInterleaves = append(potentialInterleaves, []int{0, 0, 0})
	for len(potentialInterleaves) > 0 {
		state := potentialInterleaves[len(potentialInterleaves)-1]
		potentialInterleaves = potentialInterleaves[:len(potentialInterleaves)-1]
		p1 := state[0]
		p2 := state[1]
		p3 := state[2]
		if p1 == l1 && p2 == l2 && p3 == l3 {
			return true
		}
		if p1 < l1 && s1[p1] == s3[p3] {
			nextInterleave := []int{p1 + 1, p2, p3 + 1}
			potentialInterleaves = append(potentialInterleaves, nextInterleave)
		}
		if p2 < l2 && s2[p2] == s3[p3] {
			nextInterleave := []int{p1, p2 + 1, p3 + 1}
			potentialInterleaves = append(potentialInterleaves, nextInterleave)
		}
	}
	return false
}
