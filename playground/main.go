package main

import (
	"fmt"
	"gobook/book/ch3/basename"
	"gobook/book/ch3/comma"
	"math"
)

func main() {
	fmt.Printf("The max int 8 is %d \n", math.MaxInt8)

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

	fmt.Println(basename.Basename1("a/b/c.go")) // "c"
	fmt.Println(basename.Basename1("c.d.go"))   // "c.d"
	fmt.Println(basename.Basename1("abc"))      // "abc"

	fmt.Println("-----------------")

	fmt.Println(basename.Basename2("a/b/c.go")) // "c"
	fmt.Println(basename.Basename2("c.d.go"))   // "c.d"
	fmt.Println(basename.Basename2("abc"))      // "abc"

	fmt.Println(comma.Comma("123456789"))
	fmt.Println(comma.Comma("12"))
}
