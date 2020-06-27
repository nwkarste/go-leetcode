package utils

import ()

func knightProbability(N int, K int, r int, c int) float64 {
	probChart := make([][]float64, N)
	for i := 0; i < N; i++ {
		probChart[i] = make([]float64, N)
	}
	probChart[r][c] = 1
	moves := [][]int{{1, 2}, {1, -2}, {-1, 2}, {-1, -2}, {2, 1}, {2, -1}, {-2, 1}, {-2, -1}}
	var probOOB float64
	probOOB = 0

	for jump := 0; jump < K; jump++ {
		nextProbs := make([][]float64, N)
		for i := 0; i < N; i++ {
			nextProbs[i] = make([]float64, N)
		}
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				for move := 0; move < 8; move++ {
					row := i + moves[move][0]
					col := j + moves[move][1]
					if row < 0 || row > N-1 || col < 0 || col > N-1 {
						probOOB += probChart[i][j] / 8
					} else {
						nextProbs[row][col] += probChart[i][j] / 8
					}
				}
			}
		}
		probChart = nextProbs
	}
	return 1 - probOOB
}
