package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main () {
	niter := 1000
	rand.Seed (time.Now().UnixNano())
	sum_y := 0.0;

	/* Main loop */
	for iter:=0; iter<niter+1; iter++ {
		x := rand.Float64()
		y := math.Sqrt (1.0 - x*x)
		sum_y += y

		fmt.Printf ("%4d %.10f\n", iter, sum_y/float64(iter))
	}
}
