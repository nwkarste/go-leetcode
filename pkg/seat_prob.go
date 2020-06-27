package utils

import ()

func nthPersonGetsNthSeat(n int) float64 {
	probSeatTaken := make([]float64, n)
	for lcv := 0; lcv < n; lcv++ {
		probSeatTaken[lcv] = 1 / float64(n)
	}
	for lcv := 1; lcv < n; lcv++ {
		for i := lcv + 1; i < n; i++ {
			probSeatTaken[i] += probSeatTaken[lcv] / (float64(n) - float64(lcv))
		}
	}
	return probSeatTaken[n-1]
}
