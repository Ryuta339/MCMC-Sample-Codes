package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	niter int = 10000		// set the number of samples
	ntau int = 40			// the step number of the leap flog method
	dtau float64 = 1.0		// the step size of the leap flog method
)

var grand func() float64

/* Marsaglia Poler Method */
func MarsagliaPolar () func () float64 {
	var spare float64
	var hasSpare bool = false
	return func () float64 {
		if (hasSpare) {
			hasSpare = false
			return spare
		} else {
			var u, v, s float64
			s = 0.0
			for s>=1.0 || s==0.0 {
				u = rand.Float64 () * 2.0 - 1.0
				v = rand.Float64 () * 2.0 - 1.0
				s = u*u + v*v
			}
			s = math.Sqrt (-2.0*math.Log(s) / s)
			spare = v * s
			hasSpare = true
			return u * s
		}
	}
}

/* Box Muller Method */
func BoxMuller () func() float64 {
	return func () float64 {
		r := rand.Float64 ()
		s := rand.Float64 ()
		return math.Sqrt (-2.0*math.Log(r)) * math.Sin (2.0*math.Pi*s)
	}
}

/* Calculation the action S[x] */
func calc_action (x float64) float64 {
	action := 0.5*x*x
	return action
}

/* Calculation the Hamiltonian H[x, p] */
func calc_hamiltonian (x, p float64) float64 {
	ham := calc_action (x)	// calculate the action (potential energy)
	ham += 0.5*p*p			// add kinetic energy
	return ham
}

/* Calculation dH/dx */
func calc_delh (x float64) float64 {
	// dH(x)/dx ~ (H(x+e)-H(x-e))/(2e)
	const epsilon = 5e-9
	delh := (calc_action (x+epsilon) - calc_action (x-epsilon)) / (2*epsilon)

	return delh
}

/* Time evolution by molecular dynamics */
func Molecular_Dynamics (x, ham_init, ham_fin *float64) {
//	var r1, r2 float64
	r1 := grand ()
//	r2 = grand ()
	p := r1			// generate the momentun p as a Gaussian random variable

	// calculate Hamiltonian
	*ham_init = calc_hamiltonian (*x, p)
	// first step of the leap flog method
	*x += p*0.5*dtau
	// from second to Ntau step of the leap flog method
	for iter:=1; iter<ntau; iter++ {
		delh := calc_delh (*x)
		p -= delh*dtau
		*x += p*dtau
	}

	// final step of the leap flog method
	delh := calc_delh (*x)
	p -= delh*dtau
	*x += p*0.5*dtau

	// recalculate Hamiltonian
	*ham_fin = calc_hamiltonian (*x, p)
}

func main () {
	rand.Seed (time.Now().UnixNano())
	// grand = MarsagliaPolar ()
	grand = BoxMuller ()

	/* initial configuration */
	x := 0.0
	
	var naccept int = 0
	var sum_xx float64 = 0.0
	
	/* main loop */
	for iter:=0; iter<niter; iter++ {
		backup_x := x
		var ham_init, ham_fin float64
		Molecular_Dynamics (&x, &ham_init, &ham_fin)
		metropolis := rand.Float64()
		if (math.Exp (ham_init-ham_fin) > metropolis) {
			// accept
			naccept ++
		} else {
			// reject
			x = backup_x
		}

		/* output data */
		sum_xx += x*x
		fmt.Printf ("%.6f\t%.6f\t%.6f\n",
			x,
			sum_xx / float64(iter+1),
			float64(naccept) / float64(iter+1),
		)
	}
}
