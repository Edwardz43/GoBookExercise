package main

func main() {
	//fmt.Printf("The max int 8 is %d \n", math.MaxInt8)

	// var x uint8 = 1<<1 | 1<<5
	// var y uint8 = 1<<1 | 1<<2
	// fmt.Printf("%08b\n", x)    // "00100010", the set {1, 5}
	// fmt.Printf("%08b\n", y)    // "00000110", the set {1, 2}
	// fmt.Printf("%08b\n", x&y)  // "00000010", the intersection {1}
	// fmt.Printf("%08b\n", x|y)  // "00100110", the union {1, 2, 5}
	// fmt.Printf("%08b\n", x^y)  // "00100100", the symmetric difference {2, 5}
	// fmt.Printf("%08b\n", x&^y) // "00100000", the difference {5}
	// for i := uint(0); i < 8; i++ {
	// 	if x&(1<<i) != 0 { // membership test
	// 		fmt.Println(i) // "1", "5"
	// 	}
	// }
	// fmt.Printf("%08b\n", x<<1) // "01000100", the set {2, 6}
	// fmt.Printf("%08b\n", x>>1) // "00010001", the set {0, 4}

	// medals := []string{"gold", "silver", "bronze"}
	// for i := len(medals) - 1; i >= 0; i-- {
	// 	fmt.Println(medals[i]) // "bronze", "silver", "gold"
	// }

	// o := 0666
	// fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"
	// x := int64(0xdeadbeef)
	// fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)

	// for x := 0; x < 8; x++ {
	// 	fmt.Printf("x = %d e^x = %8.3f\n", x, math.Exp(float64(x)))
	// }

	// var z float64
	// fmt.Println(z, -z, 1/z, -1/z, z/z) // "0 -0 +Inf -Inf NaN"

	// nan := math.NaN()
	// fmt.Println(nan == nan, nan < nan, nan > nan) // "false false false"

	// var x = complex(1, 2)    // 1+2i
	// var y = complex(3, 4)    // 3+4i
	// fmt.Println(x * y)       // "(-5+10i)"
	// fmt.Println(real(x * y)) // "-5"
	// fmt.Println(imag(x * y)) // "10"

	// fmt.Println(3.14i) // "10"
	// fmt.Println(time.Now().Format("2006-01-02 15:04:05"))

	// x := 123
	// y := fmt.Sprintf("%d", x)
	// fmt.Println(y, strconv.Itoa(x)) // "123 123"

	// fmt.Println(strconv.FormatInt(int64(x), 2)) // "1111011"
	// s := fmt.Sprintf("x=%b", x)                 // "x=1111011"
	// fmt.Println(s)

	// x, err := strconv.Atoi("123") // x is an int
	// if err == nil {
	// 	fmt.Println(x)
	// }
	// y, err := strconv.ParseInt("123", 10, 64) // base 10, up to 64 bits
	// if err == nil {
	// 	fmt.Println(y)
	// }

	// const noDelay time.Duration = 0
	// const timeout = 5 * time.Minute
	// fmt.Printf("%T %[1]v\n", noDelay)     // "time.Duration 0"
	// fmt.Printf("%T %[1]v\n", timeout)     // "time.Duration 5m0s"
	// fmt.Printf("%T %[1]v\n", time.Minute) // "time.Duration 1m0s"

	// const (
	// 	a = 1
	// 	b
	// 	c = 2
	// 	d
	// )
	// fmt.Println(a, b, c, d) // "1 1 2 2"

	// type Weekday int
	// const (
	// 	Sunday Weekday = iota
	// 	Monday
	// 	Tuesday
	// 	Wednesday
	// 	Thursday
	// 	Friday
	// 	Saturday
	// )
	// fmt.Println(Friday)

	// var x float32 = math.Pi
	// var y float64 = math.Pi
	// var z complex128 = math.Pi
	// fmt.Println(x, y, z)

	// const Pi64 float64 = math.Pi
	// var x1 float32 = float32(Pi64)
	// var y1 float64 = Pi64
	// var z1 complex128 = complex128(Pi64)

	// fmt.Println(x1, y1, z1)

	// var f float64 = 212
	// fmt.Println((f - 32) * 5 / 9)     // "100"; (f - 32) * 5 is a float64
	// fmt.Println(5 / 9 * (f - 32))     // "0"; 5/9 is an untyped integer, 0
	// fmt.Println(5.0 / 9.0 * (f - 32)) // "100"; 5.0/9.0 is an untyped float

	// var f float64 = 3 + 0i // untyped complex -> float64
	// fmt.Printf("%T %v\n", f, f)
	// f = 2 // untyped integer -> float64
	// fmt.Printf("%T %v\n", f, f)
	// f = 1e123 // untyped floating-point -> float64
	// fmt.Printf("%T %v\n", f, f)
	// f = 'a' // untyped rune -> float64
	// fmt.Printf("%T %v\n", f, f)

	// const (
	// 	deadbeef = 0xdeadbeef        // untyped int with value 3735928559
	// 	a        = uint32(deadbeef)  // uint32 with value 3735928559
	// 	b        = float32(deadbeef) // float32 with value 3735928576 (rounded up)
	// 	c        = float64(deadbeef) // float64 with value 3735928559 (exact)
	// 	d        = int32(deadbeef)   // compile error: constant overflows int32
	// 	e        = float64(1e309)    // compile error: constant overflows float64
	// 	f        = uint(-1)          // compile error: constant underflows uint
	// )

	// i := 0      // untyped integer; implicit int(0)
	// r := '\000' // untyped rune; implicit rune('\000')
	// f := 0.0    // untyped floating-point; implicit float64(0.0)
	// c := 0i     // untyped complex; implicit complex128(0i)

	// fmt.Println(i, r, f, c)
	// s := "Hello, 世界"
	// fmt.Println(len(s))                    // "13"
	// fmt.Println(utf8.RuneCountInString(s)) // "9"

	// // for i := 0; i < len(s); {
	// // 	r, size := utf8.DecodeRuneInString(s[i:])
	// // 	fmt.Printf("%d\t%c\n", i, r)
	// // 	i += size
	// // }

	// for i, r := range "Hello, 世界" {
	// 	fmt.Printf("%d\t%q\t%d\n", i, r, r)
	// }

	// fmt.Println(utf8.RuneCountInString(s))

	// "program" in Japanese katakana
	// s := "プログラム"
	// fmt.Printf("% x\n", s) // "e3 83 97 e3 83 ad e3 82 b0 e3 83 a9 e3 83 a0"
	// r := []rune(s)
	// fmt.Printf("%x\n", r) // "[30d7 30ed 30b0 30e9 30e0]"

	// fmt.Println(string(1234567)) // "(?)"

	// fmt.Println(basename.Basename1("a/b/c.go")) // "c"
	// fmt.Println(basename.Basename1("c.d.go"))   // "c.d"
	// fmt.Println(basename.Basename1("abc"))      // "abc"

	// fmt.Println("-----------------")

	// fmt.Println(basename.Basename2("a/b/c.go")) // "c"
	// fmt.Println(basename.Basename2("c.d.go"))   // "c.d"
	// fmt.Println(basename.Basename2("abc"))      // "abc"

	// fmt.Println(comma.Comma("123456789"))
	// fmt.Println(comma.Comma("12"))

	// var a [3]int             // array of 3 integers
	// fmt.Println(a[0])        // print the first element
	// fmt.Println(a[len(a)-1]) // print the last element, a[2]
	// // Print the indices and elements.
	// for i, v := range a {
	// 	fmt.Printf("%d %d\n", i, v)
	// }
	// // Print the elements only.
	// for _, v := range a {
	// 	fmt.Printf("%d\n", v)
	// }

	//var q [3]int = [3]int{1, 2, 3}
	// var r [3]int = [3]int{1, 2}
	// fmt.Println(r[2]) // "0"

	// q := [...]int{1, 2, 3}
	// fmt.Printf("%T\n", q) // "[3]int"

	// q := [3]int{1, 2, 3}
	// q = [4]int{1, 2, 3, 4} // compile error: cannot assign [4]int to [3]int

	// type Currency int
	// const (
	// 	USD Currency = iota // 美元
	// 	EUR                 // 歐元
	// 	GBP                 // 英鎊
	// 	RMB                 // 人民幣
	// 	NTD                 // 新台幣
	// )
	// symbol := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥", NTD: "$"}
	// fmt.Println(GBP, symbol[GBP]) // "3 ¥"

	// r := [...]int{99: -1}
	// fmt.Println(r)

	// a := [2]int{1, 2}
	// b := [...]int{1, 2}
	// c := [2]int{1, 3}
	// fmt.Println(a == b, a == c, b == c) // "true false false"
	// d := [3]int{1, 2}
	// fmt.Println(a == d) // compile error: cannot compare [2]int == [3]int

	// months := [...]string{
	// 	1: "January", 2: "February", 3: "March", 4: "April", 5: "May", 6: "June",
	// 	7: "July", 8: "August", 9: "September", 10: "October", 11: "November", 12: "December"}

	// q2 := months[4:7]
	// summer := months[6:9]

	// fmt.Printf("%s\n", q2)
	// fmt.Printf("%s\n", summer)

	// for _, s := range summer {
	// 	for _, q := range q2 {
	// 		if s == q {
	// 			fmt.Printf("%s appears in both\n", s)
	// 		}
	// 	}
	// }

	// //fmt.Println(summer[:20])    // panic: out of range
	// endlessSummer := summer[:5] // extend a slice (within capacity)
	// fmt.Println(endlessSummer)  // "[June July August September October]"

	// s := [...]int{0, 1, 2, 3, 4, 5}

	// rev.LeftShift(s[:], 6)

	// fmt.Println(s) // "[2 3 4 5 0 1]"

	// var s []int // len(s) == 0, s == nil
	// fmt.Println(s)
	// s = nil // len(s) == 0, s == nil
	// fmt.Println(s == nil)
	// s = []int(nil) // len(s) == 0, s == nil
	// fmt.Println(s == nil)
	// s = []int{} // len(s) == 0, s != nil
	// fmt.Println(s == nil)

	// var runes []rune
	// for _, r := range "Hello, 世界" {
	// 	runes = append(runes, r)
	// }
	// fmt.Printf("%q\n", runes) // "['H' 'e' 'l' 'l' 'o' ',' ' ' '世' '界']"

	// a := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	// fmt.Printf("%T %v\n", a, a)

	// s := a[1:5]
	// fmt.Printf("%T %v\n", s, s)

	// s = append.Int(s, 10, 20, 30)
	// fmt.Printf("%T %v\n", s, s)
	// fmt.Printf("%T %v\n", a, a)

	// var x, y []int
	// for i := 0; i < 10; i++ {
	// 	y = append.Int(x, i)
	// 	fmt.Printf("%d cap=%d\t%v\n", i, cap(y), y)
	// 	x = y
	// }

	// var x []int
	// x = append(x, 1)
	// x = append(x, 2, 3)
	// x = append(x, 4, 5, 6)
	// x = append(x, x...) // append the slice x
	// fmt.Println(x)      // "[1 2 3 4 5 6 1 2 3 4 5 6]"
	// data := []string{"one", "", "three"}
	// fmt.Printf("%q\n", nonempty.NonemptyV2(data)) // `["one" "three"]`
	// fmt.Printf("%q\n", data)                      // `["one" "three" "three"]`

	// ages := map[string]int{
	// 	"alice":   31,
	// 	"charlie": 34,
	// 	"ed":      36,
	// 	"ting":    31,
	// }

	// var names []string
	// for name := range ages {
	// 	names = append(names, name)
	// }
	// sort.Strings(names)
	// for _, name := range names {
	// 	fmt.Printf("%s\t%d\n", name, ages[name])
	// }

	// var ages map[string]int
	// fmt.Println(ages == nil)    // "true"
	// fmt.Println(len(ages) == 0) // "true"

	// //ages["carol"] = 21 // panic: assignment to entry in nil map

	// age, ok := ages["bob"]
	// if !ok { /* "bob" is not a key in this map; age == 0. */
	// 	fmt.Println(ok)
	// } else {
	// 	fmt.Println(age)
	// }

	// if age, ok := ages["bob"]; !ok { /* ... */
	// }

}
