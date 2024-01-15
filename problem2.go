//Problem 2
package main

import (
	"fmt"
	Go "godaaay2/gofunctions"
)

func main() {

	studentF := Go.Createf("student.csv")
	marksF := Go.Createf("marks.csv")
	s, m := Go.Initialize(studentF, marksF)
	avg, total := Go.Average(m, s)
	fmt.Println(avg)
	sm := Go.Student_vise_marks(m, s)
	fmt.Println(sm, "\n\n")
	//Go.Order_vise(sm, s)
	//student_id := "S1"
	//Go.Info(student_id, s, avg, sm, total)
	fmt.Println(avg, total)

}
