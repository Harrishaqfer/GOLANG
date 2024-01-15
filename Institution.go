package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type institution interface {
	student_info(par1 [][]string, par2 [][]string) (map[string]interface{}, map[string]float64, float64, []string, []string, []float64)
	Print_institute_report(name_arr, n_id_arr []string, mark_arr []float64, info map[string]interface{}, esa map[string]float64, tot float64)
	student_report(par1 map[string]interface{}, par2 string)
}

type SBstudents struct {
	name     string
	n_id     string
	phno     string
	grade    string
	total    float64
	submarks map[string]float64
}

type CBSEstudents struct {
	name      string
	phno      string
	grade     string
	total     float64
	sub_marks map[string]float64
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

func (s SBstudents) student_info(studentF [][]string, marksF [][]string) (map[string]interface{}, map[string]float64, float64, []string, []string, []float64) {
	info := map[string]interface{}{}
	esa := map[string]float64{}
	total := float64(0)
	var name_arr []string
	var n_id_arr []string
	var mark_arr []float64

	for _, j := range studentF {
		name_arr = append(name_arr, j[0])
		n_id_arr = append(n_id_arr, j[1])
		s.name = j[0]
		s.phno = j[2]
		tot := float64(0)
		temp := map[string]float64{}
		for _, k := range marksF {
			if j[1] == k[0] {
				f, _ := strconv.ParseFloat(k[2], 8)
				tot = tot + f
				temp[k[1]] = temp[k[1]] + f
				esa[k[1]] = esa[k[1]] + f
			}
		}
		s.submarks = temp
		s.total = tot
		mark_arr = append(mark_arr, tot)
		total = total + tot
		if tot >= 570 {
			s.grade = "A"
		} else if tot >= 530 && tot < 570 {
			s.grade = "B"
		} else if tot >= 500 && tot < 520 {
			s.grade = "B"
		} else {
			s.grade = "D"
		}
		info[j[1]] = s
	}
	sort.Strings(name_arr)
	sort.Strings(n_id_arr)
	sort.Sort(sort.Reverse(sort.Float64Slice(mark_arr)))
	return info, esa, total, name_arr, n_id_arr, mark_arr
}

func (s CBSEstudents) student_info(studentF, marksF [][]string) (map[string]interface{}, map[string]float64, float64, []string, []string, []float64) {
	info := map[string]interface{}{}
	esa := map[string]float64{}
	tot := float64(0)
	t := float64(0)

	var name_arr []string
	var n_id_arr []string
	var mark_arr []float64

	for _, i := range studentF {
		name_arr = append(name_arr, i[0])
		n_id_arr = append(n_id_arr, i[1])
		s.name = i[0]
		s.phno = i[2]
		t = 0
		temp := map[string]float64{}
		for _, j := range marksF {

			if i[1] == j[0] {
				f, _ := strconv.ParseFloat(j[3], 8)
				temp[j[1]] = temp[j[1]] + f
			}
		}
		s.sub_marks = temp

		for i, j := range temp {
			if j >= 90 {
				j = 10
			} else if j >= 85 && j < 90 {
				j = 9

			} else if j >= 80 && j < 85 {
				j = 8

			} else if j >= 70 && j < 80 {
				j = 7

			} else if j >= 60 && j < 70 {
				j = 6

			} else if j >= 50 && j < 60 {
				j = 5

			} else if j >= 40 && j < 50 {
				j = 4

			} else if j >= 30 && j < 40 {
				j = 3

			} else if j >= 20 && j < 30 {
				j = 2

			} else if j >= 10 && j < 20 {
				j = 1

			} else {
				j = 0
			}
			t = t + j
			esa[i] = esa[i] + (j / float64(len(studentF)))

		}
		tot = tot + (t / float64(len(studentF)))
		s.total = t
		mark_arr = append(mark_arr, t)
		if t >= 45 {
			s.grade = "A+"
		} else if t >= 40 && t < 45 {
			s.grade = "A"
		} else if t >= 35 && t < 40 {
			s.grade = "B+"
		} else if t >= 30 && t < 35 {
			s.grade = "B"
		} else if t >= 25 && t < 30 {
			s.grade = "C+"
		} else if t >= 20 && t < 25 {
			s.grade = "C"
		} else {
			s.grade = "F"
		}
		info[i[1]] = s

	}
	sort.Strings(name_arr)
	sort.Strings(n_id_arr)
	sort.Sort(sort.Reverse(sort.Float64Slice(mark_arr)))

	return info, esa, tot, name_arr, n_id_arr, mark_arr
}

func (CBSEstudents) Print_institute_report(name_arr, n_id_arr []string, mark_arr []float64, info map[string]interface{}, esa map[string]float64, tot float64) {
	// Specify the file name

	fileName := "INSTITUTE REPORT CBSE.csv"

	// Create or open the CSV file for writing
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to the CSV file
	header := []string{"TOTAL STUDENTS:", fmt.Sprintf("%d", len(name_arr))}
	err = writer.Write(header)
	header = []string{"EACH SUBJECT AVERAGE:"}
	err = writer.Write(header)

	// Write data to the CSV file
	for i, j := range esa {
		header = []string{i, ":", fmt.Sprintf("%f", j)}
		err := writer.Write(header)
		if err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}
	header = []string{"TOTAL AVERAGE:", fmt.Sprintf("%f", tot)}
	err = writer.Write(header)
	header = []string{"RANKING:"}
	err = writer.Write(header)

	var id_vise_name []string
	var name_vise_mark []float64
	info3 := info
	for _, j := range n_id_arr {
		id_vise_name = append(id_vise_name, info3[j].(CBSEstudents).name)
	}
	for _, i := range name_arr {

		for _, j := range info3 {
			if i == j.(CBSEstudents).name {
				name_vise_mark = append(name_vise_mark, j.(CBSEstudents).total)
			}
		}
	}
	for i, j := range mark_arr {

		if i < len(mark_arr)-1 {

			if mark_arr[i] != mark_arr[i+1] {
				for l, k := range info3 {
					if j == k.(CBSEstudents).total && k.(CBSEstudents).name != "1" {
						fmt.Println(k.(CBSEstudents).name, l, k.(CBSEstudents).total)
						header = []string{k.(CBSEstudents).name, l, fmt.Sprintf("%f", k.(CBSEstudents).total)}
						err = writer.Write(header)
					}
				}
			} else {
				for k, _ := range name_vise_mark {
					if name_vise_mark[k] == mark_arr[i] {

						for l, m := range info3 {
							if m.(CBSEstudents).name == name_arr[k] && m.(CBSEstudents).total == mark_arr[i] {
								fmt.Println(name_arr[k], l, m.(CBSEstudents).total)
								header = []string{name_arr[k], l, fmt.Sprintf("%f", m.(CBSEstudents).total)}
								err = writer.Write(header)
								te := info3[l].(CBSEstudents)
								te.name = "1"
								info3[l] = te

								break
							}

						}
						name_vise_mark[k] = 0
						break

					}
				}

			}

		} else {

			for l, k := range info3 {
				if j == k.(CBSEstudents).total && k.(CBSEstudents).name != "1" {
					fmt.Println(k.(CBSEstudents).name, l, k.(CBSEstudents).total)
					header = []string{k.(CBSEstudents).name, l, fmt.Sprintf("%f", k.(CBSEstudents).total)}
					err = writer.Write(header)
					break
				}
			}
		}
	}

}
func (SBstudents) Print_institute_report(name_arr, n_id_arr []string, mark_arr []float64, info map[string]interface{}, esa map[string]float64, tot float64) {
	// Specify the file name

	fileName := "INSTITUTE REPORT StateBoard.csv"

	// Create or open the CSV file for writing
	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)
	defer writer.Flush()

	// Write header to the CSV file
	header := []string{"TOTAL STUDENTS:", fmt.Sprintf("%d", len(name_arr))}
	err = writer.Write(header)
	header = []string{"EACH SUBJECT AVERAGE:"}
	err = writer.Write(header)

	// Write data to the CSV file
	for i, j := range esa {
		header = []string{i, ":", fmt.Sprintf("%f", j)}
		err := writer.Write(header)
		if err != nil {
			fmt.Println("Error writing record:", err)
			return
		}
	}
	header = []string{"TOTAL AVERAGE:", fmt.Sprintf("%f", tot)}
	err = writer.Write(header)
	header = []string{"RANKING:"}
	err = writer.Write(header)

	var id_vise_name []string
	var name_vise_mark []float64

	for _, j := range n_id_arr {
		id_vise_name = append(id_vise_name, info[j].(SBstudents).name)
	}
	for _, i := range name_arr {

		for _, j := range info {
			if i == j.(SBstudents).name {
				name_vise_mark = append(name_vise_mark, j.(SBstudents).total)
			}
		}
	}
	for i, j := range mark_arr {

		if i < len(mark_arr)-1 {

			if mark_arr[i] != mark_arr[i+1] {
				for l, k := range info {
					if j == k.(SBstudents).total && k.(SBstudents).name != "1" {
						fmt.Println(k.(SBstudents).name, l, k.(SBstudents).total)
						header = []string{k.(SBstudents).name, l, fmt.Sprintf("%f", k.(SBstudents).total)}
						err = writer.Write(header)
					}
				}
			} else {
				for k, _ := range name_vise_mark {
					if name_vise_mark[k] == mark_arr[i] {

						for l, m := range info {
							if m.(SBstudents).name == name_arr[k] && m.(SBstudents).total == mark_arr[i] {
								fmt.Println(name_arr[k], l, m.(SBstudents).total)
								header = []string{name_arr[k], l, fmt.Sprintf("%f", m.(SBstudents).total)}
								err = writer.Write(header)
								te := info[l].(SBstudents)
								te.name = "1"
								info[l] = te

								break
							}

						}
						name_vise_mark[k] = 0
						break

					}
				}

			}

		} else {

			for l, k := range info {
				if j == k.(SBstudents).total && k.(SBstudents).name != "1" {
					fmt.Println(k.(SBstudents).name, l, k.(SBstudents).total)
					header = []string{k.(SBstudents).name, l, fmt.Sprintf("%f", k.(SBstudents).total)}
					err = writer.Write(header)
					break
				}
			}
		}
	}

}
func (SBstudents) student_report(info map[string]interface{}, si string) {

	for i, j := range info {
		if i == si {
			fmt.Println(j.(SBstudents).name, "\n", j.(SBstudents).phno, "\n", j.(SBstudents).grade, "\n", j.(SBstudents).submarks, "\n", j.(SBstudents).total, "\n")
		}
	}
}
func (CBSEstudents) student_report(info map[string]interface{}, si string) {

	for i, j := range info {
		if i == si {
			fmt.Println(j.(CBSEstudents).name, "\n", j.(CBSEstudents).phno, "\n", j.(CBSEstudents).grade, "\n", j.(CBSEstudents).sub_marks, "\n", j.(CBSEstudents).total, "\n")
		}
	}
}
func main() {
	studentF := createf("student.csv")
	marksF := createf("marks.csv")
	studentF2 := createf("cbseStudents.csv")
	marksF2 := createf("cbseMarks.csv")

	var intr1 institution = SBstudents{}
	var intr2 institution = CBSEstudents{}

	info, esa, total, name_arr1, n_id_arr1, mark_arr1 := intr1.student_info(studentF, marksF)
	fmt.Println(info, esa, total, "\n")

	info2, esa2, total2, name_arr2, n_id_arr2, mark_arr2 := intr2.student_info(studentF2, marksF2)
	fmt.Println(info2, esa2, total2)
	intr1.Print_institute_report(name_arr1, n_id_arr1, mark_arr1, info, esa, total)
	intr2.Print_institute_report(name_arr2, n_id_arr2, mark_arr2, info2, esa2, total2)

	si := "S1"
	intr1.student_report(info, si)

	intr2.student_report(info2, si)
    
}
