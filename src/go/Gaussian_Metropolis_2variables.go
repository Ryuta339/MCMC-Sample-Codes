package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main () {
	const(
		niter int = 10000
		step_size_x float64 = 0.5
		step_size_y float64 = 0.5
	)
	rand.Seed (time.Now().UnixNano())

	/* Initialization */
	x := 0.0
	y := 0.0
	naccept := 0

	/* Main loop */
	for iter:=1; iter<niter+1; iter++ {
		backup_x := x
		backup_y := y
		action_init := 0.5*(x*x + y*y + x*y)

		dx := (rand.Float64() - 0.5)*step_size_x*2.0
		dy := (rand.Float64() - 0.5)*step_size_y*2.0
		x += dx
		y += dy
		action_fin := 0.5*(x*x + y*y + x*y)

		/* Metropolis test */
		metropolis := rand.Float64()
		if math.Exp(action_init-action_fin) > metropolis {
			// Accept
			naccept ++
		} else {
			// Reject
			x = backup_x
			y = backup_y
		}

		if iter%10 == 0 {
			fmt.Printf ("%.10f\t%.10f\t%f\n", x, y, float64(naccept)/float64(iter))
		}
	}
}

