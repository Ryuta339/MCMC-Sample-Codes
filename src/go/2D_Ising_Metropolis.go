/**************************************
 *** 2d Ising model with Metropolis ***
 **************************************/
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
	niter int = 4096000
	nx int = 64					// number of sites along x-direction
	ny int = 64					// number of sites along y-direction
	coupling_J float64 = 1.0
	coupling_h float64 = 0.0
	temperature float64 = 5.0
	nskip int = 40960			// Frequency of measurement
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
func calc_action (
	spin [][]int, 
	coupling_J, 
	coupling_h, 
	temperature float64,
) (action float64) {

	sum1 := 0
	sum2 := 0

	for ix:=0; ix<nx; ix++ {
		ixp1 := (ix+1)%nx
		for iy:=0; iy<ny; iy++ {
			iyp1 := (iy+1)%ny
			sum1 += spin[ix][iy]
			sum2 += spin[ix][iy]*spin[ixp1][iy] + spin[ix][iy]*spin[ix][iyp1]
		}
	}
	action = - (float64(sum2)*coupling_J + float64(sum1)*coupling_h)/temperature
	return
}

/*** Calculation of the change of the action
 ***  when the spin at (ix, iy) is flipped   ***/
func calc_action_change (
	spin [][]int,
	coupling_J,
	coupling_h,
	temperature float64,
	ix,
	iy int,
) (action_change float64) {

	ixp1 := (ix+1)%nx
	iyp1 := (iy+1)%ny
	ixm1 := (ix-1+nx)%nx
	iym1 := (iy-1+ny)%ny

	sum1_change := float64(2 * spin[ix][iy])
	sum2_change := float64(
		2 * spin[ix][iy] * spin[ixp1][iy] + 
		2 * spin[ix][iy] * spin[ix][iyp1] +
		2 * spin[ix][iy] * spin[ixm1][iy] + 
		2 * spin[ix][iy] * spin[ix][iym1])
	
	action_change = (sum2_change * coupling_J + sum1_change * coupling_h) / temperature
	return
}

/*** Calculation of the total spin ***/
func calc_total_spin (spin [][]int) (total_spin int) {
	total_spin = 0
	for ix:=0; ix<nx; ix++ {
		for iy:=0; iy<ny; iy++ {
			total_spin += spin[ix][iy]
		}
	}
	return
}

/*** Main ***/
func main () {
	spin := make ([][]int, nx)
	for ix:=0; ix<nx; ix++ {
		spin[ix] = make ([]int, ny)
	}
	rand.Seed (time.Now().UnixNano())

	nconfig.initialize (spin)

	/*** Main part ***/
	naccept := 0
	for iter:=0; iter<niter; iter++ {
		// choose a point randomly
		ix := rand.Intn (nx)
		iy := rand.Intn (ny)

		metropolis := rand.Float64 ()
		action_change := calc_action_change (spin, coupling_J, coupling_h, temperature, ix, iy)
		if math.Exp (-action_change) > metropolis {
			// accept
			spin[ix][iy] = -spin[ix][iy]
			naccept ++
		} else {
			// reject
		}
		total_spin := calc_total_spin (spin)
		energy := calc_action (spin, coupling_J, coupling_h, temperature) * temperature

		/*** data output ***/
		if ((iter+1)%nskip == 0) {
			fmt.Printf ("%4d\t%.4f\t%.4f\n",
				total_spin,
				energy,
				float64(naccept)/float64(iter+1))
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
