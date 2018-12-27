package rev

// Reverse reverses the int array
func Reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// LeftShift shift the int array left for n steps
func LeftShift(s []int, n int) {

	l := len(s)

	n = n % l

	if n == 0 {
		return
	}

	Reverse(s[:n])
	Reverse(s[n:])
	Reverse(s[:])
}
