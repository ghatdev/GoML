package main

import "fmt"

type Any interface{}

type Variable struct {
	Value Any
}

func main() {
	var af, ag float64
	var bf, bg float64
	var cg float64

	x1 := []float64{0, 0, 1, 1}
	x2 := []float64{0, 1, 0, 1}
	y := []float64{0, 1, 1, 1}

	alpha := 0.01

	for j := 0; j < 10000; j++ {
		for i := 0; i < len(x1); i++ {
			g := ag*x1[i] + bg*x2[i] + cg
			f := af*g + bf

			e := y[i] - f
			se := e * e * 0.5
			fmt.Printf("Squared Error: %f\n", se)

			Eaf := -e * g
			Ebf := -e * 1
			Eag := -e * af * x1[i]
			Ebg := -e * af * x2[i]
			Ecg := -e * af * 1

			af -= alpha * Eaf
			bf -= alpha * Ebf
			ag -= alpha * Eag
			bg -= alpha * Ebg
			cg -= alpha * Ecg
		}
		g := ag*0 + bg*0 + cg
		f := af*g + bf

		fmt.Printf("af: %f, bf: %f\n", af, bf)
		fmt.Printf("ag: %f, bg: %f, cg: %f\n", ag, bg, cg)
		fmt.Printf("Result: %f\n\n", f)
	}

}
