package sum

func Sum(vals ...int) int {
	total := 0
	for _, v := range vals {
		total += v
	}
	return total

}
