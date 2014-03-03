package kanjidic2

import "testing"

var numKanji = 13108
var kanjidic2Parser Kanjidic2Parser

func TestParser(t *testing.T) {
	got, err := ParseKanjiDic2("kanjidic2.xml")
	if err != nil {
		t.Fatalf("ParseKanjiDic2: %v", err)
	}
	if len(got) != numKanji {
		t.Fatalf("ParseKanjiDic2 length incorrect: got %d, want %d", got, numKanji)
	}
	// set kanjiList for later tests
	kanjidic2Parser = got
}

var strokeCountTests = []struct {
	Literal string
	Count   int
}{
	{"一", 1},
	{"二", 2},
	{"三", 3},
	{"中", 4},
	{"店", 8},
	{"感", 13},
	{"機", 16},
}

func TestStrokeCount(t *testing.T) {
	for _, tt := range strokeCountTests {
		if got := kanjidic2Parser[tt.Literal].StrokeCount; got != tt.Count {
			t.Errorf("TestStrokeCount (%s): got %d, want %d", tt.Literal, got, tt.Count)
		}
	}
}
