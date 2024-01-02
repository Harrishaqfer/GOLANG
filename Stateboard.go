package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type SBstudents struct {
	name  string
	n_id  string
	phno  string
	grade string
	total float64
}

type SBmarks struct {
	n_id string
	s_id string
	mark float64
}

func createf(path string) [][]string {
	fil, err := os.Open(path)
	if err != nil {
		fmt.Println("error")
	}

	defer fil.Close()
	fileReader := csv.NewReader(fil)
	records, error := fileReader.ReadAll()

	if error != nil {
		fmt.Println(error)
	}
	return records
}
func grading(s []SBstudents, m []SBmarks) []SBstudents {
	for i, j := range s {
		tot := float64(0)
		for _, i := range m {

			if j.n_id == i.n_id {
				tot = tot + i.mark
			}
		}

		s[i].total = tot

		if tot >= 570 {
			s[i].grade = "A"
		} else if tot >= 530 && tot < 570 {
			s[i].grade = "B"
		} else if tot >= 500 && tot < 520 {
			s[i].grade = "B"
		} else {
			s[i].grade = "D"
		}

	}
	return s

}
func Average(m []SBmarks, s []SBstudents) (map[string]float64, float64) {
	esa := map[string]float64{}
	var total float64 = 0
	for _, j := range m {
		total = total + j.mark

		esa[j.s_id] = esa[j.s_id] + j.mark
	}

	fmt.Printf("Total average mark is:%f\n\n", total/float64(len(s)))
	for i, j := range esa {
		fmt.Printf("Average mark of %s is: %f", i, j/float64(len(s)))
	}

	return esa, total
}

func main() {

	studentF := createf("student.csv")
	marksF := createf("marks.csv")
	var s []SBstudents
	var m []SBmarks

	for _, j := range studentF {

		s = append(s, SBstudents{name: j[0], n_id: j[1], phno: j[2], grade: "ki", total: 0})
	}
	for _, j := range marksF {

		f, _ := strconv.ParseFloat(j[2], 8)
		m = append(m, SBmarks{n_id: j[0], s_id: j[1], mark: f})
	}

	graded := grading(s, m)
	fmt.Println(graded)

	student_name := "John Doe"

	for _, j := range graded {

		if j.name == student_name || j.n_id == student_name {
			fmt.Println(j.name, " ", j.n_id, " ", j.phno, " ", j.grade, " ", j.total)

			for _, i := range m {
				if i.n_id == j.n_id {
					fmt.Println(i.s_id, ":", i.mark)
				}
			}
		}

	}
	esa, tot1 := Average(m, s)
	fmt.Println(esa, tot1)

}
