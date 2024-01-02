package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"sort"
	"strconv"
)

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
func student_info(studentF, marksF [][]string) (map[string]CBSEstudents, map[string]float64, float64, []string, []string, []float64) {
	info := map[string]CBSEstudents{}
	esa := map[string]float64{}
	tot := float64(0)
	t := float64(0)
	var s CBSEstudents
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
func Print_institute_report(name_arr, n_id_arr []string, mark_arr []float64, info map[string]CBSEstudents, esa map[string]float64, tot float64) {
	// Specify the file name
	fileName := "INSTITUTE REPORT.csv"

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
		id_vise_name = append(id_vise_name, info[j].name)
	}
	for _, i := range name_arr {

		for _, j := range info {
			if i == j.name {
				name_vise_mark = append(name_vise_mark, j.total)
			}
		}
	}
	for i, j := range mark_arr {

		if i < len(mark_arr)-1 {

			if mark_arr[i] != mark_arr[i+1] {
				for l, k := range info {
					if j == k.total && k.name != "1" {
						fmt.Println(k.name, l, k.total)
						header = []string{k.name, l, fmt.Sprintf("%f", k.total)}
						err = writer.Write(header)
					}
				}
			} else {
				for k, _ := range name_vise_mark {
					if name_vise_mark[k] == mark_arr[i] {

						for l, m := range info {
							if m.name == name_arr[k] && m.total == mark_arr[i] {
								fmt.Println(name_arr[k], l, m.total)
								header = []string{name_arr[k], l, fmt.Sprintf("%f", m.total)}
								err = writer.Write(header)
								te := info[l]
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
				if j == k.total && k.name != "1" {
					fmt.Println(k.name, l, k.total)
					header = []string{k.name, l, fmt.Sprintf("%f", k.total)}
					err = writer.Write(header)
					break
				}
			}
		}
	}

}
func main() {

	studentF := createf("cbseStudents.csv")
	marksF := createf("cbseMarks.csv")
	info, esa, tot, name_arr, n_id_arr, mark_arr := student_info(studentF, marksF)
	fmt.Println(info, "\n", esa, "\n", tot)
	Print_institute_report(name_arr, n_id_arr, mark_arr, info, esa, tot)

}
