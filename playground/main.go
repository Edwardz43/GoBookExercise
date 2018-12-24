package main

import "fmt"

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

	i := 0      // untyped integer; implicit int(0)
	r := '\000' // untyped rune; implicit rune('\000')
	f := 0.0    // untyped floating-point; implicit float64(0.0)
	c := 0i     // untyped complex; implicit complex128(0i)

	fmt.Println(i, r, f, c)
}
