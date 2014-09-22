/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* minmax define functions for determining a Min and/or Max int32
 */
package nlp_math

const (
	MaxInteger = int(^uint(0) >> 1) // max largest int 9223372036854775807
	MinInteger = (-MaxInteger - 1)
)

func MinInt(a ...int) int {
	min := MaxInteger
	for _, i := range a {
		if i < min {
			min = i
		}
	}
	return min
}

func MaxInt(a ...int) int {
	max := int(0)
	for _, i := range a {
		if i > max {
			max = i
		}
	}
	return max
}
