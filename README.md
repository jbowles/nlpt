# Natural Language Processing Toolkit
Written in Go.

Gopher images use Creative Commons Attributions 3.0. All credit belongs to **Renee French**.
See blog post http://blog.golang.org/gopher.

Progress has been slow, but I am fully comitted to NLPT; even if there are silent spells. Some of my work on NLP and Go cannot be open-sourced for given periods of time, and now I have chance to be working in a more open-source friendly capacity!

I am a linguist by trade and so NLP is always something I want to write code for. I now have the chance to work on Go professionally in a more open-source friendly capacity and this bodes well for this project. 

[siw](https://github.com/jbowles/siw) is the Simple Words project and is the place I've playing out some basic ideas in shorter form. If you don't see any progress on `nlpt`, you can be sure that `siw` has got something brewing.

[exp branch](https://github.com/jbowles/nlpt/tree/exp) is the experimental branch; this will be the default branch until more progress is made on the core architecture (tokenzier, stemmer, and tagger).

This is the tested, stable, and production ready brnach of a research project to write natural language processing tools in Go. NLPT is built up from multiple sub-packages (each separately accessible).

Get it:
* `go get github.com/jbowles/nlpt` 
  * (or update: `go get -u github.com/jbowles/nlpt`)

Functionality is separated into sub packages, which are usable outside the scope of the main NLPT project. Naming of each subpackage will be consistent as per the `first 3 letter prefix` + `subpackage name`. For example: tokenizer = `nlptokenizer`, stemmer = `nlpstemmer`, tagger = `nlptagger`.
Get a subpackage:
* `go get github.com/jbowles/nlpt/nlptokenizer` 
  * (or update: `go get -u github.com/jbowles/nlpt/nlptokenizer`)

Thanks to the Go Berlin users group for letting me ripoff their gopher image!

![Alt text](https://github.com/jbowles/nlpt/raw/exp/nlpt.jpg "Natural Language Processing Toolkit in Go")

## Branches
* `exp`
  * Low-level development and general messiness
* `stable`
  * Testing, performance, run standard data sets, P&R (precision and recall) where appropriate

Development workflow == `exp` -> `stable` -> `master`

## Criteria each sub-package:
* `Stability` (Experimental, Stable, Production) to determine whether the API is production ready. 
* `Volatility` (Radical, Mild, Stable) to determine whether the API is likely to change.
* `Test` (Nil, Some, Stable) to signal range of coverage for tests over the API.
* `Examples` link to external repo with more documentation and examples.

## General
NLPT broadly supports minimal functionality for the full range of 4 bit unicode code points.


## Tokenizer

* **Done**:
  * Basic Whitespace tokenizer
   * only whitespace, cannot parse punctuation
  * Unicode tokenizer (good for noisy data sets)
   * captures all unicode code points into a `Bucket`
   * cannot return sequences of numbers (e.g., dates)
   * will reconstruct 'nosiy" text (e.g., `th0s7!e` => `["thse"], [0, 7], ["!"]`)

* **In Progress**:
  * State Machine-like lexer (inspired by [Rob Pike: Lexical Scanning in Go](http://cuddle.googlecode.com/hg/talk/lex.html#landing-slide))
  * Punkt algorithm (see python NLTK project)

Run tests and benchmarks:

```sh
go test -v
go test -benchmem -bench .
```

    Stability:  2   - Stable
    Volatility: 2   - Stable
    Tests:      2   - Stable
Examples:   []()

### TODO
* Support for Arabic and Mandarin are coming, though probably not until late 2014.
* Eventually move to a more probabilistic model.

### Description
Tokenizer it leverages the Go Rune Type (`int32` aliases for Unicode). Basically, you can **build custom unicode alphabets** that are used for pattern matching (instead of regular expressions). General goals:

1. Broader spectrum of Unicode characaters used across ever expanding and changing media
1. Special or nonstandard characters used in software application logs
1. The rise of new languages on the web, and the velocity of disappearing endangered langauges

### Sources
* Unicode code points table for Go Runes
  * [UTF8 Character Table (in decimal)](http://www.utf8-chartable.de/unicode-utf8-table.pl?utf8=dec)


## TF-IDF: Term Frequency-Inverse Document Frequency
The Tf-Idf stuff is not done, I've just been playing with different ways of doing it. There is not a full model finished yet and so the first implementation is not complete.

    Stability:  0   - Not Started
    Volatility: 0   - Not Started
    Tests:      0   - Not Started
Examples:  []()

## Stemmer

    Stability:  0   - Not Started
    Volatility: 0   - Not Started
    Tests:      0   - Not Started
Examples:  []()

## POS: Part of Speech Tagger

    Stability:  0   - Not Started
    Volatility: 0   - Not Started
    Tests:      0   - Not Started
Examples:  []()

# Resources

## Text Processing in Go
* [Go machine learning nlp libraries](http://biosphere.cc/software-engineering/go-machine-learning-nlp-libraries/)

## Tokenizer
* [Xerox Tokenizer Service](http://open.xerox.com/Services/fst-nlp-tools/Consume/175)

## Stemming
* [Xerox Morphological Analysis](http://open.xerox.com/Services/fst-nlp-tools/Pages/morphology) 

## Part of Speech Tags
* [Part of Speech Tagger in about 200 lines of Python](http://honnibal.wordpress.com/)
* [Xerox POS Tagger Service](http://open.xerox.com/Services/fst-nlp-tools/Consume/178)
* [Xerox POS Tagset Standard](http://open.xerox.com/Services/fst-nlp-tools/Pages/English%20Part-of-Speech%20Tagset)

## Tf-Idf
See the IDF entry in the [Information Retrieval](http://nlp.stanford.edu/IR-book/html/htmledition/inverse-document-frequency-1.html) book provided by Stanford and authors for detail.

## Unicode
* [Lorem Ipsum generator for multiple languages](http://generator.lorem-ipsum.info/)
  * Includes Arabic, Mandarin, Hebrew, Cyrillic, and others
* [Library of Congress Standards](http://www.loc.gov/standards/)
* [ISO 693-2 Standard for Natural Language Codes and Names](http://www.loc.gov/standards/iso639-2/php/code_list.php)
* [ISO 639-5 Standard for Natural Language Families and Names](http://www.loc.gov/standards/iso639-5/id.php)

## NLP APIs for result comaprison testing
One reason for knowing about NLP services is they can be used for testing and comparing results. When more sub-packages become available tests will be written against these APIs to compare results.

#### Xerox
* [Xerox Linguistic Tools](http://open.xerox.com/Services/fst-nlp-tools/Pages/API%20Docs)
  * "Finite State Technology Tools for Natural Language Processing"

  ```go
  /*
     ** DESCRIPTION FROM XEROX **
  These tools, called xfst, twolc, and lexc, are used in many linguistic applications such as morphological analysis, tokenisation, and shallow parsing of a wide variety of natural languages. The finite state tools here are built on top of a software library that provides algorithms to create automata from regular expressions and equivalent formalisms and contains both classical operations, such as union and composition, and new algorithms such as replacement and local sequentialisation.

  Finite-state linguistic resources are used in a series of applications and prototypes that range from OCR to terminology extraction, comprehension assistants, digital libraries and authoring and translation systems.

  The components provided here are:

  Tokenization
  Morphology
  Part of Speech Disambiguation (Tagging) 
  */
  ```

#### Alchemy
* [AlchemyApi](http://www.alchemyapi.com/)
  * "... cloud-based and on-premise text analysis infrastructure"

  ```go
  /*
  ** DESCRIPTION FROM ALCHEMY **
  AlchemyAPI uses natural language processing technology and machine learning algorithms to extract semantic meta-data from content, such as information on people, places, companies, topics, facts, relationships, authors, and languages.

  API endpoints are provided for performing content analysis on Internet-accessible web pages, posted HTML or text content.
  */
  ```
