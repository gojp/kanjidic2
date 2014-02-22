package kanjidic2

import (
	"testing"
)

var numKanji = 13108
var kanjiList []Kanji

func TestParser(t *testing.T) {
	got, err := ParseKanjiDic2("kanjidic2.xml")
	if err != nil {
		t.Fatalf("ParseKanjiDic2: %v", err)
	}
	if len(got) != numKanji {
		t.Fatalf("ParseKanjiDic2 length incorrect: got %d, want %d", got, numKanji)
	}
	// set kanjiList for later tests
	kanjiList = got
}

var strokeCountTests = []struct {
	Literal string
	Count   int
}{
	{"ー", 1},
	{"二", 2},
	{"三", 3},
	{"中", 4},
	{"店", 8},
	{"感", 13},
	{"機", 16},
}

func TestStrokeCount(t *testing.T) {
	for _, k := range kanjiList {
		for _, tt := range strokeCountTests {
			if k.Literal == tt.Literal {
				if got := k.StrokeCount; got != tt.Count {
					t.Errorf("TestStrokeCount (%s): got %d, want %d", tt.Literal, got, tt.Count)
				}
			}
		}
	}
}
