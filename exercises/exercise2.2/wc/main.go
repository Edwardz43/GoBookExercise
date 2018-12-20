package main

import (
	"fmt"
	"os"
	"strconv"

	"GoBook/exercises/exercise2.2/weightconv"
)

func main() {
	for _, arg := range os.Args[1:] {
		w, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "wc : %v\n", err)
			os.Exit(1)
		}

		k := weightconv.KiloGram(w)
		p := weightconv.Pound(w)
		fmt.Printf("%s = %s, %s = %s\n",
			p, weightconv.IToM(p), k, weightconv.MToI(k))
	}
}
