package main

import (
	"fmt"
	"math"
)

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
	y := []float64{0, 1, 1, 0}

	alpha := 0.01
	af = 0.1
	ag = 0.2
	bf = 0.4
	bg = 0.3
	cg = 0.12

	for j := 0; j < 1000; j++ {
		for i := 0; i < 4; i++ {
			g := Sigmoid(ag*x1[i] + bg*x2[i] + bg)
			f := Sigmoid(af*g + bf)

			e := y[i] - f
			se := e * e * 0.5
			fmt.Printf("f: %f, g: %f\n", f, g)
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

		fmt.Printf("af: %f, bf: %f\n", af, bf)
		fmt.Printf("ag: %f, bg: %f, cg: %f\n", ag, bg, cg)

		for k := 0; k < 4; k++ {
			g := Sigmoid(ag*x1[k] + bg*x2[k] + bg)
			f := Sigmoid(af*g + bf)
			fmt.Printf("Result: %f\n\n", f)
		}

	}

}

func Sigmoid(input float64) (output float64) {
	output = 1 / (1 + math.Exp(-input))
	return
}
