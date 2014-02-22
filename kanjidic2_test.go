package kanjidic2

import (
	"testing"
)

var numKanji = 13108

func TestParser(t *testing.T) {
	got, err := ParseKanjiDic2("kanjidic2.xml")
	if err != nil {
		t.Fatalf("ParseKanjiDic2: %v", err)
	}
	if len(got) != numKanji {
		t.Fatalf("ParseKanjiDic2 length incorrect: got %d, want %d", got, numKanji)
	}
}
