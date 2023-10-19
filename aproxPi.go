/*
* The idea is to generate random points inside a square.
* This square has a circle inscribed inside it.
* If we divide the number of points inside the circle, by the total points generated
* and multiply it by 4 we optain a number that aproximates Pi.
*
*	#########
*	# #   # #
*	##     ##
*	#       #
*	#       #
*	#       #
*	##     ##
*	# #   # #
*	#########
*
 */

package main

import (
	"fmt"
	"math"
	"math/rand"
)

// Side of the square = 1.000
const side float32 = 1000

func main() {

	var iterations = []int{1, 10, 100, 270, 1000, 1000000, 1000000000}

	for _, it := range iterations {
		fmt.Printf("%d points:\n", it)

		var in int
		var x, y float32

		for i := 0; i < it; i++ {
			x = rand.Float32() * side
			y = rand.Float32() * side

			if math.Hypot(float64(x-side/2), float64(y-side/2)) < float64(side/2) {
				in++
			}
		}

		if in == 0 || it-in == 0 {
			fmt.Printf(" Too few points generated.\n  Points inside circle  = %d\n  Points outside circle = %d\n", in, it-in)
		} else {
			fmt.Printf(" Aproximated π is %0.6f.\n  Points inside circle  = %d\n  Points outside circle = %d\n", (float32(in)/float32(it))*4.0, in, it-in)
		}
	}
}
