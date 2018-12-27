package nonempty

// Nonempty returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func Nonempty(s []string) []string {
	i := 0
	for _, v := range s {
		if v != "" {
			s[i] = v
			i++
		}
	}
	return s[:i]
}

// NonemptyV2 returns a slice holding only the non-empty strings.
// The underlying array is modified during the call.
func NonemptyV2(s []string) []string {
	out := s[:0]
	for _, v := range s {
		if v != "" {
			out = append(out, v)
		}
	}
	return out
}
