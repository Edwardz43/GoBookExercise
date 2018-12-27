package append

//Int append push y into x
func Int(x []int, y ...int) []int {
	var z []int
	zlen := len(x) + len(y)
	if zlen <= cap(x) {
		z = x[:zlen]
		//z[len(x):] = y
	} else {
		zcap := zlen
		if zcap < 2*len(x) {
			zcap = 2 * len(x)
		}
		z = make([]int, zlen, zcap)
		//copy(z[len(x):], y)
	}
	copy(z[len(x):], y)
	return z
}
