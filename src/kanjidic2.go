package main

import (
	"encoding/xml"
	"fmt"
	"log"
	"os"
)

type Reading struct {
	RType         string `xml:"r_type,attr"`
	ReadingString string `xml:",innerxml"`
}

type Meaning struct {
	MLang         string `xml:"m_lang,attr"`
	MeaningString string `xml:",innerxml"`
}

type DicRef struct {
	DrType    string `xml:"dr_type,attr"`
	RefNumber string `xml:",innerxml"`
}

type Kanji struct {
	Literal     string    `xml:"literal"`
	Grade       int       `xml:"misc>grade"`
	StrokeCount int       `xml:"misc>stroke_count"`
	Freq        int       `xml:"misc>freq"`
	JLPT        int       `xml:"misc>jlpt"`
	DicRefs     []DicRef  `xml:"dic_number>dic_ref"`
	Readings    []Reading `xml:"reading_meaning>rmgroup>reading"`
	Meanings    []Meaning `xml:"reading_meaning>rmgroup>meaning"`
	Nanori      []string  `xml:"nanori"`
}

func ParseKanjiDic2(filename string) (kanjiList []Kanji) {
	xmlFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer xmlFile.Close()
	decoder := xml.NewDecoder(xmlFile)
	for {
		token, _ := decoder.Token()
		if token == nil {
			break
		}
		switch startElement := token.(type) {
		case xml.StartElement:
			if startElement.Name.Local == "character" {
				var kanji Kanji
				decoder.DecodeElement(&kanji, &startElement)
				kanjiList = append(kanjiList, kanji)
			}
		}
	}
	return
}

func main() {
	kanjidic := ParseKanjiDic2("/Users/shawn/Downloads/kanjidic2.xml")
	fmt.Println(kanjidic[0])
}
