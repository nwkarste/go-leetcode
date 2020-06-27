package utils

import ()

func permute(nums []int) [][]int {
	var queue []node
	initialPerm := make([]int, len(nums))
	initialTaken := make([]bool, len(nums))
	initialNode := node{
		perm:  initialPerm,
		taken: initialTaken,
		total: 0,
	}
	var final [][]int
	queue = append(queue, initialNode)
	for _, num := range nums {
		queueLen := len(queue)
		for lcv := 0; lcv < queueLen; lcv++ {
			nodeOld := queue[0]
			queue = queue[1:]
			for j := 0; j < len(nums); j++ {
				perm := make([]int, len(nums))
				taken := make([]bool, len(nums))
				copy(perm, nodeOld.perm)
				copy(taken, nodeOld.taken)
				if !taken[j] {
					perm[j] = num
					taken[j] = true
					if nodeOld.total+1 == len(nums) {
						final = append(final, perm)
					} else {
						newNode := node{
							perm:  perm,
							taken: taken,
							total: nodeOld.total + 1,
						}
						queue = append(queue, newNode)
					}
				}
			}
		}
	}
	if len(nums) == 0 {
		final = append(final, []int{})
	}
	return final
}

type node struct {
	perm  []int
	taken []bool
	total int
}
