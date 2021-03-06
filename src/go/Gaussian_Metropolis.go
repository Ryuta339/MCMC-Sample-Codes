package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main () {
	const niter int = 1000000
	const step_size float64 = 0.5
	rand.Seed (time.Now().UnixNano())

	/* Initialization */
	var x float64 = 0.0
	var naccept int = 0 // counter of acceptance

	/* Main loop */
	for iter:=1; iter<niter+1; iter++ {
		backup_x := x
		action_init := 0.5*x*x
		
		dx := rand.Float64 ()
		x += (dx-0.5) * step_size * 2.0

		action_fin := 0.5*x*x

		/* Metropolis test */
		metropolis := rand.Float64 ()
		if math.Exp (action_init-action_fin) > metropolis {
			// Accept
			naccept ++
		} else {
			// Reject
			x = backup_x
		}
		/* Output the result */
		fmt.Printf ("%.10f\t%f\n", x, float64(naccept)/float64(iter))
	}
}
