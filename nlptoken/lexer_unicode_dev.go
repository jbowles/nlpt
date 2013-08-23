package main

import "fmt"

/*
type Encodings interface {
	scope() (rune, rune)
}
*/

// create a new type struct that has CodePoint in it and utype

type TokenRange struct {
	cp     []CodePoint
	uniset []rune
}

type CodePoint struct {
	order []rune
	utyp  unicodeType
}

type unicodeType int

const (
	itemBasicLatin                         unicodeType = iota
	itemCyrillic                                       //itemType = iota
	itemSamaritan                                      //itemType = iota
	itemTelugu                                         //itemType = iota
	itemMyanmar                                        //itemType = iota
	itemUnifiedCanadianAboriginalSyllabics             //itemType = iota
	itemMongolian                                      //itemType = iota
	itemLepcha                                         //itemType = iota
	itemGeneralPunctuation                             //itemType = iota
	itemControlPictures                                //itemType = iota
	itemBraillePatterns                                //itemType = iota
	itemGlagolitic                                     //itemType = iota
	itemCjkSymbolsPunctuation                          //itemType = iota
	itemCjkUnifiedIdeographsExtA                       //itemType = iota
	itemCjkUnifiedIdeographs                           //itemType = iota
)

var (
	BasicLatin                         = CodePoint{order: []rune{0, 1023}, utyp: itemBasicLatin}                            //
	Cyrillic                           = CodePoint{order: []rune{1024, 2047}, utyp: itemCyrillic}                           //
	Samaritan                          = CodePoint{order: []rune{2048, 3071}, utyp: itemSamaritan}                          //
	Telugu                             = CodePoint{order: []rune{3072, 4095}, utyp: itemTelugu}                             //
	Myanmar                            = CodePoint{order: []rune{4096, 5119}, utyp: itemMyanmar}                            //
	UnifiedCanadianAboriginalSyllabics = CodePoint{order: []rune{5120, 6143}, utyp: itemUnifiedCanadianAboriginalSyllabics} //
	Mongolian                          = CodePoint{order: []rune{6144, 7167}, utyp: itemMongolian}                          //
	Lepcha                             = CodePoint{order: []rune{7168, 8191}, utyp: itemLepcha}                             //
	GeneralPunctuation                 = CodePoint{order: []rune{8192, 9125}, utyp: itemGeneralPunctuation}                 //
	ControlPictures                    = CodePoint{order: []rune{9216, 10239}, utyp: itemControlPictures}                   //
	BraillePatterns                    = CodePoint{order: []rune{10240, 11263}, utyp: itemBraillePatterns}                  //
	Glagolitic                         = CodePoint{order: []rune{11264, 12287}, utyp: itemGlagolitic}                       //
	CjkSymbolsPunctuation              = CodePoint{order: []rune{12288, 13311}, utyp: itemCjkSymbolsPunctuation}            // Chinese, Japanese, Korean
	CjkUnifiedIdeographsExtA           = CodePoint{order: []rune{13312, 20479}, utyp: itemCjkUnifiedIdeographsExtA}         // Chinese, Japanese, Korean
	CjkUnifiedIdeographs               = CodePoint{order: []rune{20480, 40959}, utyp: itemCjkUnifiedIdeographs}             // Chinese, Japanese, Korean
	//YiSyllables							= CodePoint{order: []rune{0, 0}}  //
)

//3072 - 4095 == Telugu
//4096 - 5119 == Myanmar
//5120 - 6143 == Unified Canadian Aboriginal Syllabics
//6144 - 7167 == Mongolian
//7168 - 8191 == Lepcha
//8192 - 9125 == General Punctuation
/*
func (cp CodePoint) scope() {
	//fmt.Println("cp order", cp.order)
	return
}
*/

func UnicodeSet(sets ...CodePoint) TokenRange {
	t := TokenRange{uniset: make([]rune, 0), cp: make([]CodePoint, 0)}

	for _, cop := range sets {
		startidx := cop.order[0]
		stopidx := cop.order[1]
		t.cp = append(t.cp, cop)
		tmp := make([]rune, stopidx, stopidx)
		for i := startidx; i < stopidx; i++ {
			//fmt.Println(tmp[i])
			tmp[i] = rune(i)
		}
		t.uniset = append(t.uniset, tmp[startidx:]...)
	}
	return t
}

func main() {
	//fmt.Println("BasicLatin:", UnicodeSet(BasicLatin))
	//fmt.Println(BasicLatin, Cyrillic)
	//fmt.Println(BasicLatin.order,len(BasicLatin.order),cap(BasicLatin.order))
	//fmt.Println("BasicLatin:", UnicodeSet(BasicLatin))
	th := UnicodeSet(BasicLatin, Cyrillic)
	//fmt.Println(th)

	for _, iot := range th.cp {
		fmt.Println("UnicodeType iota value", iot.utyp)
		for _, c := range th.uniset {
			fmt.Println("Char", string(c), "Decimal Code Point value:", c)
		}
	}
}
