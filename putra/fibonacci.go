package putra

type InitFibonacci struct {
	A, B, Limit int
}

func (f *InitFibonacci) FindNumberOfFibonacci() []int {
	var result = []int{f.A, f.B}

	for {
		c := f.A + f.B
		f.A = f.B
		f.B = c

		if c > f.Limit {
			break
		}

		result = append(result, c)
	}

	return result
}
