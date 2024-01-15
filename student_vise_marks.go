package gofunctions

func Student_vise_marks(m []Marks, s []Students) map[string]float64 {
	sm := map[string]float64{}

	for _, j := range s {
		for _, k := range m {
			if k.N_id == j.N_id {

				sm[j.Name] = sm[j.Name] + k.Mark
			}

		}

	}

	return sm
}
