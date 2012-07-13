package kanjidic2

import (
	"encoding/xml"
	"log"
	"os"
)

type Reading struct {
	RType string `xml:"r_type,attr"`
	Value string `xml:",innerxml"`
}

type Meaning struct {
	MLang string `xml:"m_lang,attr"`
	Value string `xml:",innerxml"`
}

type DicRef struct {
	DrType string `xml:"dr_type,attr"`
	Value  string `xml:",innerxml"`
}

type CpValue struct {
	CpType string `xml:"cp_type,attr"`
	Value  string `xml:",innerxml"`
}

type RadValue struct {
	RadType string `xml:"rad_type,attr"`
	Value   int    `xml:",innerxml"`
}

type Variant struct {
	VarType string `xml:"var_type,attr"`
	Value   string `xml:",innerxml"`
}

type QCode struct {
	QcType string `xml:"qc_type,attr"`
	Value  string `xml:",innerxml"`
}

type Kanji struct {
	Literal     string     `xml:"literal"`
	CodePoints  []CpValue  `xml:"codepoint>cp_value"`
	Radicals    []RadValue `xml:"radical>rad_value"`
	Grade       int        `xml:"misc>grade"`
	StrokeCount int        `xml:"misc>stroke_count"`
	Variant     Variant    `xml:"misc>variant"`
	Freq        int        `xml:"misc>freq"`
	JLPT        int        `xml:"misc>jlpt"`
	DicRefs     []DicRef   `xml:"dic_number>dic_ref"`
	QueryCodes  []QCode    `xml:"query_code>q_code"`
	Readings    []Reading  `xml:"reading_meaning>rmgroup>reading"`
	Meanings    []Meaning  `xml:"reading_meaning>rmgroup>meaning"`
	Nanori      []string   `xml:"nanori"`
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
