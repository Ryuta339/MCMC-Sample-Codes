package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main () {
	const niter int = 1000000
	// const step_size float64 = 0.5
	const min_step float64 = -0.5
	const max_step float64 = 1.0
	const step_size float64 = max_step - min_step
	rand.Seed (time.Now().UnixNano())

	/* Initialization */
	var x float64 = 0.0
	var naccept int = 0 // counter of acceptance

	/* Main loop */
	for iter:=1; iter<niter+1; iter++ {
		backup_x := x
		action_init := 0.5*x*x
		
		dx := step_size*rand.Float64 () + min_step
		x += dx

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
