package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

const (
	nbeta int = 2000
	niter = 1000000
	nskip = 10000
	stepSize float64 = 0.1
	dbeta float64 = 0.5
)

func calcF (x float64) float64 {
	return (x-1.0)*(x-1.0)*((x+1.0)*(x+1.0)+0.01)
}

func main () {
	naccept := make ([]int, nbeta)
	x := make ([]float64, nbeta)
	beta := make ([]float64, nbeta)

	rand.Seed (time.Now().UnixNano ())

	/* Set the initial configuration */
	for ibeta:=0; ibeta<nbeta; ibeta++ {
		x[ibeta] = 0.0
		beta[ibeta] = float64(ibeta+1)*dbeta
		naccept[ibeta] = 0.0
	}

	/* Main loop */
	for iter:=0; iter<niter; iter++ {
		for ibeta:=0; ibeta<nbeta; ibeta++ {
			backupX := x[ibeta]
			actionInit := calcF (x[ibeta]) * beta[ibeta]

			dx := (rand.Float64()-0.5) * stepSize * 2.0
			x[ibeta] += dx

			actionFin := calcF (x[ibeta]) * beta[ibeta]

			/* Metropolis test */
			metropolis := rand.Float64 ()
			if math.Exp (actionInit - actionFin) > metropolis {
				/* accept */
				naccept[ibeta] += 1
			} else {
				/* reject */
				x[ibeta] = backupX
			}
		}

		for ibeta:=0; ibeta<nbeta-1; ibeta ++ {
			tmp1 := calcF(x[ibeta]); tmp2 := calcF(x[ibeta+1])
			actionInit := tmp1*beta[ibeta] + tmp2*beta[ibeta+1]
			actionFin := tmp1*beta[ibeta+1] + tmp2*beta[ibeta]

			/* Metropolis test */
			metropolis := rand.Float64 ()
			if math.Exp (actionInit - actionFin) > metropolis {
				/* accept = exchange */
				backupX := x[ibeta]
				x[ibeta] = x[ibeta+1]
				x[ibeta+1] = backupX
			}
		}
		/* data output */
		fmt.Printf ("%f\t%f\t%f\n", x[19], x[199], x[1999])
	}

}
