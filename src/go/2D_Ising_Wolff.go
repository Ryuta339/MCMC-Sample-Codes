/*** 2d Ising model with Wolff algorithm ***/
package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	niter = 10000
	nx int = 64			// number of sites along x-direction
	ny int = 64			// number of sites along y-direction
	coupling_J float64 = 1.0
	coupling_h float64 = 0.1
	temperature float64 = 5.0
	nskip = 100		// Frequency of measurement
	nconfig = 10000	// Frequency of saving config; 0->don't save
)

/*** Calculation of the action ***/
func calcAction (
	spin [][]int,
	coupling_J,
	coupling_h,
	temperature float64,
) float64 {

	sum1 := 0.0
	sum2 := 0.0
	for ix:=0; ix<nx; ix++ {
		ixp1 := (ix+1)%nx
		for iy:=0; iy<ny; iy++ {
			iyp1 := (iy+1)%ny
			sum1 += float64(spin[ix][iy])
			sum2 += float64(spin[ix][iy]*spin[ixp1][iy] + spin[ix][iy]*spin[ix][iyp1])
		}
	}
	action := - (sum2*coupling_J + sum1*coupling_h) / temperature
	return action
}

/*** Calculation of the total spin ***/
func calcTotalSpin (spin [][]int) int {
	totalSpin := 0
	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			totalSpin += spin[ix][iy]
		}
	}

	return totalSpin
}

/*** Construct the cluster ***/
func makeCluster (
	spin [][]int,
	coupling_J,
	temperature float64,
	iCluster [][2]int,
) (int, int) {
	isIn := make ([][]bool, nx)
	for ix:=0; ix<nx; ix++ {
		isIn[ix] = make ([]bool, ny)
	}

	// choose a point randomly
	ix := rand.Intn (nx)
	iy := rand.Intn (ny)
	isIn[ix][iy] = true
	iCluster[0][0] = ix
	iCluster[0][1] = iy
	spinCluster := spin[ix][iy]
	nCluster := 1
	probability := 1.0 - math.Exp (-2.0*coupling_J/temperature)

	dict := [][]int{{1,0},{(nx-1),0},{0,1},{0,(ny-1)}}

	k := 0
	for k < nCluster {
		for _, v := range dict {
			ix2 := (ix+v[0])%nx
			iy2 := (iy+v[1])%ny
			if (spin[ix2][iy2]==spinCluster) && (!isIn[ix2][iy2]) && (rand.Float64()<probability) {
				iCluster[nCluster][0] = ix2
				iCluster[nCluster][1] = iy2
				nCluster ++
				isIn[ix2][iy2] = true
			}
		}
		k ++
	}
	return nCluster, spinCluster
}

/*** Main ***/
func main () {
	spin := make ([][]int, nx)
	for ix:=0; ix<nx; ix++ {
		spin[ix] = make ([]int, ny)
		for iy:=0; iy<ny; iy++ {
			spin[ix][iy] = 1
		}
	}
	rand.Seed (time.Now ().UnixNano())

	fp, err := os.Create ("output_config.txt")
	if err != nil {
		panic (err)
	}
	defer fp.Close ()

	iCluster := make([][2]int, nx*ny)
	/*** Main part ***/
	for iter:=int64(0); iter<niter; iter++ {
		nCluster, spinCluster := makeCluster (spin, coupling_J, temperature, iCluster)

		metropolis := rand.Float64()
		prob := math.Exp (-2.0*coupling_h*float64(spinCluster*nCluster)/temperature)
		if prob > metropolis {
			for k:=0; k<nCluster; k++ {
				ix := iCluster[k][0]
				iy := iCluster[k][1]
				spin[ix][iy] *= -1
			}
		}

		totalSpin := calcTotalSpin (spin)
		energy := calcAction (spin, coupling_J, coupling_h, temperature) * temperature

		/*** data output ***/
		if (iter+1)%nskip == 0 {
			fmt.Printf ("%4d\t%.4f\n", totalSpin, energy)
		}

		/*** config output ***/
		if (nconfig > 0) && ((iter+1)%nconfig == 0) {
			for ix:=0; ix<nx; ix++ {
				for iy:=0; iy<ny; iy++ {
					fmt.Fprintf (fp, "%d %d %d\n", ix, iy, spin[ix][iy])
				}
			}
		}
	}
}

