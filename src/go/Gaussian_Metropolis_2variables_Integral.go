package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func s_physics (x float64) float64 {
	return (2.0 + math.Tanh(x)) / 3.0
}

func s_baseball (y float64) (s float64) {
	if y <= 2.0 {
		s = 0.0
	} else {
		s = y*y/2.0
	}
	return
}

func calc_action (x, y float64) float64 {
	return 0.5*(x*x + y*y + x*y)
}

func metropolis_test (action_init, action_fin float64) bool {
	metropolis := rand.Float64 ()
	return math.Exp (action_init-action_fin) > metropolis
}

func calc_expectation (niter int, step_size_x, step_size_y float64) {
	const (
		wstep int = 10
		siz int = 3
	)
	fmt.Printf ("%d %d\n", niter, wstep)
	k := siz-1

	xs := new ([siz]float64)
	ys := new ([siz]float64)
	naccept := 0
	nsamples := 0
	expectation_physics := 0.0
	expectation_baseball1 := 0.0

	expectation_baseball2 := new([siz]float64);
	alphas := new ([siz]float64)
	alphas[0] = 0.0
	alphas[1] = 1.5
	alphas[2] = 3.0
	for iter:=1; iter<=niter; iter++ {
		for i:=0; i<=k; i++ {
			action_init := calc_action (xs[i], ys[i]-alphas[i])

			dx := (rand.Float64()-0.5) * step_size_x * 2.0
			dy := (rand.Float64()-0.5) * step_size_y * 2.0
			candidate_x := xs[i] + dx
			candidate_y := ys[i] + dy
			action_fin := calc_action (candidate_x, candidate_y-alphas[i])

			if metropolis_test (action_init, action_fin) {
				/* Accept */
				naccept ++
				xs[i] = candidate_x
				ys[i] = candidate_y
			}
		}

		if iter%wstep == 0 {
			nsamples ++
			expectation_physics += s_physics (xs[0])
			expectation_baseball1 += s_baseball (ys[0])
			for i:=0; i<k; i++ {
				expectation_baseball2[i] += math.Exp (- (calc_action(xs[i],ys[i]-alphas[i+1]) - calc_action(xs[i],ys[i]-alphas[i])))
			}
			expectation_baseball2[k] += math.Exp(-(calc_action(xs[k],ys[k])-calc_action(xs[k],ys[k]-alphas[k]))) * s_baseball (ys[k])
			prod := 1.0
			for i:=0; i<=k; i++ {
				prod *= expectation_baseball2[i] / float64(nsamples)
			}
			fmt.Printf ("%.10f\t%.10f\t%.10f\t%f\n", 
				expectation_physics/float64(nsamples),
				expectation_baseball1/float64(nsamples),
				prod,
				float64(naccept)/float64(iter))
		}
	}
}


func main () {
	const (
		niter int = 1000000
		step_size_x float64 = 0.5
		step_size_y float64 = 0.5
	)
	rand.Seed (time.Now().UnixNano())
	calc_expectation (niter, step_size_x, step_size_y)
}
