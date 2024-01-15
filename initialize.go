package gofunctions

import "strconv"

type Students struct {
	Name string
	N_id string
	Phno string
}
type Marks struct {
	N_id string
	S_id string
	Mark float64
}

func Initialize(studentF, marksF [][]string) ([]Students, []Marks) {

	var s []Students
	var m []Marks
	for _, j := range studentF {

		s = append(s, Students{Name: j[0], N_id: j[1], Phno: j[2]})
	}
	for _, j := range marksF {

		f, _ := strconv.ParseFloat(j[2], 8)
		m = append(m, Marks{N_id: j[0], S_id: j[1], Mark: f})
	}
	return s, m

}
