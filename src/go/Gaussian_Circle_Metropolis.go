package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func calcAction (x float64) float64 {
	if x >= 0 {
		// Gaussian
		return 0.5*x*x + 0.5*math.Log (2.0*math.Pi)
	} else if x >= -1 {
		// Circle
		return - math.Log (2.0*math.Sqrt(1 - x*x)/math.Pi)
	}
	// returns infty in the case x < -1
	return math.Inf (1)
}

func metropolisCheck (action_init, action_fin float64) bool {
	if math.IsInf (action_fin, 1) {
		// the case x < -1 is rejected
		return false
	}
	
	metropolis := rand.Float64 ()
	return math.Exp (action_init - action_fin) > metropolis
}

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
		action_init := calcAction (x)
		
		dx := rand.Float64 ()
		x += (dx-0.5) * step_size * 2.0

		action_fin := calcAction (x)

		/* Metropolis test */
		if metropolisCheck (action_init, action_fin) {
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
