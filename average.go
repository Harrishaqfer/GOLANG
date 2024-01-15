package gofunctions

import (
	"fmt"
)

func Average(m []Marks, s []Students) (map[string]float64, float64) {
	esa := map[string]float64{}
	var total float64 = 0
	for _, j := range m {
		total = total + j.Mark

		esa[j.S_id] = esa[j.S_id] + j.Mark
	}

	fmt.Printf("Total average mark is:%f\n\n", total/float64(len(s)))
	for i, j := range esa {
		fmt.Printf("Average mark of %s is: %f", i, j/float64(len(s)))
	}

	return esa, total
}
