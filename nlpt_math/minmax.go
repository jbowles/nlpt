/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* minmax define functions for determining a Min and/or Max int32
 */
package nlpt_math

import "math"

func MinInt32(a ...int) int {
	min := math.MaxInt32
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MaxInt32(a ...int) int {
	max := math.MinInt32 //int(0)
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
