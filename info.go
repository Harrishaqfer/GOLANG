package gofunctions

import (
	"fmt"
)

func Info(student_id string, s []Students, avg, sm map[string]float64, total float64) {
	for _, j := range s {
		if j.N_id == student_id {
			fmt.Printf("\nStudent ID:  %s Name: %s Phone no: %s\n Marks: %f", j.N_id, j.Name, j.Phno, sm[j.Name])
			for i, j := range avg {
				fmt.Printf("Average mark of %s is: %f\n", i, j/float64(len(s)))
			}
			fmt.Println("Total marks:", total)
		}
	}
}
