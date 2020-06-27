package utils

import (
	"sort"
)

func findOrder(numCourses int, prerequisites [][]int) []int {
	order := make(map[int][]int)
	rankings := make(map[int]int)
	for i := 0; i < numCourses; i++ {
		rankings[i] = 0
	}
	for _, pair := range prerequisites {
		order[pair[1]] = append(order[pair[1]], pair[0])
		if rankings[pair[0]] < 1+rankings[pair[1]] {
			rankings[pair[0]] = 1 + rankings[pair[1]]
			cascade(order, rankings, pair[0])
		}
	}
	courseRanks := make([][]int, numCourses)
	i := 0
	for course, prerequisites := range rankings {
		courseRanks[i] = make([]int, 2)
		courseRanks[i][0] = prerequisites
		courseRanks[i][1] = course
		i = i + 1
	}
	sort.Slice(courseRanks, func(i, j int) bool {
		return courseRanks[i][0] < courseRanks[j][0]
	})
	finalOrder := make([]int, numCourses)
	for i, courseRank := range courseRanks {
		finalOrder[i] = courseRank[1]
	}
	return finalOrder
}

func cascade(order map[int][]int, rank map[int]int, course int) {
	for _, dependent := range order[course] {
		if rank[dependent] < rank[course]+1 {
			rank[dependent] = 1 + rank[course]
			cascade(order, rank, dependent)
		}
	}
}
