package gofunctions

import (
	"fmt"
	"sort"
)

func Order_vise(sm map[string]float64, s []Students) {
	var mark_arr []float64

	for _, j := range sm {

		mark_arr = append(mark_arr, j)

	}

	sort.Float64s(mark_arr)

	stu_id := map[string]string{}
	for _, j := range s {
		stu_id[j.Name] = j.N_id
	}
	for i := len(mark_arr) - 1; i >= 0; i-- {
		for l, k := range sm {
			if mark_arr[i] == k {
				fmt.Printf("%s ID:%s Mark:%f\n", l, stu_id[l], k)
			}
		}
	}
}
