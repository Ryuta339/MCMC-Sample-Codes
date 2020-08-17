package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	niter int     = 1000000
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


func main () {
	A := new ([ndim][ndim]float64)

	A[0][0] = 1.0; A[1][1] = 2.0; A[2][2] = 2.0;
	A[0][1] = 1.0; A[0][2] = 1.0; A[1][2] = 1.0;
	for idim:=1; idim<ndim; idim ++ {
		for jdim:=0; jdim<idim; jdim++ {
			A[idim][jdim] = A[jdim][idim]
		}
	}

	rand.Seed (time.Now().UnixNano ())
	grand = MarsagliaPolar ()

	x, y, z := 0.0, 0.0, 0.0;

	/*** Main part ***/
	for iter:=0; iter<niter; iter++ {
		var mu, sigma float64

		// Update x
		sigma = 1.0 / math.Sqrt (A[0][0])
		mu = -A[0][1]/A[0][0]*y - A[0][2]/A[0][0]*z
		x = sigma * grand () + mu

		// Update y
		sigma = 1.0 / math.Sqrt (A[1][1])
		mu = -A[1][0]/A[1][1]*x - A[1][2]/A[1][1]*z
		y = sigma * grand () + mu

		// Update z
		sigma = 1.0 / math.Sqrt (A[2][2])
		mu = -A[2][0]/A[1][1]*x - A[2][1]/A[2][2]*y
		z = sigma * grand () + mu

		if (iter+1)%10==0 {
			fmt.Printf ("%.6f\t%.6f\t%.6f\n", x, y, z)
		}
	}
}
