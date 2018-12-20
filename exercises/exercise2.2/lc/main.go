package main

import (
	"fmt"
	"os"
	"strconv"

	"GoBook/exercises/exercise2.2/lengthconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		l, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "lc : %v\n", err)
			os.Exit(1)
		}

		i := lengthconv.Inch(l)
		m := lengthconv.MilliMeter(l)
		fmt.Printf("%s = %s, %s = %s\n",
			i, lengthconv.IToM(i), m, lengthconv.MToI(m))
	}
}
