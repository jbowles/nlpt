/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* tfidf == Term Frequency Inverse Document Frequency.
 */
package nlptir

import (
	"math"
	"sort"
	"strings"
)

// TYPES //////////////////////////////////////////////////

// Vector contains values for tf-idf value, document number, and index location of token/term for quicker lookup
type Vector struct {
	docNum     int
	index      int
	dotProduct float64
}

// Field contains a space of the map of the token/term to its Vectors
type VecField struct {
	Space map[string][]Vector
}

// A data structure to hold a key/value pair.
type Pair struct {
	Key   string
	Value []Vector
}

// A slice of Pairs that implements sort.Interface to sort by Value of Hash Map.
type PairList []Pair

// SORT //////////////////////////////////////////////////

// Create needed Sort methods: Len(), Less(), Swap()
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Len() int           { return len(p) }
func (p PairList) Less(i, j int) bool { return p[i].Value[0].dotProduct < p[j].Value[0].dotProduct }

// A function to turn a map into a PairList, then sort and return it.
func (m VecField) sortByTfIdf() PairList {
	p := make(PairList, len(m.Space))
	i := 0
	for k, v := range m.Space {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	return p
}

// TF-IDF //////////////////////////////////////////////////

// TermCount gets the total count of all terms in a document
// param:
// return:
func TermCount(doc string) float64 {
	words := strings.Split(doc, " ") // TODO: use the tokenizer here
	return float64(len(words))
}

// TokenFreq gets the frequency of term in a document
// param:
// return:
func TokenFreq(word, doc string) float64 {
	total := TermCount(doc)
	count := 0.0
	//TODO: replace strings.Fields with tokenizer
	for _, w := range strings.Fields(doc) {
		switch w {
		case word:
			count += 1.0
		}
	}
	return count / total
}

// NumDocsContain calculates the number of documents that cotain one term
// param:
// return:
func NumDocsContain(word string, doc_list []string) (count float64) {
	for _, doc := range doc_list {
		if TokenFreq(word, doc) > 0.0 {
			count += 1.0
		}
	}
	return
}

// Tf is the technical term frequency of tf-idf
// param:
// return:
func Tf(word, doc string) float64 {
	return (TokenFreq(word, doc) / TermCount(doc))
}

// Idf is the inverse document frequency of tf-idf
// param:
// return:
func Idf(word string, doc_list []string, log string) (idf float64) {
	// set val for reuse; +1 so we don't get +Inf values
	val := float64(len(doc_list)+1) / (NumDocsContain(word, doc_list) + 1)
	switch log {
	case "log":
		idf = math.Log(val) //Log returns the natural logarithm of x.
	case "log10":
		idf = math.Log10(val) //Log10 returns the decimal logarithm of x.
	case "nolog":
		idf = val //no logarithm
	case "log1p":
		idf = math.Log1p(val) //Log1p natural log of 1 plus its argument x
	case "log2":
		idf = math.Log2(val) //Log2 returns the binary log of x.
	default:
		idf = math.Log(val)
	}
	return
}

// TfIdf returns the Term Frequency-Inverse Document Frequency for a word and all documents
func TfIdf(word, doc string, doc_list []string, log string) float64 {
	return (Tf(word, doc) * Idf(word, doc_list, log))
}

func (f *VecField) Compose(documents []string) {
	//initialize Space map
	f.Space = make(map[string][]Vector)
	for doc_num, doc := range documents {
		for idx, word := range strings.Fields(doc) {
			v, ok := f.Space[word]
			if !ok {
				v = nil
			}
			tfidf_product := TfIdf(word, doc, documents, "log")
			f.Space[word] = append(v, Vector{doc_num, idx, tfidf_product})
		}
	}
}
