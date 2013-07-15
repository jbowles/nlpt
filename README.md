# Natural Language Processing Toolkit
Written in Go.
This is the tested, stable, and production ready brnach of a research project to write natural language processing tools in Go. NLPT is built up from multiple sub-packages (each separately accessible).

Get it:
* `go get github.com/jbowles/nlpt` 
  * (or update: `go get -u github.com/jbowles/nlpt`)

Get a subpackage (for example, the tokenizer package `tkn`):
* `go get github.com/jbowles/nlpt/tkn` 
  * (or update: `go get -u github.com/jbowles/nlpt/tkn`)

![Alt text](https://github.com/jbowles/nlpt/raw/master/nlpt2.jpg "Natural Language Processing Toolkit in Go")
## Branches
* `stable` comes out of `experiment` and for writing tests and ensuring stability
* `experiment` is for low-level development and general messiness

Workflow == `experiment` -> `stable` -> `master`

## Criteria each sub-package:
* `Stability` (Experimental, Usable, Stable) to determine whether the API is production ready. 
* `Volatility` (Radical, Mild, Stable) to determine whether the API is likely to change.
* `Test` (Nil, Some, Stable) to signal range of coverage for tests over the API.
* `Examples` link to external repo with more documentation and examples.

## General
NLPT broadly supports minimal functionality for text in following language sets:

* Text in Roman alphabet with diacritics (English, Spanish, French, German, etc...)
* Text in Cyrillic alphabet (Russian, Belarusian, Ukrainian, Rusyn, Serbian, Bulgarin, Macedonian, Chechen, and other Slavic langauges.
* Text with Greek alphabet
* **Support for Arabic and Mandarin are coming in late 2014**

## Tokenizer

    Stability:  3   - Stable
    Volatility: 3   - Stable
    Tests:      3   - Stable
Examples:   []()

### TODO
* Support for Arabic and Mandarin are coming, though probably not until late 2014.
* Eventually move to probabilistic model.

### Description
Tokenizer **does not use regular expressions for its unicode character pattern matching**. Instead, it leverages the Go Rune Type (`int32` aliases for Unicode characters) and the String Types defined within Go Struct Types. Basically, you can **build custom unicode alphabets** that are used as an allowable set of characters for the tokenizing. Rivet Tokenizer seeks to provide an easy to use, flexible work flow that can adapt to novel requirements and be extended easily. I am building it with three main concerns:

1. Broader spectrum of Unicode characaters used across ever expanding and changing media
1. Special or nonstandard characters used in software application logs
1. The rise of new languages on the web, and the velocity of disappearing endangered langauges

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

## Tokenizer
* [Xerox Tokenizer Service](http://open.xerox.com/Services/fst-nlp-tools/Consume/175)

## Stemming
* [Xerox Morphological Analysis](http://open.xerox.com/Services/fst-nlp-tools/Pages/morphology) 

## Part of Speech Tags
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


