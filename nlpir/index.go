/*
* Copyright ©2013 The nlpt Authors. All rights reserved.
* Use of this source code is governed by a BSD-style
* license that can be found in the LICENSE file.
*
* Index allows creation of simple index types:
* Create a forward index: simple mapping for word index to document number
* Create an inverted index: simple mapping for string to string index and document number
 */
package nlpir

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Types
type Document struct {
	docNum   int
	wordIndx int
}

type ForwardIndex map[int]string
type InvertedIndex map[string][]Document
type IndexContainer struct {
	iIndex *InvertedIndex
	fIndex *ForwardIndex
}

////////////////////////// Initialize Indexes ////////////////
// Initialize new Forward Index
func NewForwardIndex() *ForwardIndex {
	i := make(ForwardIndex)
	return &i
}

// Initialize new Inverted Index
func NewInvertedIndex() *InvertedIndex {
	i := make(InvertedIndex)
	return &i
}

///////////////////// Find items in Indexes ////////////////
// Find item by index in Forward Index
func (idx *ForwardIndex) ItemAt(i int) string {
	return (*idx)[i]
}

// Find item by word in Inverted Index
func (idx *InvertedIndex) Search(query string) []Document {
	ref, ok := (*idx)[query]
	if ok {
		//fmt.Printf("***Results for query: '%v'\n", query)
		return ref
	}
	return nil
}

////////////////////////// Size of Indexes ////////////////
// Size of Inverted Index
func (idx *InvertedIndex) Size() int {
	return len(map[string][]Document(*idx))
}

// Size of Forward Index
func (idx *ForwardIndex) Size() int {
	return len(*idx)
}

/////////////// Add documents to Indexes ////////////////
// AddDoc for Forward Index
func (idx *ForwardIndex) AddDoc(docNum int, doc string) {
	for _, word := range strings.Fields(doc) {
		_, ok := (*idx)[docNum]
		if !ok {
			(*idx)[docNum] = word
		}
	}
}

// AddDoc for Inverted Index
func (idx *InvertedIndex) AddDoc(docNum int, doc string) {
	for wordIndx, word := range strings.Fields(doc) {
		ref, ok := (*idx)[word]
		if !ok {
			ref = nil
		}

		(*idx)[word] = append(ref, Document{docNum: docNum, wordIndx: wordIndx})
	}
}

// Initilialize index creation:
// read file/line; craete bloom filter, add to inverted/forward indices
func InitIndex(iIndex *InvertedIndex, fIndex *ForwardIndex, corpusPath string) {
	// Get Corpus Documents
	file, file_err := os.Open(corpusPath)
	if file_err != nil {
		fmt.Println("Error reading file:", file_err)
	}
	defer file.Close()

	docID := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scan_err := scanner.Err(); scan_err != nil {
			fmt.Println("Scan Line read error:", scan_err)
		}
		scan_line := scanner.Text()

		iIndex.AddDoc(docID, scan_line) //insert into inverted index
		fIndex.AddDoc(docID, scan_line) //insert into forward index

		docID++
	}

}
