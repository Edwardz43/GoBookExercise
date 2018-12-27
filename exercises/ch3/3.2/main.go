package main

import (
	"fmt"
	"io"
	"math"
	"os"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 12        // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

type zFunc func(x, y float64) float64

func main() {
	errMsg := "ex3.2 saddle|eggbox"
	if len(os.Args) < 2 {
		fmt.Println(os.Stderr, errMsg)
		os.Exit(1)
	}
	//fmt.Println("http://localhost:8000")
	var f zFunc
	switch os.Args[1] {
	case "saddle":
		f = saddle
	case "eggbox":
		f = eggbox
	case "test":
		f = test
	default:
		fmt.Println(os.Stderr, errMsg)
		os.Exit(1)
	}
	svg(os.Stdout, f)
}
func corner(i, j int, f zFunc) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy

}

func surface(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

func eggbox(x, y float64) float64 {
	return 0.2 * (math.Cos(x) + math.Cos(y))
}

func test(x, y float64) float64 {
	return 0.05 * math.Exp(x*y) / math.Exp(1+x*x)
}

func saddle(x, y float64) float64 {
	a := 25.0
	b := 17.0
	return (y*y/(a*a) - (x*x)/(b*b))
}

func svg(w io.Writer, f zFunc) {
	fmt.Printf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, f)
			bx, by := corner(i, j, f)
			cx, cy := corner(i, j+1, f)
			dx, dy := corner(i+1, j+1, f)
			if math.IsNaN(ax) || math.IsNaN(bx) || math.IsNaN(cx) || math.IsNaN(dx) {
				continue
			}
			// fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
			// 	ax, ay, bx, by, cx, cy, dx, dy)
			fmt.Fprintf(w, "<polygon style='stroke: %s; fill: #222222' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				"#666666", ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}
