package main

import (
	"fmt"
	"math"
	"os"

	plot "gonum.org/v1/plot"
	plotter "gonum.org/v1/plot/plotter"
)

const ITERATIONS = 80

func calculate(rx float64, ix float64, cx float64, cy float64) (float64, float64) {
	a := rx * rx
	b := 2 * rx * ix
	c := ix * ix * (-1)

	real := a + c + cx
	img := b + cy

	return real, img
}

func mandlebrot(rx float64, ix float64, cx float64, cy float64, counter int64) (float64, float64, float64) {
	real := rx
	img := ix
	i := 0
	for i < ITERATIONS {
		real, img = calculate(real, img, cx, cy)
		if real == math.Inf(1) || img == math.Inf(1) || real == math.Inf(-1) || img == math.Inf(-1) {
			return real, img, float64(i)
		}
		i = i + 1
	}
	return real, img, ITERATIONS
}

func representation(points plotter.XYZs) {
	represent := plot.New()

	f, err2 := os.Create("mandlebrot.png")
	if err2 != nil {
		panic("Png was not created")
	}

	s, errPl := plotter.NewScatter(points)

	if errPl != nil {
		panic("Did not scatter")
	}

	represent.Add(s)

	wt, err1 := represent.WriterTo(1024, 1024, "png")
	if err1 != nil {
		panic("Error creating window")
	}

	_, err3 := wt.WriteTo(f)
	if err3 != nil {
		panic("Did not write to file")
	}

	if error := f.Close(); error != nil {
		panic("Did not close")
	}
}

func main() {
	fmt.Println("Mandlebrot set is starting!")

	cx := -2.20
	cy := -2.20
	var points plotter.XYZs

	for cx < 2.20 {
		cy = -2.00
		for cy < 2.20 {
			var x, y, iterations = mandlebrot(0, 0, cx, cy, 0)
			if y != math.Inf(1) && x != math.Inf(1) && y != math.Inf(-1) && x != math.Inf(-1) {
				points = append(points, struct{ X, Y, Z float64 }{cx, cy, iterations})
			}
			cy = cy + 0.002
		}

		cx = cx + 0.002
	}

	representation(points)
}
