package main

import (
	"sort"
	"testing"
)

func TestThreeSum(t *testing.T) {
	test1 := []int{-1, 0, 1, 2, -1, -4}
	result1 := threeSum(test1)

	if len(result1) != 2 {
		t.Errorf("Expected 2, got %d", len(result1))
	}
}

func threeSum(nums []int) [][]int {
	// result := [][]int{}
	// comps는 키가 []int{int, int, int}이고 value가 bool인 map

	// comps := map[string]bool{}

	comps := map[int]map[int]map[int]bool{}

	answers := [][]int{}

	for index1, i := range nums[:len(nums)-2] {
		if comps[i] == nil {
			comps[i] = map[int]map[int]bool{}
		}
		if i == 0 {
			if nums[index1+1] == 0 && nums[index1+2] == 0 {

				if comps[0][0] == nil {
					comps[0][0] = map[int]bool{}
				}
				if !comps[0][0][0] {
					answers = append(answers, []int{0, 0, 0})
					comps[0][0][0] = true
				}
			}

		}

		for index2 := index1 + 1; index2 < len(nums)-1; index2++ {
			j := nums[index2]
			for index3 := len(nums) - 1; index3 > index2; index3-- {

				k := nums[index3]
				if i+j+k == 0 {
					comp := []int{i, j, k}
					sort.Ints(comp)
					if comps[comp[0]] == nil {
						comps[comp[0]] = map[int]map[int]bool{}
					}

					if comps[comp[0]][comp[1]] == nil {
						comps[comp[0]][comp[1]] = map[int]bool{}
					}
					// comps[comp] = true
					if !comps[comp[0]][comp[1]][comp[2]] {
						answers = append(answers, comp)
						comps[comp[0]][comp[1]][comp[2]] = true
					}
					// comps[fmt.Sprintf("%d|%d|%d", comp[0], comp[1], comp[2])] = true

					// comps[comp[0]][comp[1]][comp[2]] = true
				}
			}
		}
	}

	return answers
	// return result

}
