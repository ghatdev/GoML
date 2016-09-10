package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	s := rand.NewSource(time.Now().UnixNano())

	a := rand.New(s).Float64()
	b := rand.New(s).Float64()

	x1 := []float64{0,0,1,1}
  x2 := []float64{0,1,0,1}
	y := []float64{0, 1, 1, 0}

	alpha := 0.01

	for j := 0; j < 10000; j++ {
		for i := 0; i < len(x); i++ {
			e := y[i] - a*x[i] - b

			Ea := e * -x[i]
			Eb := e * -1

			a -= alpha * Ea
			b -= alpha * Eb
		}
	}

	f := a*4 + b

	fmt.Println(f)

}
