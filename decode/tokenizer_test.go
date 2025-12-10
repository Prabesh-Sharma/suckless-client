package decode

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestDecodeString(t *testing.T) {
	p := NewBencodeParser([]byte("11:hello world"))
	got, err := p.decodeString()
	if err != nil {
		fmt.Println(err)
	}
	want := "hello world"
	fmt.Println(p.offset)
	if got != want {
		t.Errorf("error occured while decoding got '%s' want '%s'", got, want)
	}
}

func TestDecodeList(t *testing.T) {
	p := NewBencodeParser([]byte("l5:helloi33ee"))
	got, err := p.decodeList()
	if err != nil {
		fmt.Println(err)
	}
	want := []any{"hello", 33}
	fmt.Println(p.offset)
	if !slices.Equal(want, got) {
		t.Errorf("error occured while decoding got '%v' want '%v'", got, want)
	}
}

func TestDecodeInt(t *testing.T) {
	p := NewBencodeParser([]byte("i3334e"))
	got, err := p.decodeInt()
	if err != nil {
		fmt.Println(err)
	}
	want := 3334
	fmt.Println(p.offset)
	if got != want {
		t.Errorf("error occured while decoding got '%d' want '%d'", got, want)
	}
}

func TestDecodeDict(t *testing.T) {
	p := NewBencodeParser([]byte("d5:helloi33ee"))

	got, err := p.decodeDict()
	if err != nil {
		t.Fatal(err)
	}

	want := map[string]any{"hello": 33}

	if !reflect.DeepEqual(want, got) {
		t.Errorf("error occurred while decoding got '%v' want '%v'", got, want)
	}
}
