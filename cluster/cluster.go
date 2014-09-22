// Copyright ©2012 The bíogo.cluster Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package cluster provides interfaces and types for data clustering in ℝⁿ.
package cluster

// Indices is a list of indexes into a array or slice of Values.
type Indices []int

// Clusterer is the common interface implemented by clustering types.
type Clusterer interface {
	// Cluster the data.
	Cluster() error

	// Centers returns a slice of centers of the clusters.
	Centers() []Center

	// Values returns the internal representation of the original data.
	Values() []Value
}

// Interface is a type that can be clustered by a Clusterer.
type Interface interface {
	Len() int               // Return the length of the data vector.
	Values(i int) []float64 // Return the data values for element i as a slice of float64.
}

// Weighter is an extension of the Interface that allows values represented by the Interface to be
// differentially weighted.
type Weighter interface {
	Weight(i int) float64 // Return the weight for element i.
}

// Point represents a point in ℝⁿ.
type Point interface {
	V() []float64
}

// A Value is the representation of a data point within the clustering object.
type Value interface {
	Point

	// Cluster returns an index into the slice returned by Clusterer.Centers() that
	// refers to the Center associated with the Value.
	Cluster() int
}

// A Center is a representation of a cluster center.
type Center interface {
	Point

	// Members returns a set of indices into the slice returned by Clusterer.Values() that
	// refers to the Values associated with the Center.
	Members() Indices
}
