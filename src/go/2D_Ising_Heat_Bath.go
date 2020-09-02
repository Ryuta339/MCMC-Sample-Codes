package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	niter int64 = 4096000
	nx int = 64					// number of sites along x-direction
	ny int = 64					// number of sites along y-direction
	coupling_J float64 = 1.0
	coupling_h float64 = 0.1
	temperature float64 = 5.0
	nskip = 40960			// Frequency of measurement
)

/*** Configuration ***/
type Configuration interface {
	initialize (spin [][]int)
}

/*** Up Configuration ***/
type Up struct {
}
func (up *Up) initialize (spin [][]int) {
	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			spin[ix][iy] = 1
		}
	}
}
func NewUp () *Up {
	return &Up {}
}

/*** Down Configuration ***/
type Down struct {
}
func (down *Down) initialize (spin [][]int) {
	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			spin[ix][iy] = -1
		}
	}
}
func NewDown () *Down {
	return &Down {}
}

/*** Down Configuration ***/
type Read struct {
	filename string
}
func (read *Read) initialize (spin [][]int) {
	fp, err := os.Open (read.filename)
	if err!= nil {
		panic (err)
	}
	defer fp.Close ()			// file close in the end

	scanner := bufio.NewScanner (fp)
	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			elems := strings.Split (scanner.Text (), " ")
			spin[ix][iy], _ = strconv.Atoi (elems[2])
		}
	}
}
func NewRead (filename string) *Read {
	return &Read { filename: filename }
}

var nconfig Configuration = NewUp ()

/*** Calculation of the action ***/
func calcAction (spin [][]int, coupling_J, coupling_h, temperature float64) float64 {
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
	action := -(sum2*coupling_J + sum1*coupling_h) / temperature
	return action
}

/*** Calculation of the probability          ***
 ***     when the spin at (ix,iy) is updated ***/
func heatBathProbability (
	spin [][]int,
	coupling_J,
	coupling_h,
	temperature float64,
	ix,
	iy int,
) float64 {

	var Ep, Em float64

	ixp1 := (ix+1)%nx
	iyp1 := (iy+1)%ny
	ixm1 := (ix+nx-1)%nx
	iym1 := (iy+ny-1)%ny

	temp := coupling_h
	temp += float64( spin[ixp1][iy] + spin[ix][iyp1] + spin[ixm1][iy] + spin[ix][iym1]) * coupling_J

	Ep = - temp / temperature
	Em = temp / temperature

	ratio := math.Exp (-Ep) / (math.Exp (-Ep) + math.Exp (-Em))
	return ratio
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

/*** Main ***/
func main () {
	spin := make([][]int, nx)
	for ix:=0; ix<nx; ix++ {
		spin[ix] = make ([]int, ny)
	}
	rand.Seed (time.Now().UnixNano ())

	/* Set the initial configuration */
	nconfig.initialize (spin)

	/*** Main part ***/
	for iter:=int64(0); iter<niter; iter++ {
		// choose a point randomly
		ix := rand.Intn (nx)
		iy := rand.Intn (ny)
		metropolis := rand.Float64 ()
		ratio := heatBathProbability (spin, coupling_J, coupling_h, temperature, ix, iy)

		if ratio > metropolis {
			spin[ix][iy] = 1
		} else {
			spin[ix][iy] = -1
		}

		totalSpin := calcTotalSpin (spin)
		energy := calcAction (spin, coupling_J, coupling_h, temperature) * temperature

		/*** data output ***/
		if (iter+1)%nskip == 0 {
			fmt.Printf ("%4d\t%.4f\n", totalSpin, energy)
		}
	}

	/*** save final config ***/
	fp, err := os.Create ("output_config.txt")
	if err != nil {
		panic (err)
	}
	defer fp.Close ()			// file close in the end

	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			fmt.Fprintf (fp, "%d %d %d\n", ix, iy, spin[ix][iy])
		}
	}
}
