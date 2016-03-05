package tkz

/*
Wikipedia streaming XML parser.
Based mostly on code and blog post by David Singleton:
  https://github.com/dps/go-xml-parse/blob/master/go-xml-parse.go
  http://blog.davidsingleton.org/parsing-huge-xml-files-with-go/
*/

import (
	"bufio"
	"encoding/xml"
	"log"
	"net/url"
	"os"
	"regexp"
	"strings"
	"sync"
)

// Here is an example article from the Wikipedia XML dump
//
// <page>
// 	<title>Apollo 11</title>
//      <redirect title="Foo bar" />
// 	...
// 	<revision>
// 	...
// 	  <text xml:space="preserve">
// 	  {{Infobox Space mission
// 	  |mission_name=&lt;!--See above--&gt;
// 	  |insignia=Apollo_11_insignia.png
// 	...
// 	  </text>
// 	</revision>
// </page>
//
// Note how the tags on the fields of Page and Redirect below
// describe the XML schema structure.

type Redirect struct {
	Title string `xml:"title,attr"`
}

type Page struct {
	Title string   `xml:"title"`
	Redir Redirect `xml:"redirect"`
	Text  string   `xml:"revision>text"`
}

func CanonicalizeTitle(title string) string {
	can := strings.ToLower(title)
	can = strings.Replace(can, " ", "_", -1)
	can = url.QueryEscape(can)
	return can
}

func WritePage(tokenDelimiter byte, txtByte []byte, outFile, title string) {
	f, err := os.OpenFile(outFile, os.O_APPEND|os.O_WRONLY, 0600)
	if err != nil {
		panic(err)
	}
	writer := bufio.NewWriter(f)
	defer f.Close()
	var writeInput []byte
	for _, token := range txtByte {
		writeInput = append(writeInput, token+tokenDelimiter)
	}
	writer.Write(txtByte)
	log.Printf("wikipedia page %s...\n", title)
	writer.Flush()
}

func StreamTokenizedWikipediaDump(wg *sync.WaitGroup, tokenDelim byte, inFile, outFile, tkzTyp string) {
	var filter, _ = regexp.Compile("^file:.*|^talk:.*|^special:.*|^wikipedia:.*|^wiktionary:.*|^user:.*|^user_talk:.*")
	//overwrite the output file and close, so subsuequent runs don't append but pipes do
	f, ferr := os.Create(outFile)
	f.Close()
	if ferr != nil {
		wg.Done()
		panic(ferr)
	}
	//overwritten output file and closed it so we start fresh with Append

	xmlFile, err := os.Open(inFile)
	if err != nil {
		wg.Done()
		panic(err)
	}
	defer xmlFile.Close()

	decoder := xml.NewDecoder(xmlFile)
	total := 0
	var inElement string
	for {
		// Read tokens from the XML document in a stream.
		t, _ := decoder.Token()
		if t == nil {
			break
		}
		// Inspect the type of the token just read.
		switch se := t.(type) {
		case xml.StartElement:
			// If we just read a StartElement token
			inElement = se.Name.Local
			// ...and its name is "page"
			if inElement == "page" {
				var p Page
				// decode chunk of XML into variable p Page
				decoder.DecodeElement(&p, &se)

				// Do some stuff with the page.
				//p.Title = CanonicalizeTitle(p.Title)
				m := filter.MatchString(p.Title)
				if !m && p.Redir.Title == "" {
					WritePage(tokenDelim, TokenizeBytes([]byte(p.Text), tkzTyp).Bytes, outFile, p.Title)
					total++
				}
			}
		default:
		}

	}

	log.Printf("Total articles: %d \n", total)
	wg.Done()
}
