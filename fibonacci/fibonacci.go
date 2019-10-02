package fibonacci

type Init struct {
	A, B, limit int
}

func (f *Init) FindNumberOfFibonacci() []int {
	var result = []int{f.A, f.B}

	for {
		c := f.A + f.B
		f.A = f.B
		f.B = c

		if (c > f.limit) {
			break
		}

		result = append(result, c)
	}

	return result
}
