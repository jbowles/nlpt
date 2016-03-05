/*
# Sources
Got most of these sources from the original word2vec project. Maybe rework this so its all Go and I can just have binary.

## News crawler

```sh
wget http://www.statmt.org/wmt14/training-monolingual-news-crawl/news.2012.en.shuffled.gz
wget http://www.statmt.org/wmt14/training-monolingual-news-crawl/news.2013.en.shuffled.gz
gzip -d news.2012.en.shuffled.gz
gzip -d news.2013.en.shuffled.gz
```

## 1B lang modelling benchmark

```sh
wget http://www.statmt.org/lm-benchmark/1-billion-word-language-modeling-benchmark-r13output.tar.gz
tar -xvf 1-billion-word-language-modeling-benchmark-r13output.tar.gz
```

## UMBC

```sh
wget http://ebiquity.umbc.edu/redirect/to/resource/id/351/umbc_webbase_corpus.tar.gz
tar -xvf umbc_webbase_corpus.tar.gz '*.txt'
wget http://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2
```


## Wikipedia

```
wget http://dumps.wikimedia.org/enwiki/latest/enwiki-latest-pages-articles.xml.bz2
bzip2 -c -d enwiki-latest-pages-articles.xml.bz2
```
*/

package main

import (
	tkz "github.com/jbowles/nlpt_tkz"
	"sync"
	"time"
)

var bigwordRootfp string = "/Users/jbowles/x/training_data/8_billion_word/"
var billionBenchmarkfp string = (bigwordRootfp + "1-billion-word-language-modeling-benchmark-r13output/")
var tkzr string = "unicode"

var BigWordData = map[int][]string{
	0: {"wikipedia", bigwordRootfp + "enwiki-latest-pages-articles.xml", "datasets/enwiki-latest-pages-articles-tokenized.txt"},
	1: {"file", bigwordRootfp + "news.2012.en.shuffled", "datasets/news.2012.en.shuffled.tokenized.txt"},
	2: {"file", bigwordRootfp + "news.2013.en.shuffled", "datasets/news.2013.en.shuffled.tokenized.txt"},
	3: {"dir", billionBenchmarkfp + "heldout-monolingual.tokenized.shuffled", "datasets/heldout-monolingual.tokenized.txt"},
	4: {"dir", billionBenchmarkfp + "training-monolingual.tokenized.shuffled", "datasets/training-monolingual.tokenized.txt"},
	5: {"dir", bigwordRootfp + "webbase_all", "datasets/umbc_webbase_tokenized.txt"},
}

func streamDataOps() {
	var wg sync.WaitGroup
	for _, value := range BigWordData {
		switch value[0] {
		case "dir":
			wg.Add(1)
			go tkz.StreamTokenizedDirectory(&wg, time.Minute*90, value[1], value[2], tkzr)
		case "file":
			wg.Add(1)
			go tkz.StreamTokenizedFile(&wg, time.Minute*90, value[1], value[2], tkzr)
		case "wikipedia":
			wg.Add(1)
			go tkz.StreamTokenizedWikipediaDump(&wg, value[1], value[2], tkzr)
		}
	}
	wg.Wait()
}

func main() {
	streamDataOps()
}
