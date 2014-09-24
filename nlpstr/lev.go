/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Levenshtein (edit) Distance
 */
package nlpstr

import "github.com/jbowles/nlpt/nlp_math"

// EX: fmt.Println(LevenshteinThree("stri","str"))
func Levenshtein(s1, s2 string) int {
	m1 := len(s1)
	n2 := len(s2)
	width := n2 - 1
	vcell := make([]int, m1*n2)

	// cells for i of m(s1)
	for i1 := 1; i1 < m1; i1++ {
		vcell[i1*width+0] = i1
	}

	// cell for j of n(s2)
	for j2 := 1; j2 < n2; j2++ {
		vcell[0*width+j2] = j2
	}

	for j2 := 1; j2 < n2; j2++ {
		for i1 := 1; i1 < m1; i1++ {
			if s1[i1] == s2[j2] {
				vcell[i1*width+j2] = vcell[(i1-1)*width+(j2-1)]
			} else {
				deletion := vcell[(i1-1)*width+j2] + 1
				insertion := vcell[(i1*width+(j2-1))] + 1
				substitution := vcell[((i1-1)*width+(j2-1))] + 1
				vcell[i1*width+j2] = nlp_math.MinInt32(deletion, insertion, substitution)
			}
		}
	}
	return vcell[m1*width+0]
}
