package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main () {
	const niter int = 1000
	rand.Seed (time.Now().UnixNano())

	sum_z := 0.0
	n_in := 0		// counter

	/* Main loop */
	for iter:=0; iter<niter+1; iter++ {
		x := rand.Float64()
		y := rand.Float64()
		rr := x*x + y*y
		if (rr < 1.0) {
			n_in += 1
			sum_z += math.Sqrt (1.0 - rr)
		}
		// Print the expectation value
		fmt.Printf ("%4d %.10f\n", iter, sum_z/float64(n_in)*2.0*math.Pi)
	}
}
