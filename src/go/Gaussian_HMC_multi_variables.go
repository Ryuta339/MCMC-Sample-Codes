package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	niter int     = 10000
	ntau  int     = 20
	dtau  float64 = 0.5
	ndim  int     = 3
)

var grand func () float64

/*** Gaussian Random Number Generator with Marsaglia Polar Method ***/
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

/*** Gaussian Random Number Generator with Box Muller Algorithm ***/
func BoxMuller () func() float64 {
	return func () float64 {
		r := rand.Float64 ()
		s := rand.Float64 ()
		return math.Sqrt (-2.0*math.Log(r)) * math.Sin (2.0*math.Pi*s)
	}
}

/*** Calculation of the action ***/
func calc_action (
	x *[ndim]float64,
	A *[ndim][ndim]float64,
) float64 {
	action := 0.0
	for idim:=0; idim<ndim; idim++ {
		for jdim:=0; jdim<idim; jdim++ {
			action += x[idim]*A[idim][jdim]*x[jdim]
		}
		action += 0.5*x[idim]*A[idim][idim]*x[idim]
	}

	return action
}

/*** Calculation of the Hamiltonian ***/
func calc_hamiltonian (
	x *[ndim]float64,
	p *[ndim]float64,
	A *[ndim][ndim]float64,
) float64 {
	ham := calc_action (x, A)
	for idim:=0; idim<ndim; idim++ {
		ham += 0.5 * p[idim] * p[idim]
	}

	return ham
}

/*** Calculation of dH/dx ***/
func calc_delh (
	x *[ndim]float64,
	A *[ndim][ndim]float64,
	delh *[ndim]float64,
) {
	const epsilon = 5e-9

	p := new ([ndim]float64)

	for idim:=0; idim<ndim; idim++ {
		delh[idim] = 0.0
	}

	for idim:=0; idim<ndim; idim++ {
		var hp, hm, backup float64

		backup = x[idim]
		x[idim] = backup + epsilon
		hp = calc_hamiltonian (x, p, A)
		x[idim] = backup - epsilon
		hm = calc_hamiltonian (x, p, A)

		delh[idim] = (hp-hm) / (2*epsilon)
	}
}

/*** Molecular evolution ***/
func Molecular_Dynamics (
	x *[ndim]float64,
	A *[ndim][ndim]float64,
) (ham_init float64, ham_fin float64) {
// var r1, r2 float64
	p := new ([ndim]float64)
	delh := new ([ndim]float64)

	for idim:=0; idim<ndim; idim++ {
		r1 := grand()
		p[idim] = r1
	}

	// Calculate Hamiltonian
	ham_init = calc_hamiltonian (x, p, A)
	// First step of leap flog
	for idim:=0; idim<ndim; idim++ {
		x[idim] += p[idim] * 0.5 * dtau
	}

	// Second, ..., Ntau-th steps 
	for step:=1; step<ntau; step++ {
		calc_delh (x, A, delh)
		for idim:=0; idim<ndim; idim++ {
			p[idim] -= delh[idim] * dtau
			x[idim] += p[idim] * dtau
		}
	}
	// Last step of leap flog
	calc_delh (x, A, delh)
	for idim:=0; idim<ndim; idim++ {
		p[idim] -= delh[idim]*dtau
		x[idim] += p[idim] * 0.5 * dtau
	}
	// Calculate Hamiltonian again
	ham_fin = calc_hamiltonian (x, p, A)

	return
}


func main () {
	x := new ([ndim]float64)
	A := new ([ndim][ndim]float64)

	A[0][0] = 1.0; A[1][1] = 2.0; A[2][2] = 2.0;
	A[0][1] = 1.0; A[0][2] = 1.0; A[1][2] = 1.0;
	for idim:=1; idim<ndim; idim ++ {
		for jdim:=0; jdim<idim; jdim++ {
			A[idim][jdim] = A[jdim][idim]
		}
	}

	rand.Seed (time.Now().UnixNano())
	grand = MarsagliaPolar ()

	/*** Main part ***/
	naccept := 0
	for iter:=0; iter<niter; iter++ {
		backup_x := new ([ndim]float64)
		for idim:=0; idim<ndim; idim++ {
			backup_x[idim] = x[idim]
		}
		ham_init, ham_fin := Molecular_Dynamics (x, A)

		metropolis := rand.Float64 ()
		if (math.Exp (ham_init-ham_fin) > metropolis) {
			// accept
			naccept ++
		} else {
			// reject
			for idim:=0; idim<ndim; idim++ {
				x[idim] = backup_x[idim]
			}
		}

		/*** Data output ***/
		if (iter+1)%10 == 0 {
			fmt.Printf ("%.6f\t%.6f\t%.6f\t%.6f\n",
				x[0],
				x[1],
				x[2],
				float64(naccept)/float64(iter+1))
		}
	}
}

