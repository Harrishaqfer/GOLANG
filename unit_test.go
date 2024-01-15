package main

import (
	Go "godaaay2/gofunctions" // This is the package i have created. gofunctions package has six files each with a funcion.

	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_my_functions1(t *testing.T) {

	studentF := Go.Createf("student.csv") // This function will convert the csv file into golang data structure (slice of slice of strings)
	marksF := Go.Createf("marks.csv")

	s, m := Go.Initialize(studentF, marksF) // This function  will initialize the files data to the structures fields (Hardcoded)

	avg, total := Go.Average(m, s) // This is the fuction to be tested.

	expected_avg := map[string]float64{"E": 255.8, "M": 255.8, "S": 255.8}
	expected_total := float64(767.400000)

	assert.Equal(t, expected_avg, avg, "total should match") // Here im checking whether both the actual and expected outputs are same.
	assert.Equal(t, expected_total, total, "total should match")

}
func Test_my_functions2(t *testing.T) {

	studentF := Go.Createf("student.csv")

	marksF := Go.Createf("marks.csv")

	s, m := Go.Initialize(studentF, marksF)
	sm := Go.Student_vise_marks(m, s)

	ex := map[string]float64{"Alice Johnson": 234, "Jane Smith": 276.9, "John Doe": 256.5}
	assert.Equal(t, ex, sm, "total should match")

}
