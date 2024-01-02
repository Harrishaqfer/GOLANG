package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type students struct {
	name string
	n_id string
	phno string
}
type marks struct {
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
func Average(m []marks, s []students) (map[string]float64, float64) {
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
func student_vise_marks(m []marks, s []students) map[string]float64 {
	sm := map[string]float64{}

	for _, j := range s {
		for _, k := range m {
			if k.n_id == j.n_id {

				sm[j.name] = sm[j.name] + k.mark
			}

		}

	}
	fmt.Println(sm)
	return sm
}
func order_vise(sm map[string]float64, s []students) {
	var mark_arr []float64

	for _, j := range sm {

		mark_arr = append(mark_arr, j)

	}

	sort.Float64s(mark_arr)

	stu_id := map[string]string{}
	for _, j := range s {
		stu_id[j.name] = j.n_id
	}
	for i := len(mark_arr) - 1; i >= 0; i-- {
		for l, k := range sm {
			if mark_arr[i] == k {
				fmt.Printf("%s ID:%s Mark:%f\n", l, stu_id[l], k)
			}
		}
	}
}

func info(student_id string, s []students, avg, sm map[string]float64, total float64) {
	for _, j := range s {
		if j.n_id == student_id {
			fmt.Printf("\nStudent ID:  %s Name: %s Phone no: %s\n Marks: %f", j.n_id, j.name, j.phno, sm[j.name])
			for i, j := range avg {
				fmt.Printf("Average mark of %s is: %f\n", i, j/float64(len(s)))
			}
			fmt.Println("Total marks:", total)
		}
	}
}
func main() {

	studentF := createf("student.csv")
	marksF := createf("marks.csv")
	var s []students
	var m []marks

	for _, j := range studentF {

		s = append(s, students{name: j[0], n_id: j[1], phno: j[2]})
	}
	for _, j := range marksF {

		f, _ := strconv.ParseFloat(j[2], 8)
		m = append(m, marks{n_id: j[0], s_id: j[1], mark: f})
	}
	fmt.Println(m)

	avg, total := Average(m, s)
	sm := student_vise_marks(m, s)
	order_vise(sm, s)
	student_id := "S1"
	info(student_id, s, avg, sm, total)

}
