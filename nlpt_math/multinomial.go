/*
* Copyright ©2014 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* multinomial provides the Multinomial Distribution
 */
package nlpt_math

// Multinomial Counts
type MultiCount struct {
	counts map[int]int
}

// Multinomial Distribution
type Multinomial struct {
	distribution map[int]float64
}

func (mc *MultiCount) Count() {
	mc.counts = make(map[int]int)
}

func (mc *MultiCount) Increment(key int) {
	if _, ok := mc.counts[key]; ok {
		mc.counts[key] = 0
	}
	mc.counts[key]++
}

func (multi *Multinomial) Probability(key int) float64 {
	return multi.distribution[key]
}

func (multi *Multinomial) MaximumLikelihood(mc MultiCount) {
	multi.distribution = make(map[int]float64)
	sum := 0.0
	for _, count := range mc.counts {
		sum += float64(count)
	}
	for key, count := range mc.counts {
		multi.distribution[key] = float64(count) / sum
	}
}
