// Copyright ©2012 The bíogo.cluster Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package kmeans implements Lloyd's k-means clustering for ℝⁿ data.
package kmeans

import (
	"errors"
	"math/rand"

	"github.com/jbowles/nlpt/cluster"
)

type point []float64

func (p point) V() []float64 { return p }

type value struct {
	point
	w       float64
	cluster int
}

func (v *value) Weight() float64 { return v.w }
func (v *value) Cluster() int    { return v.cluster }

type center struct {
	point
	w       float64
	count   int
	indices cluster.Indices
}

func (c *center) zero() {
	p := c.point
	for i := range p {
		p[i] = 0
	}
	*c = center{point: p}
}

func (c *center) Members() cluster.Indices { return c.indices }

// Kmeans implements clustering of ℝⁿ data according to the Lloyd k-means algorithm.
type Kmeans struct {
	dims   int
	values []value
	means  []center
}

// New creates a new k-means object populated with data from an Interface value, data.
func New(data cluster.Interface) (*Kmeans, error) {
	v, d, err := convert(data)
	if err != nil {
		return nil, err
	}
	return &Kmeans{
		dims:   d,
		values: v,
	}, nil
}

// convert renders data to the internal float64 representation for a Kmeans.
func convert(data cluster.Interface) ([]value, int, error) {
	va := make([]value, data.Len())
	if data.Len() == 0 {
		return nil, 0, errors.New("kmeans: no data")
	}
	dim := len(data.Values(0))
	for i := 0; i < data.Len(); i++ {
		vec := data.Values(i)
		if len(vec) != dim {
			return nil, 0, errors.New("kmeans: mismatched dimensions")
		}
		va[i] = value{point: append(point(nil), vec...)}
	}
	if w, ok := data.(cluster.Weighter); ok {
		for i := 0; i < data.Len(); i++ {
			va[i].w = w.Weight(i)
		}
	} else {
		for i := 0; i < data.Len(); i++ {
			va[i].w = 1
		}
	}

	return va, dim, nil
}

// Seed generates the initial means for the k-means algorithm according to the k-means++
// algorithm
func (km *Kmeans) Seed(k int) {
	km.means = make([]center, k)
	for i := range km.means {
		km.means[i].point = make(point, km.dims)
	}

	copy(km.means[0].point, km.values[rand.Intn(len(km.values))].point)
	if k == 1 {
		return
	}
	d := make([]float64, len(km.values))
	for i := 1; i < k; i++ {
		sum := 0.
		for j, v := range km.values {
			_, min := km.nearest(v.point)
			d[j] = min
			sum += d[j]
		}
		target := rand.Float64() * sum
		j := 0
		for sum = d[0]; sum < target; sum += d[j] {
			j++
		}
		copy(km.means[i].point, km.values[j].point)
	}
}

// SetCenters sets the locations of the centers to c.
func (km *Kmeans) SetCenters(c []cluster.Center) {
	km.means = make([]center, len(c))
	for i, cv := range c {
		km.means[i] = center{point: append(point(nil), cv.V()...)}
	}
}

// Find the nearest center to the point v. Returns c, the index of the nearest center
// and min, the square of the distance from v to that center.
func (km *Kmeans) nearest(v point) (c int, min float64) {
	var ad float64
	for j := range v {
		ad = v[j] - km.means[0].point[j]
		min += ad * ad
	}

	for i := 1; i < len(km.means); i++ {
		var d float64
		for j := range v {
			ad = v[j] - km.means[i].point[j]
			d += ad * ad
		}
		if d < min {
			min = d
			c = i
		}
	}

	return c, min
}

// Cluster runs a clustering of the data using the k-means algorithm.
func (km *Kmeans) Cluster() error {
	if len(km.means) == 0 {
		return errors.New("kmeans: no centers")
	}
	for i, v := range km.values {
		n, _ := km.nearest(v.point)
		km.values[i].cluster = n
	}

	for {
		for i := range km.means {
			km.means[i].zero()
		}
		for _, v := range km.values {
			for j := range km.means[v.cluster].point {
				km.means[v.cluster].point[j] += v.point[j] * v.w
			}
			km.means[v.cluster].w += v.w
			km.means[v.cluster].count++
		}
		for i := range km.means {
			inv := 1 / km.means[i].w
			for j := range km.means[i].point {
				km.means[i].point[j] *= inv
			}
		}

		deltas := 0
		for i, v := range km.values {
			if n, _ := km.nearest(v.point); n != v.cluster {
				deltas++
				km.values[i].cluster = n
			}
		}
		if deltas == 0 {
			break
		}
	}
	return nil
}

// Total calculates the total sum of squares for the data relative to the data mean.
func (km *Kmeans) Total() float64 {
	p := make([]float64, km.dims)
	for _, v := range km.values {
		for j := range p {
			p[j] += v.point[j]
		}
	}
	inv := 1 / float64(len(km.values))
	for j := range p {
		p[j] *= inv
	}

	var ss float64
	for _, v := range km.values {
		for j := range p {
			d := p[j] - v.point[j]
			ss += d * d
		}
	}

	return ss
}

// Within calculates the sum of squares within each cluster.
// Returns nil if Cluster has not been called.
func (km *Kmeans) Within() []float64 {
	if km.means == nil {
		return nil
	}
	ss := make([]float64, len(km.means))

	for _, v := range km.values {
		for j := range v.point {
			d := km.means[v.cluster].point[j] - v.point[j]
			ss[v.cluster] += d * d
		}
	}

	return ss
}

// Centers returns the k centers determined by a previous call to Cluster.
func (km *Kmeans) Centers() []cluster.Center {
	c := make([]cluster.Indices, len(km.means))
	for i := range c {
		c[i] = make([]int, 0, km.means[i].count)
	}
	for i, v := range km.values {
		c[v.cluster] = append(c[v.cluster], i)
	}

	cs := make([]cluster.Center, len(km.means))
	for i := range km.means {
		km.means[i].indices = c[i]
		cs[i] = &km.means[i]
	}

	return cs
}

// Values returns a slice of the values in the Kmeans.
func (km *Kmeans) Values() []cluster.Value {
	vs := make([]cluster.Value, len(km.values))
	for i := range km.values {
		vs[i] = &km.values[i]
	}
	return vs
}
