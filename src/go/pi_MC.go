package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	niter := 1000
	rand.Seed (time.Now().UnixNano())
	n_in := 0 // Initialize the counter

	/* Main loop */
	for iter:=0; iter<niter+1; iter++ {
		x := rand.Float64()
		y := rand.Float64()

		if x*x + y*y < 1.0 {
			n_in += 1
		}

		fmt.Printf("%4d %.10f\n", iter, float64(n_in)/float64(iter))
	}
}
